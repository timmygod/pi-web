#!/usr/bin/env python3
"""Unit tests for pi_web.py. Uses a fake urlopen — does not start pi-web."""

import io
import json
import sys
import tempfile
import unittest
import urllib.error
from pathlib import Path
from unittest.mock import patch

sys.path.insert(0, str(Path(__file__).resolve().parent))
import pi_web


class FakeResponse:
    def __init__(self, body, status=200):
        if isinstance(body, str):
            body = body.encode()
        self._body = body
        self.status = status

    def read(self):
        return self._body

    def __enter__(self):
        return self

    def __exit__(self, *exc):
        return False


class FakeHTTPError(urllib.error.HTTPError):
    def __init__(self, code, body):
        super().__init__("http://127.0.0.1/x", code, "err", hdrs=None, fp=io.BytesIO(body.encode()))


class TestDiscover(unittest.TestCase):
    def test_loopback_uses_state_port(self):
        url = pi_web.discover_base_url(state={"port": "31416", "host": "100.64.0.1"})
        self.assertEqual(url, "http://127.0.0.1:31416")

    def test_missing_state_defaults_port(self):
        with tempfile.TemporaryDirectory() as tmp:
            env = {"PI_CODING_AGENT_DIR": tmp}
            url = pi_web.discover_base_url(env=env, homedir=tmp)
            self.assertEqual(url, "http://127.0.0.1:31415")

    def test_reads_state_file(self):
        with tempfile.TemporaryDirectory() as tmp:
            env = {"PI_CODING_AGENT_DIR": tmp}
            path = Path(tmp) / "pi-web" / "pi-web-state.json"
            path.parent.mkdir()
            path.write_text(json.dumps({"port": "9999", "host": "127.0.0.1"}))
            url = pi_web.discover_base_url(env=env)
            self.assertEqual(url, "http://127.0.0.1:9999")

    def test_ignores_dev_state_filename(self):
        with tempfile.TemporaryDirectory() as tmp:
            env = {"PI_CODING_AGENT_DIR": tmp}
            web = Path(tmp) / "pi-web"
            web.mkdir()
            (web / "pi-web-state-dev.json").write_text(json.dumps({"port": "31416"}))
            (web / "pi-web-state.json").write_text(json.dumps({"port": "31415"}))
            url = pi_web.discover_base_url(env=env)
            self.assertEqual(url, "http://127.0.0.1:31415")


class TestToken(unittest.TestCase):
    def test_env_wins(self):
        with tempfile.TemporaryDirectory() as tmp:
            env_file = Path(tmp) / ".config" / "pi-web"
            env_file.mkdir(parents=True)
            (env_file / "env").write_text("PI_WEB_TOKEN=fromfile\n")
            token = pi_web.read_token(env={"PI_WEB_TOKEN": "fromenv"}, homedir=tmp)
            self.assertEqual(token, "fromenv")

    def test_reads_env_file(self):
        with tempfile.TemporaryDirectory() as tmp:
            env_file = Path(tmp) / ".config" / "pi-web"
            env_file.mkdir(parents=True)
            (env_file / "env").write_text("PATH=/bin\nPI_WEB_TOKEN=secret\n")
            token = pi_web.read_token(env={}, homedir=tmp)
            self.assertEqual(token, "secret")

    def test_headers_use_x_pi_token_not_query(self):
        headers = pi_web.request_headers("secret")
        self.assertEqual(headers["X-Pi-Token"], "secret")
        self.assertNotIn("token", headers.get("Accept", "").lower())


class TestTimezoneAndCron(unittest.TestCase):
    def test_sg_alias(self):
        self.assertEqual(pi_web.resolve_timezone("sg"), "Asia/Singapore")
        self.assertEqual(pi_web.resolve_timezone("SGT"), "Asia/Singapore")
        self.assertEqual(pi_web.resolve_timezone("singapore"), "Asia/Singapore")

    def test_iana_passthrough(self):
        self.assertEqual(pi_web.resolve_timezone("Asia/Singapore"), "Asia/Singapore")

    def test_empty_timezone(self):
        self.assertEqual(pi_web.resolve_timezone(""), "")
        self.assertEqual(pi_web.resolve_timezone(None), "")

    def test_invalid_timezone(self):
        with self.assertRaises(pi_web.CtlError) as ctx:
            pi_web.resolve_timezone("not-a-zone")
        self.assertIn("unknown timezone", str(ctx.exception))
        self.assertIn("sg", str(ctx.exception))

    def test_build_cron_matches_js(self):
        self.assertEqual(pi_web.build_cron("manual"), "")
        self.assertEqual(pi_web.build_cron("hourly", minute=30), "30 * * * *")
        self.assertEqual(pi_web.build_cron("daily", minute=5, hour=9), "5 9 * * *")
        self.assertEqual(pi_web.build_cron("weekdays", minute=0, hour=8), "0 8 * * 1-5")
        self.assertEqual(pi_web.build_cron("weekly", minute=0, hour=17, weekday=5), "0 17 * * 5")
        self.assertEqual(pi_web.build_cron("every-hours", every_hours=2), "0 */2 * * *")

    def test_build_cron_clamps(self):
        self.assertEqual(pi_web.build_cron("daily", minute=99, hour=40), "59 23 * * *")

    def test_weekday_names(self):
        self.assertEqual(pi_web.parse_weekday("mon"), 1)
        self.assertEqual(pi_web.parse_weekday("Sunday"), 0)
        self.assertEqual(pi_web.parse_weekday("5"), 5)

    def test_default_name(self):
        self.assertEqual(pi_web.default_name("Summarize inbox"), "Summarize inbox")
        self.assertEqual(pi_web.default_name(""), "Scheduled task")
        long_name = "x" * 80
        self.assertEqual(len(pi_web.default_name(long_name)), 60)


class TestClientRequest(unittest.TestCase):
    def test_sends_token_header_not_query(self):
        captured = []

        def fake_urlopen(req, timeout=None):
            captured.append(req)
            return FakeResponse(json.dumps({"schedules": []}))

        with tempfile.TemporaryDirectory() as tmp:
            env = {"PI_CODING_AGENT_DIR": tmp, "PI_WEB_TOKEN": "s3cret"}
            client = pi_web.Client(env=env, homedir=tmp, urlopen=fake_urlopen)
            client.request("GET", "/api/schedules")

        req = captured[0]
        self.assertEqual(req.full_url, "http://127.0.0.1:31415/api/schedules")
        self.assertNotIn("token=", req.full_url)
        sent = {k.lower(): v for k, v in req.header_items()}
        self.assertEqual(sent.get("x-pi-token"), "s3cret")

    def test_health_failure_message(self):
        def fake_urlopen(req, timeout=None):
            raise OSError("connection refused")

        with tempfile.TemporaryDirectory() as tmp:
            env = {"PI_CODING_AGENT_DIR": tmp}
            client = pi_web.Client(env=env, homedir=tmp, urlopen=fake_urlopen)
            with self.assertRaises(pi_web.CtlError) as ctx:
                client.ensure_running()
        self.assertIn("not running", str(ctx.exception))
        self.assertIn("/pi-web start", str(ctx.exception))

    def test_api_error_uses_json_error_field(self):
        def fake_urlopen(req, timeout=None):
            raise FakeHTTPError(400, json.dumps({"error": "name is required"}))

        with tempfile.TemporaryDirectory() as tmp:
            env = {"PI_CODING_AGENT_DIR": tmp}
            client = pi_web.Client(env=env, homedir=tmp, urlopen=fake_urlopen)
            with self.assertRaises(pi_web.CtlError) as ctx:
                client.request("POST", "/api/schedules", {"name": ""})
        self.assertEqual(str(ctx.exception), "name is required")


class TestScheduleCommands(unittest.TestCase):
    def setUp(self):
        self.calls = []
        self.responses = {}

        def fake_urlopen(req, timeout=None):
            self.calls.append((req.get_method(), req.full_url, req.data))
            key = (req.get_method(), req.full_url.split("?", 1)[0])
            body = self.responses.get(key) or self.responses.get(req.get_method())
            if body is None:
                body = {"ok": True}
            return FakeResponse(json.dumps(body))

        self.tmp = tempfile.TemporaryDirectory()
        env = {"PI_CODING_AGENT_DIR": self.tmp.name}
        self.client = pi_web.Client(env=env, homedir=self.tmp.name, urlopen=fake_urlopen)

    def tearDown(self):
        self.tmp.cleanup()

    def test_create_daily_sg(self):
        self.responses["POST"] = {
            "schedule": {
                "id": "abc",
                "name": "Inbox",
                "cronExpr": "0 2 * * *",
                "timezone": "Asia/Singapore",
                "nextRunAt": "2026-09-18T18:00:00Z",
            }
        }
        buf = io.StringIO()
        args = pi_web.build_parser().parse_args(
            [
                "schedule",
                "create",
                "--name",
                "Inbox",
                "--instructions",
                "Summarize inbox",
                "--daily",
                "--hour",
                "2",
                "--minute",
                "0",
                "--timezone",
                "sg",
                "--project",
                "/tmp/assistant",
            ]
        )
        with patch("sys.stdout", buf):
            args.func(self.client, args)
        method, url, data = self.calls[0]
        self.assertEqual(method, "POST")
        self.assertTrue(url.endswith("/api/schedules"))
        payload = json.loads(data.decode())
        self.assertEqual(payload["cronExpr"], "0 2 * * *")
        self.assertEqual(payload["timezone"], "Asia/Singapore")
        self.assertEqual(payload["projectPath"], "/tmp/assistant")
        self.assertEqual(payload["name"], "Inbox")
        self.assertIn("abc", buf.getvalue())

    def test_create_defaults_name_from_instructions(self):
        self.responses["POST"] = {"schedule": {"id": "x", "name": "Do the thing"}}
        args = pi_web.build_parser().parse_args(
            ["schedule", "create", "--instructions", "Do the thing", "--manual", "--project", "/p"]
        )
        with patch("sys.stdout", io.StringIO()):
            args.func(self.client, args)
        payload = json.loads(self.calls[0][2].decode())
        self.assertEqual(payload["name"], "Do the thing")
        self.assertEqual(payload["cronExpr"], "")

    def test_delete_matches_by_name(self):
        self.responses[("GET", "http://127.0.0.1:31415/api/schedules")] = {
            "schedules": [{"id": "id-1", "name": "Inbox"}]
        }
        self.responses["DELETE"] = {"ok": True}
        args = pi_web.build_parser().parse_args(["schedule", "delete", "Inbox"])
        with patch("sys.stdout", io.StringIO()):
            args.func(self.client, args)
        methods = [c[0] for c in self.calls]
        self.assertIn("DELETE", methods)
        delete = [c for c in self.calls if c[0] == "DELETE"][0]
        self.assertIn("id=id-1", delete[1])

    def test_ambiguous_name(self):
        self.responses[("GET", "http://127.0.0.1:31415/api/schedules")] = {
            "schedules": [
                {"id": "a", "name": "Inbox"},
                {"id": "b", "name": "Inbox"},
            ]
        }
        args = pi_web.build_parser().parse_args(["schedule", "get", "Inbox"])
        with self.assertRaises(pi_web.CtlError) as ctx:
            args.func(self.client, args)
        self.assertIn("multiple schedules", str(ctx.exception))

    def test_split_model(self):
        self.assertEqual(pi_web.split_model("anthropic/claude"), ("anthropic", "claude"))
        with self.assertRaises(pi_web.CtlError):
            pi_web.split_model("claude")


class TestNotesAndSettings(unittest.TestCase):
    def test_notes_append_separator(self):
        self.assertEqual(pi_web.notes_append_chunk("old", "new"), "\n\nnew")
        self.assertEqual(pi_web.notes_append_chunk("", "new"), "new")
        self.assertEqual(pi_web.notes_append_chunk("   ", "new"), "new")

    def test_setting_aliases(self):
        self.assertEqual(pi_web.resolve_setting_key("theme"), "pi-web-theme")
        self.assertEqual(pi_web.resolve_setting_key("pi-web-theme"), "pi-web-theme")
        with self.assertRaises(pi_web.CtlError) as ctx:
            pi_web.resolve_setting_key("nope")
        self.assertIn("unknown setting", str(ctx.exception))

    def test_bool_coerce(self):
        self.assertEqual(pi_web.coerce_setting_value("pi-web:v1:auto-title:enabled", "off"), "false")
        self.assertEqual(pi_web.coerce_setting_value("pi-web:v1:auto-title:enabled", "on"), "true")
        self.assertEqual(pi_web.coerce_setting_value("pi-web-theme", "nord"), "nord")

    def test_notes_append_posts_mode(self):
        calls = []

        def fake_urlopen(req, timeout=None):
            calls.append((req.get_method(), req.full_url, req.data))
            if req.get_method() == "GET":
                return FakeResponse(json.dumps({"content": "existing"}))
            return FakeResponse(json.dumps({"ok": True, "content": "existing\n\nmore"}))

        with tempfile.TemporaryDirectory() as tmp:
            client = pi_web.Client(
                env={"PI_CODING_AGENT_DIR": tmp},
                homedir=tmp,
                urlopen=fake_urlopen,
            )
            args = pi_web.build_parser().parse_args(
                ["notes", "append", "--text", "more", "--project", "/p"]
            )
            with patch("sys.stdout", io.StringIO()):
                args.func(client, args)
        post = [c for c in calls if c[0] == "POST"][0]
        payload = json.loads(post[2].decode())
        self.assertEqual(payload["mode"], "append")
        self.assertEqual(payload["content"], "\n\nmore")
        self.assertEqual(payload["project"], "/p")

    def test_settings_set_uses_alias(self):
        calls = []

        def fake_urlopen(req, timeout=None):
            calls.append((req.get_method(), req.full_url, req.data))
            return FakeResponse(json.dumps({"ok": True, "settings": {"pi-web-theme": "nord"}}))

        with tempfile.TemporaryDirectory() as tmp:
            client = pi_web.Client(
                env={"PI_CODING_AGENT_DIR": tmp},
                homedir=tmp,
                urlopen=fake_urlopen,
            )
            args = pi_web.build_parser().parse_args(["settings", "set", "theme", "nord"])
            with patch("sys.stdout", io.StringIO()):
                args.func(client, args)
        payload = json.loads(calls[0][2].decode())
        self.assertEqual(payload, {"settings": {"pi-web-theme": "nord"}})


class TestMain(unittest.TestCase):
    def test_ctl_error_exits_1(self):
        class Boom:
            def request(self, *args, **kwargs):
                raise pi_web.CtlError("nope")

        with patch("sys.stderr", io.StringIO()):
            code = pi_web.main(["schedule", "list"], client=Boom())
        self.assertEqual(code, 1)


if __name__ == "__main__":
    unittest.main()
