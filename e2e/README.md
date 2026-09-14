# pi-web E2E (Playwright)

These tests run against the local-model edition. The Local Mode acceptance
scenarios and upstream synchronization policy are documented in
[Local-model edition development](../docs/dev/local-llm-development.md).

End-to-end browser tests that run against the built `pi-web` binary across
desktop, mobile, and iPad viewports.

```bash
make e2e-setup   # one-time: install deps + browsers
make e2e         # build binary + run the full suite
```

See **[docs/dev/e2e-testing.md](../docs/dev/e2e-testing.md)** for the full guide:
project matrix, scripted server lifecycle, sanitized fixtures, the stub `pi`, CI,
and how to add tests.
