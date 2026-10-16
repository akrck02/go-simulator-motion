# 📏 Rules of the motion plugin

Version 1.0.0, protocol 2. Day cycle, wind gusts, drifting clouds, tides and the motions entities ask for (bob, spin, path).

This reference is generated from what the plugin itself tells the host (the `describe` message), so it always
matches the code. Parameters go under `with:` in a rules file; see go-simulator's
[rules file reference](https://github.com/akrck02/go-simulator/blob/main/docs/rules-files.md).

| Rule | What it does |
|---|---|
| [`day_cycle`](#day_cycle) | Turns the sun around the celestial pole; the stars follow it. |
| [`gusts`](#gusts) | Varies the wind strength and direction around the wind the world started with. |
| [`cloud_drift`](#cloud_drift) | Moves the clouds with the wind. |
| [`tides`](#tides) | Raises and lowers the water: turns the tide angle a full turn per period. Heights and phases come with the scene. |
| [`entity_motion`](#entity_motion) | Moves the entities that ask for it in their properties: bob (float and roll), spin (turn on an axis) and path (follow points). |

## `day_cycle`

Turns the sun around the celestial pole; the stars follow it.

- **Reads:** the environment
- **Writes:** `environment`

| Parameter | Type | Default | Meaning |
|---|---|---|---|
| `seconds_per_day` | number | `120` | how long a whole day lasts |

## `gusts`

Varies the wind strength and direction around the wind the world started with.

- **Reads:** the environment
- **Writes:** `environment`

| Parameter | Type | Default | Meaning |
|---|---|---|---|
| `variation` | number | `0.5` | how much the strength changes (0..1) |
| `veering` | number | `0.35` | how far the direction swings, in radians |

## `cloud_drift`

Moves the clouds with the wind.

- **Reads:** the environment
- **Writes:** `environment`

| Parameter | Type | Default | Meaning |
|---|---|---|---|
| `speed` | number | `0.05` | cloud-pattern units per second at wind strength 1 |

## `tides`

Raises and lowers the water: turns the tide angle a full turn per period. Heights and phases come with the scene.

- **Reads:** the environment
- **Writes:** `environment`

| Parameter | Type | Default | Meaning |
|---|---|---|---|
| `period_seconds` | number | `0` | seconds per tidal cycle (0 = 12.42 h in proportion to seconds_per_day) |
| `seconds_per_day` | number | `120` | day length used when period_seconds is 0 |

## `entity_motion`

Moves the entities that ask for it in their properties: bob (float and roll), spin (turn on an axis) and path (follow points).

- **Reads:** entities with property `bob`, `spin`, `path`
- **Writes:** nothing of its own (it moves the entities that carry `bob`, `spin` or `path`)

No parameters.
