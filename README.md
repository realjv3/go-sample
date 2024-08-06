# Go sample

This repository is a sample Go project that can be used to simulate CI / CD pipelines.

To trigger various behaviors, specify `-tags` when running Go commands. Available tags:

- `failtest`: runs a failing test
- `failbuild`: fails the build
- `failgenerate`: fails the code generation
- `failgeneratedirty`: passes code generation but leaves repository in dirty state
