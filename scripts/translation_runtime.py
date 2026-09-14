"""Shared local-model invocation for repository translation scripts."""

from __future__ import annotations

import os


DEFAULT_TRANSLATION_MODEL = (
    "llama-cpp/AtomicChat/Qwen3.8-Flash-Next-AD-4.27bpw-Q4_K_M-M64"
)
TRANSLATION_MODEL_ENV = "PI_TRANSLATION_MODEL"


def translation_model() -> str:
    """Return the pi model reference used for local translations."""
    return os.environ.get(TRANSLATION_MODEL_ENV, "").strip() or DEFAULT_TRANSLATION_MODEL


def pi_translation_command(prompt: str) -> list[str]:
    """Build a non-agentic, sessionless pi command for one translation."""
    return [
        "pi",
        "-p",
        "--model",
        translation_model(),
        "--no-session",
        "--no-tools",
        "--no-extensions",
        "--no-context-files",
        "--thinking",
        "off",
        prompt,
    ]
