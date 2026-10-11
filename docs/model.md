# Motions

A still scene feels like a photograph. The motion plugin makes it feel alive: the sun crosses the
sky, the wind picks up and dies down, clouds drift past, the sea rises and falls, boats bob and mill
sails turn. It works in real time, while you watch the 3D scene. It changes the environment (sun, wind, clouds,
tide) and the transforms of entities that ask for a motion. Its rules belong on a `realtime` clock;
the speed keys of the viewer (`,` `.`) and the pause (`P`) apply to them.

## The environment

| Rule | What it does |
|---|---|
| `day_cycle` | Turns the sun around the celestial pole once every `seconds_per_day` (120); the stars follow it, and the renderer derives day, twilight and night from the sun's height |
| `gusts` | Varies the wind's strength (`variation`, 0.5) and direction (`veering`, 0.35 rad) around the wind the world started with; plants marked to sway bend with it and clouds drift |
| `cloud_drift` | Moves the clouds with the wind (`speed`, 0.05) |
| `tides` | Turns the tide angle a full turn per period (`period_seconds`; 0 = 12.42 h in proportion to `seconds_per_day`). The tide's heights and phases come with the scene; the renderer raises the water from them |

## Entities

`entity_motion` moves the entities that ask for it in their properties. glTF node extras arrive as
properties, so a scene can ask for motions without any code.

| Property | Motion |
|---|---|
| `"bob": {"height": h, "period": s, "roll": radians}` | Floats up and down with a slight roll (a boat) |
| `"spin": {"axis": [x, y, z], "speed": radians per second}` | Turns around a local axis (a mill's sails) |
| `"path": {"points": [[x, y, z], …], "speed": units per second, "loop": true}` | Follows the points facing where it goes (+Z forward); without `loop` it turns around at the end |

An entity takes one motion: bob wins over spin, spin over path. Entities added later by other
plugins start moving as soon as they appear: the rule reads entities with those properties.

## Use

```bash
cd ../go-simulator
go run ./cmd/simulator -rules ../go-simulator-motion/examples/living-scene.yaml scene.glb
```

`examples/living-scene.yaml` runs the plugin with `go run`; after `go install` the command can be
just `[go-simulator-motion]`. During development go-simulator and go-renderer are used from the
folders next to this one (`replace` in `go.mod`).
