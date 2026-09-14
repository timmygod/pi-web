# Bounded Local Mode checkpoints

Local workers replace Pi's accumulating summary with a fresh rolling checkpoint
on every compaction, including Force Compact. Cloud workers are unchanged.

The checkpoint prioritizes the current goal, hard constraints, active decisions,
unfinished work, blockers, and exact recovery references. Completed work is reduced
to outcomes; obsolete plans, repeated history, and tool transcripts are dropped.
The previous checkpoint is input evidence, not text that must be preserved in full.
Pi retains its original recent-message boundary and owns append-only persistence.

The first request targets 2,400 tokens with at most 4,096 generated tokens. The
checkpoint must finish normally, contain all five state sections and an end
marker, and fit both 6,000 Unicode characters and 12,000 UTF-8 bytes. These are
hard text-size limits, not claims of exact tokenizer counts. Budgets scale down
for small model context/output limits. An invalid or capped response triggers
one independent rewrite from the original inputs with a tighter target, never
from an incomplete draft. Each request has a five-minute timeout and honors
user cancellation. No default unbounded summarizer fallback is allowed.

After two failures, preserve the old context and stop with a diagnostic. Successful
automatic compaction still requires a meaningful reduction (5% or 512 estimated
tokens). This is protection against a loop, not a limit on the number of
successful checkpoints in a long project. Summarization is lossy: budgets reduce
growth but cannot guarantee perfect recall or progress through provider outages.

Tests exercise bounded output, cap/empty/error retries, cancellation, budget
scaling, preservation of the recent boundary, and no-progress protection.
