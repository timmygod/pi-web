#!/usr/bin/env python3
"""pi-web-ctl: call the local pi-web HTTP API from pi skills.

Talks to 127.0.0.1 using the port in pi-web-state.json. Auth uses X-Pi-Token
(never a ?token= query — that 302s and skips the API handler).
"""

from __future__ import annotations

import argparse
import json
import os
import sys
import urllib.error
import urllib.parse
import urllib.request
from pathlib import Path
from zoneinfo import ZoneInfo, ZoneInfoNotFoundError

DEFAULT_PORT = "31415"

TZ_ALIASES = {
    "sg": "Asia/Singapore",
    "sgt": "Asia/Singapore",
    "singapore": "Asia/Singapore",
    "jst": "Asia/Tokyo",
    "tokyo": "Asia/Tokyo",
    "jp": "Asia/Tokyo",
    "utc": "UTC",
    "gmt": "UTC",
    "pt": "America/Los_Angeles",
    "pst": "America/Los_Angeles",
    "pdt": "America/Los_Angeles",
    "et": "America/New_York",
    "est": "America/New_York",
    "edt": "America/New_York",
    "ct": "America/Chicago",
    "cst": "America/Chicago",
    "cdt": "America/Chicago",
    "london": "Europe/London",
    "uk": "Europe/London",
    "bst": "Europe/London",
}

WEEKDAYS = {
    "sun": 0,
    "sunday": 0,
    "mon": 1,
    "monday": 1,
    "tue": 2,
    "tues": 2,
    "tuesday": 2,
    "wed": 3,
    "wednesday": 3,
    "thu": 4,
    "thur": 4,
    "thurs": 4,
    "thursday": 4,
    "fri": 5,
    "friday": 5,
    "sat": 6,
    "saturday": 6,
}


SETTING_ALIASES = {
    "theme": "pi-web-theme",
    "language": "pi-web:v1:locale",
    "locale": "pi-web:v1:locale",
    "font-ui": "pi-web:v1:font-ui",
    "font-content": "pi-web:v1:font-content",
    "font-code": "pi-web:v1:font-code",
    "font-ui-size": "pi-web:v1:font-ui-size",
    "font-content-size": "pi-web:v1:font-content-size",
    "spinner": "pi-sessions:spinner-style",
    "notify-on-done": "pi-share:v1:notify-on-done",
    "done-sound": "pi-share:v1:done-sound",
    "layout": "pi-sessions:view-layout",
    "show-btw": "pi-web:v1:show-btw-in-index",
    "cat": "pi-web:v1:cat:enabled",
    "cat-focus": "pi-web:v1:cat:focus-min",
    "cat-break": "pi-web:v1:cat:break-min",
    "bedtime": "pi-web:v1:cat:bedtime",
    "wakeup": "pi-web:v1:cat:wakeup",
    "sleep-min": "pi-web:v1:cat:sleep-min",
    "auto-title": "pi-web:v1:auto-title:enabled",
    "auto-title-mode": "pi-web:v1:auto-title:mode",
    "auto-title-model": "pi-web:v1:auto-title:model",
    "artifacts": "pi-web:v1:artifacts:enabled",
    "artifacts-include": "pi-web:v1:artifacts:include",
    "thinking": "pi-web:v1:toggle:thinking",
    "tools": "pi-web:v1:toggle:tools",
    "tool-outputs": "pi-web:v1:toggle:tool-outputs",
}

SETTING_KEYS = set(SETTING_ALIASES.values())
BOOL_SETTING_KEYS = {
    "pi-share:v1:notify-on-done",
    "pi-web:v1:show-btw-in-index",
    "pi-web:v1:cat:enabled",
    "pi-web:v1:auto-title:enabled",
    "pi-web:v1:artifacts:enabled",
    "pi-web:v1:toggle:thinking",
    "pi-web:v1:toggle:tools",
    "pi-web:v1:toggle:tool-outputs",
}


class CtlError(Exception):
    """User-facing CLI error; main() prints the message and exits 1."""


def agent_dir(env=None, homedir=None):
    env = os.environ if env is None else env
    if env.get("PI_CODING_AGENT_DIR"):
        return Path(env["PI_CODING_AGENT_DIR"]).expanduser()
    home = Path.home() if homedir is None else Path(homedir)
    return home / ".pi" / "agent"


def state_path(env=None, homedir=None):
    return agent_dir(env=env, homedir=homedir) / "pi-web" / "pi-web-state.json"


def read_state(env=None, homedir=None):
    path = state_path(env=env, homedir=homedir)
    try:
        return json.loads(path.read_text())
    except FileNotFoundError:
        return None
    except (OSError, json.JSONDecodeError) as err:
        raise CtlError(f"could not read {path}: {err}") from err


def discover_base_url(env=None, homedir=None, state=None):
    """Always loopback. State host is ignored so we never send the token off-box."""
    if state is None:
        state = read_state(env=env, homedir=homedir)
    port = DEFAULT_PORT
    if isinstance(state, dict) and str(state.get("port") or "").strip():
        port = str(state["port"]).strip()
    return f"http://127.0.0.1:{port}"


def read_token(env=None, homedir=None):
    env = os.environ if env is None else env
    from_env = (env.get("PI_WEB_TOKEN") or "").strip()
    if from_env:
        return from_env
    home = Path.home() if homedir is None else Path(homedir)
    path = home / ".config" / "pi-web" / "env"
    try:
        raw = path.read_text()
    except OSError:
        return None
    for line in raw.splitlines():
        if line.startswith("PI_WEB_TOKEN="):
            value = line.split("=", 1)[1].strip()
            return value or None
    return None


def request_headers(token):
    headers = {"Accept": "application/json"}
    if token:
        headers["X-Pi-Token"] = token
    return headers


def clamp_int(value, lo, hi, fallback):
    try:
        n = int(value)
    except (TypeError, ValueError):
        return fallback
    return min(hi, max(lo, n))


def parse_weekday(value):
    if value is None or value == "":
        return 1
    key = str(value).strip().lower()
    if key in WEEKDAYS:
        return WEEKDAYS[key]
    try:
        n = int(key)
    except ValueError as err:
        raise CtlError(f"unknown weekday {value!r}; use sun-sat or 0-6") from err
    if n < 0 or n > 6:
        raise CtlError(f"weekday must be 0-6, got {value!r}")
    return n


def build_cron(frequency, hour=9, minute=0, weekday=1, every_hours=None):
    """Match web/src/index/schedules.js buildCron, plus --every-hours."""
    m = clamp_int(minute, 0, 59, 0)
    h = clamp_int(hour, 0, 23, 9)
    d = clamp_int(weekday, 0, 6, 1)
    if frequency in (None, "", "manual"):
        return ""
    if frequency == "hourly":
        return f"{m} * * * *"
    if frequency == "daily":
        return f"{m} {h} * * *"
    if frequency == "weekdays":
        return f"{m} {h} * * 1-5"
    if frequency == "weekly":
        return f"{m} {h} * * {d}"
    if frequency == "every-hours":
        n = clamp_int(every_hours, 1, 23, 1)
        return f"{m} */{n} * * *"
    raise CtlError(f"unknown frequency {frequency!r}")


def resolve_timezone(value):
    raw = (value or "").strip()
    if not raw:
        return ""
    compact = raw.lower().replace(" ", "")
    name = TZ_ALIASES.get(compact) or TZ_ALIASES.get(raw.lower()) or raw
    try:
        ZoneInfo(name)
    except ZoneInfoNotFoundError as err:
        aliases = ", ".join(sorted(set(TZ_ALIASES)))
        raise CtlError(
            f"unknown timezone {value!r}; use an IANA name or one of: {aliases}"
        ) from err
    except Exception as err:
        # zoneinfo can raise other errors on some platforms for junk input.
        raise CtlError(f"unknown timezone {value!r}: {err}") from err
    return name


def split_model(value):
    raw = (value or "").strip()
    if not raw:
        return "", ""
    if "/" not in raw:
        raise CtlError("--model must be provider/id (example: anthropic/claude-sonnet-4-5)")
    provider, model_id = raw.split("/", 1)
    provider, model_id = provider.strip(), model_id.strip()
    if not provider or not model_id:
        raise CtlError("--model must be provider/id")
    return provider, model_id


def default_name(instructions):
    line = (instructions or "").strip().splitlines()[0].strip() if instructions else ""
    if not line:
        return "Scheduled task"
    if len(line) > 60:
        return line[:57].rstrip() + "..."
    return line


def _url_has_token_query(url):
    return "token=" in url.split("?", 1)[-1].lower() if "?" in url else False


class Client:
    def __init__(self, env=None, homedir=None, urlopen=urllib.request.urlopen, timeout=10):
        self.env = os.environ if env is None else env
        self.homedir = homedir
        self.urlopen = urlopen
        self.timeout = timeout
        self.base_url = discover_base_url(env=self.env, homedir=self.homedir)
        self.token = read_token(env=self.env, homedir=self.homedir)

    def health_ok(self):
        url = self.base_url + "/"
        req = urllib.request.Request(url, method="GET")
        try:
            with self.urlopen(req, timeout=1) as resp:
                return getattr(resp, "status", 200) in (200, 401, 403)
        except urllib.error.HTTPError as err:
            return err.code in (401, 403)
        except OSError:
            return False

    def ensure_running(self):
        if self.health_ok():
            return
        raise CtlError(
            f"pi-web is not running (tried {self.base_url}). Start it with /pi-web start."
        )

    def request(self, method, path, body=None):
        url = self.base_url + path
        if _url_has_token_query(url):
            raise CtlError("refusing to put the token in the query string")
        data = None
        headers = request_headers(self.token)
        if body is not None:
            data = json.dumps(body).encode()
            headers["Content-Type"] = "application/json"
        req = urllib.request.Request(url, data=data, headers=headers, method=method)
        try:
            with self.urlopen(req, timeout=self.timeout) as resp:
                raw = resp.read()
        except urllib.error.HTTPError as err:
            raw = err.read()
            message = _api_error_message(raw, err)
            raise CtlError(message) from err
        except OSError as err:
            raise CtlError(f"request failed: {err}") from err
        if not raw:
            return {}
        try:
            return json.loads(raw.decode())
        except json.JSONDecodeError as err:
            raise CtlError(f"pi-web returned non-JSON: {raw[:200]!r}") from err


def _api_error_message(raw, err):
    text = raw.decode(errors="replace") if raw else ""
    try:
        payload = json.loads(text)
        if isinstance(payload, dict) and payload.get("error"):
            return str(payload["error"])
    except json.JSONDecodeError:
        pass
    if text.strip():
        return f"HTTP {err.code}: {text.strip()[:300]}"
    return f"HTTP {err.code}"


def emit(payload):
    json.dump(payload, sys.stdout, indent=2)
    sys.stdout.write("\n")


def _resolve_schedule(client, token):
    token = (token or "").strip()
    if not token:
        raise CtlError("schedule id or name is required")
    data = client.request("GET", "/api/schedules")
    items = data.get("schedules") or []
    for sc in items:
        if sc.get("id") == token:
            return sc
    matches = [sc for sc in items if (sc.get("name") or "").lower() == token.lower()]
    if len(matches) == 1:
        return matches[0]
    if len(matches) > 1:
        ids = ", ".join(sc.get("id", "") for sc in matches)
        raise CtlError(f"multiple schedules named {token!r}; pass an id: {ids}")
    raise CtlError(f"schedule not found: {token}")


def _schedule_input(sc, **overrides):
    body = {
        "name": sc.get("name") or "",
        "instructions": sc.get("instructions") or "",
        "modelProvider": sc.get("modelProvider") or "",
        "modelId": sc.get("modelId") or "",
        "thinkingLevel": sc.get("thinkingLevel") or "",
        "projectPath": sc.get("projectPath") or "",
        "cronExpr": sc.get("cronExpr") or "",
        "timezone": sc.get("timezone") or "",
        "enabled": sc.get("enabled", True),
    }
    body.update(overrides)
    return body


def cmd_schedule_list(client, _args):
    emit(client.request("GET", "/api/schedules"))


def cmd_schedule_get(client, args):
    sc = _resolve_schedule(client, args.target)
    emit({"schedule": sc})


def cmd_schedule_create(client, args):
    instructions = args.instructions
    if not (instructions or "").strip():
        raise CtlError("--instructions is required")
    cron_expr = args.cron
    if cron_expr is None:
        frequency = args.frequency
        if args.every_hours is not None:
            frequency = "every-hours"
        elif frequency is None:
            frequency = "manual"
        cron_expr = build_cron(
            frequency,
            hour=args.hour,
            minute=args.minute,
            weekday=parse_weekday(args.weekday),
            every_hours=args.every_hours,
        )
    timezone = resolve_timezone(args.timezone)
    provider, model_id = split_model(args.model)
    project = args.project if args.project is not None else os.getcwd()
    name = (args.name or "").strip() or default_name(instructions)
    body = {
        "name": name,
        "instructions": instructions,
        "modelProvider": provider,
        "modelId": model_id,
        "thinkingLevel": (args.thinking or "").strip(),
        "projectPath": project,
        "cronExpr": cron_expr,
        "timezone": timezone,
        "enabled": not args.paused,
    }
    created = client.request("POST", "/api/schedules", body)
    emit(created)


def cmd_schedule_update(client, args):
    sc = _resolve_schedule(client, args.target)
    overrides = {}
    if args.name:
        overrides["name"] = args.name
    if args.instructions:
        overrides["instructions"] = args.instructions
    if args.project is not None:
        overrides["projectPath"] = args.project
    if args.timezone is not None:
        overrides["timezone"] = resolve_timezone(args.timezone)
    if args.thinking is not None:
        overrides["thinkingLevel"] = args.thinking
    if args.model is not None:
        provider, model_id = split_model(args.model)
        overrides["modelProvider"] = provider
        overrides["modelId"] = model_id
    if args.cron is not None:
        overrides["cronExpr"] = args.cron
    elif args.frequency or args.every_hours is not None:
        frequency = "every-hours" if args.every_hours is not None else args.frequency
        overrides["cronExpr"] = build_cron(
            frequency,
            hour=args.hour,
            minute=args.minute,
            weekday=parse_weekday(args.weekday),
            every_hours=args.every_hours,
        )
    if args.paused:
        overrides["enabled"] = False
    body = _schedule_input(sc, **overrides)
    emit(client.request("POST", f"/api/schedule?id={sc['id']}", body))


def cmd_schedule_delete(client, args):
    sc = _resolve_schedule(client, args.target)
    emit(client.request("DELETE", f"/api/schedule?id={sc['id']}"))


def cmd_schedule_enable(client, args):
    sc = _resolve_schedule(client, args.target)
    body = _schedule_input(sc, enabled=True)
    emit(client.request("POST", f"/api/schedule?id={sc['id']}", body))


def cmd_schedule_disable(client, args):
    sc = _resolve_schedule(client, args.target)
    body = _schedule_input(sc, enabled=False)
    emit(client.request("POST", f"/api/schedule?id={sc['id']}", body))


def cmd_schedule_run(client, args):
    sc = _resolve_schedule(client, args.target)
    emit(client.request("POST", f"/api/schedule/run?id={sc['id']}", {}))


def cmd_schedule_runs(client, args):
    sc = _resolve_schedule(client, args.target)
    emit(client.request("GET", f"/api/schedule/runs?id={sc['id']}"))


def _project_path(args):
    if args.project is not None:
        return args.project
    return os.getcwd()


def notes_append_chunk(existing, text):
    if (existing or "").strip() and text:
        return "\n\n" + text
    return text


def cmd_notes_read(client, args):
    project = _project_path(args)
    emit(client.request("GET", "/api/scratchpad?project=" + urllib.parse.quote(project)))


def cmd_notes_append(client, args):
    project = _project_path(args)
    text = args.text if args.text is not None else ""
    existing = ""
    try:
        existing = (client.request("GET", "/api/scratchpad?project=" + urllib.parse.quote(project)).get("content") or "")
    except CtlError:
        pass
    chunk = notes_append_chunk(existing, text)
    emit(
        client.request(
            "POST",
            "/api/scratchpad",
            {"project": project, "content": chunk, "mode": "append"},
        )
    )


def cmd_notes_replace(client, args):
    project = _project_path(args)
    emit(
        client.request(
            "POST",
            "/api/scratchpad",
            {"project": project, "content": args.text if args.text is not None else "", "mode": "replace"},
        )
    )


def resolve_setting_key(alias):
    raw = (alias or "").strip()
    if not raw:
        raise CtlError("setting name is required")
    if raw in SETTING_ALIASES:
        return SETTING_ALIASES[raw]
    if raw in SETTING_KEYS:
        return raw
    aliases = ", ".join(sorted(SETTING_ALIASES))
    raise CtlError(f"unknown setting {alias!r}; known aliases: {aliases}")


def coerce_setting_value(key, value):
    text = "" if value is None else str(value)
    if key not in BOOL_SETTING_KEYS:
        return text
    lowered = text.strip().lower()
    if lowered in ("1", "true", "on", "yes"):
        return "true"
    if lowered in ("0", "false", "off", "no"):
        return "false"
    raise CtlError(f"{key} expects on/off (got {value!r})")


def cmd_settings_get(client, args):
    data = client.request("GET", "/api/settings")
    settings = data.get("settings") or {}
    if args.key:
        key = resolve_setting_key(args.key)
        emit({"key": key, "value": settings.get(key, "")})
        return
    emit({"settings": settings, "aliases": SETTING_ALIASES})


def cmd_settings_set(client, args):
    key = resolve_setting_key(args.key)
    value = coerce_setting_value(key, args.value)
    emit(client.request("POST", "/api/settings", {"settings": {key: value}}))


def _add_schedule_write_flags(parser, *, for_update=False):
    parser.add_argument("--name")
    inst = parser.add_argument("--instructions")
    if not for_update:
        inst.required = True
    parser.add_argument("--project", help="Project path (default: current directory)")
    parser.add_argument("--timezone", help="IANA name or alias (sg, jst, utc, pt, et, …)")
    parser.add_argument("--model", help="provider/id")
    parser.add_argument("--thinking")
    parser.add_argument("--hour", type=int, default=9)
    parser.add_argument("--minute", type=int, default=0)
    parser.add_argument("--weekday", default="mon", help="sun-sat or 0-6 (weekly)")
    parser.add_argument("--paused", action="store_true", help="Create disabled / pause on update")
    freq = parser.add_mutually_exclusive_group()
    freq.add_argument("--daily", action="store_const", const="daily", dest="frequency")
    freq.add_argument("--hourly", action="store_const", const="hourly", dest="frequency")
    freq.add_argument("--weekdays", action="store_const", const="weekdays", dest="frequency")
    freq.add_argument("--weekly", action="store_const", const="weekly", dest="frequency")
    freq.add_argument("--manual", action="store_const", const="manual", dest="frequency")
    freq.add_argument("--cron", help="Raw 5-field cron expression")
    freq.add_argument("--every-hours", type=int, dest="every_hours")


def build_parser():
    parser = argparse.ArgumentParser(
        prog="pi-web-ctl",
        description="Call the local pi-web HTTP API (schedules, notes, settings).",
    )
    sub = parser.add_subparsers(dest="group", required=True)

    sched = sub.add_parser("schedule", help="Create and manage pi-web schedules")
    sched_sub = sched.add_subparsers(dest="action", required=True)

    p = sched_sub.add_parser("list")
    p.set_defaults(func=cmd_schedule_list)

    p = sched_sub.add_parser("get")
    p.add_argument("target", help="Schedule id or exact name")
    p.set_defaults(func=cmd_schedule_get)

    p = sched_sub.add_parser("create")
    _add_schedule_write_flags(p)
    p.set_defaults(func=cmd_schedule_create)

    p = sched_sub.add_parser("update")
    p.add_argument("target", help="Schedule id or exact name")
    _add_schedule_write_flags(p, for_update=True)
    p.set_defaults(func=cmd_schedule_update)

    p = sched_sub.add_parser("delete")
    p.add_argument("target")
    p.set_defaults(func=cmd_schedule_delete)

    p = sched_sub.add_parser("enable")
    p.add_argument("target")
    p.set_defaults(func=cmd_schedule_enable)

    p = sched_sub.add_parser("disable")
    p.add_argument("target")
    p.set_defaults(func=cmd_schedule_disable)

    p = sched_sub.add_parser("run")
    p.add_argument("target")
    p.set_defaults(func=cmd_schedule_run)

    p = sched_sub.add_parser("runs")
    p.add_argument("target")
    p.set_defaults(func=cmd_schedule_runs)

    notes = sub.add_parser("notes", help="Read or write the per-project scratchpad")
    notes_sub = notes.add_subparsers(dest="action", required=True)
    p = notes_sub.add_parser("read")
    p.add_argument("--project")
    p.set_defaults(func=cmd_notes_read)
    p = notes_sub.add_parser("append")
    p.add_argument("--text", required=True)
    p.add_argument("--project")
    p.set_defaults(func=cmd_notes_append)
    p = notes_sub.add_parser("replace")
    p.add_argument("--text", required=True)
    p.add_argument("--project")
    p.set_defaults(func=cmd_notes_replace)

    settings = sub.add_parser("settings", help="Read or change pi-web settings")
    settings_sub = settings.add_subparsers(dest="action", required=True)
    p = settings_sub.add_parser("get")
    p.add_argument("key", nargs="?", help="Alias or storage key")
    p.set_defaults(func=cmd_settings_get)
    p = settings_sub.add_parser("set")
    p.add_argument("key")
    p.add_argument("value")
    p.set_defaults(func=cmd_settings_set)

    return parser


def main(argv=None, client=None):
    parser = build_parser()
    args = parser.parse_args(argv)
    try:
        if client is None:
            client = Client()
            client.ensure_running()
        args.func(client, args)
    except CtlError as err:
        print(err, file=sys.stderr)
        return 1
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
