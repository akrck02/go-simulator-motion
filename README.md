# go-simulator-motion

A [go-simulator](../go-simulator) plugin with the motions of a living scene.

| Rule | What it does | Parameters |
|---|---|---|
| `day_cycle` | Turns the sun around the celestial pole; the stars follow it | `seconds_per_day` (120) |
| `gusts` | Varies the wind strength and direction around the wind the world started with | `variation` (0.5), `veering` (0.35 rad) |
| `cloud_drift` | Moves the clouds with the wind | `speed` (0.05) |
| `tides` | Raises and lowers the water (turns the tide angle; heights and phases come with the scene) | `period_seconds` (0 = 12.42 h in proportion to `seconds_per_day`) |
| `entity_motion` | Moves the entities that ask for it in their properties | — |

Properties `entity_motion` reads (glTF node extras arrive as properties):

| Property | Motion |
|---|---|
| `"bob": {"height": h, "period": s, "roll": radians}` | Floats up and down with a slight roll |
| `"spin": {"axis": [x, y, z], "speed": radians per second}` | Turns around a local axis |
| `"path": {"points": [[x, y, z], ...], "speed": units per second, "loop": true}` | Follows the points, facing where it goes (+Z forward); without `loop` it turns around at the end |

An entity takes one motion: bob wins over spin, spin over path. Entities added later by other
plugins start moving as soon as they appear.

## Use

```bash
cd ../go-simulator
go run ./cmd/simulator -rules ../go-simulator-motion/examples/living-scene.yaml scene.glb
```

`examples/living-scene.yaml` runs the plugin with `go run`; after `go install` the command can be
just `[go-simulator-motion]`.

During development go-simulator and go-renderer are used from the folders next to this one
(`replace` in `go.mod`).
