# go-simulator-motion

> **Note:** this project was written with an AI assistant and is pending thorough review. See the [AI disclosure](#ai-disclosure).

A [go-simulator](https://github.com/akrck02/go-simulator) plugin with the motions of a living scene: the passing of the
day, wind gusts, drifting clouds, tides and the motions entities ask for in their properties
(bob, spin, path).

```bash
cd ../go-simulator
go run ./cmd/simulator -rules ../go-simulator-motion/examples/living-scene.yaml scene.glb
```

Documentation in [docs/](docs/README.md): [motions](docs/model.md) and the
[rules reference](docs/rules.md).

## AI disclosure

I write most of my code by hand, so I want to be clear about this project: **it was written by an
AI assistant (Claude, by Anthropic) under my direction.** I decided what to build and how it should
behave; the code, the tests and the documentation were written by the AI. Every commit made this
way carries a `Co-Authored-By: Claude` line (all of them).

**Status: pending thorough review.** The code compiles and its tests pass, but I have not yet
reviewed it line by line. Treat it as a working draft, not as code I vouch for.
