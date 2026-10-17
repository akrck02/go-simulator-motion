# 🚀 Getting started

## The folders it needs

The plugin is developed next to the projects it uses: `go.mod` replaces go-simulator and
go-renderer with the folders beside this one. Clone them side by side:

```
go-renderer/
go-simulator/
go-simulator-motion/   ← this repository
```

## Running the example

```bash
cd ../go-simulator
go run ./cmd/simulator -rules ../go-simulator-motion/examples/living-scene.yaml scene.glb
```

`scene.glb` is any glTF 2.0 scene. The viewer shows it while the plugin moves it.

`examples/living-scene.yaml` starts the plugin with `go run` from this repository, so it needs no
installing. After `go install github.com/akrck02/go-simulator-motion@latest` the command can be
just `[go-simulator-motion]`.

```yaml
plugins:
  - name: motion
    command: [go, run, .]
    directory: ..

stages:
  - name: motion
    clock: realtime
    rules:
      - rule: day_cycle
        with: {seconds_per_day: 120}
        enabled: false
        toggle_key: T
      - rule: gusts
        with: {variation: 0.5, veering: 0.35}
      - rule: cloud_drift
      - rule: tides
        with: {seconds_per_day: 120}
        toggle_key: M
      - rule: entity_motion
```

All the rules sit in one stage on a `realtime` clock, so they follow the viewer's clock. The day
cycle starts switched off; the rules file's [reference](https://github.com/akrck02/go-simulator/blob/main/docs/4-guides/1-rules-files.md)
explains every key.

## The keys

| Key | What it does |
|---|---|
| **T** | Turns the day cycle on and off (`toggle_key` in the example) |
| **M** | Turns the tides off and on (`toggle_key` in the example) |
| **P** | Pauses and resumes time |
| **,** and **.** | Halve or double the speed of time |

The viewer's own keys (moving the camera and so on) work as usual.

## Asking for a motion in a scene

glTF node extras arrive as the entity's properties, so a scene asks for a motion without any code.
A boat that bobs, sails that turn and a ferry on a loop, in a scene's `nodes`:

```json
{ "name": "boat",  "extras": { "bob":  { "height": 0.5, "period": 4, "roll": 0.05 } } },
{ "name": "sails", "extras": { "spin": { "axis": [0, 0, 1], "speed": 0.8 } } },
{ "name": "ferry", "extras": { "path": { "points": [[-30, 0, -30], [30, 0, -30], [30, 0, 30], [-30, 0, 30]],
                                         "speed": 4, "loop": true } } }
```

Every key and what it is when left out are in [motions](4-the-model.md). Entities that other
plugins add later with these properties start moving too.

## The tests

```bash
go vet ./...
go test ./...
```

The tests run the plugin inside a go-simulator host and check that
the sun turns, the clouds drift with the wind, the tide turns and can be switched off, the wind
gusts within its range, that a path goes there and back, and that bob wins over spin.
