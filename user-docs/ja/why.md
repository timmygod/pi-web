# Why pi-web?

I'm kind of addicted to Claude Code. I am always using it. If I am not sitting in front of the computer, I am thinking about it. I feel like I am not burning enough tokens. It was the early days of Claude Code. And I was thinking, why can't I resume from my phone? I set up Termius and I didn't really like it.

I started to create my own and stopped when Claude introduced their Claude Code mobile app.

Then I got a herniated disc and I couldn't really do anything that much. Time went on and I felt recovered a bit and I wanted to continue my Claude Code via web/PWA project.

Then Claude Code started banning usage outside of their own harness. And I feel like it's not worth it.

Then I found pi.dev and explored a bit but hadn't really dived in. I read about it, watched videos about it and decided to give a full try and now I am totally into pi.

Since it's open source I feel like it's worth building for. I get different provider choices as well. I also feel like relying on one provider/model like Anthropic/Claude is not sustainable.

So I am building it here.

This checkout is maintained as a local-model edition of pi-web. It follows the
upstream project for shared improvements, while keeping local deployment,
context stability, and local-model testing on a separately released track.

## Why a local model needs a different operating profile

The original pi-web experience is an excellent foundation, but local inference
has different failure modes from a typical hosted model. A local model may slow
down sharply as context grows, share limited memory with the rest of the machine,
stop after producing only reasoning, or lose a long run to a transient local
transport failure. Treating those cases exactly like cloud failures makes the UI
look compatible while the actual session remains fragile.

This edition approaches the problem in layers:

1. **Preserve upstream first.** Shared UI and session behavior continue to come
   from pi-web; local changes are isolated behind effective Local Mode.
2. **Prevent before recovering.** A percentage-based 65% context boundary is
   enforced before subsequent model calls, including calls inside long tool loops.
3. **Recover only with evidence.** Automatic continuation is limited to recognized
   context, transport, and thinking-only incidents—not authentication, quota, or
   arbitrary provider errors.
4. **Bound every autonomous action.** Recovery incidents are deduplicated,
   progress is required before another rescue, and startup considers at most one
   recently active Local session.
5. **Keep a manual exit.** Force Compact summarizes rather than wipes history, so
   the user can rescue a session without pretending the context never existed.
6. **Protect cloud compatibility.** Cloud Mode keeps the upstream semantics and
   controls; local-model optimizations do not silently redefine cloud sessions.

That is the real difference in this fork: it treats local inference as a distinct
operational environment, not merely another model name in a dropdown.
