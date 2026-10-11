# go-simulator-motion

A [go-simulator](https://github.com/akrck02/go-3d-simulator) plugin with the motions of a living scene: the passing of the
day, wind gusts, drifting clouds, tides and the motions entities ask for in their properties
(bob, spin, path).

```bash
cd ../go-simulator
go run ./cmd/simulator -rules ../go-simulator-motion/examples/living-scene.yaml scene.glb
```

Documentation in [docs/](docs/README.md): [motions](docs/model.md) and the
[rules reference](docs/rules.md).
