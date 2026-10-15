# 📚 go-simulator-motion documentation

| | Page | What is in it |
|---|---|---|
| 🚀 | [Getting started](getting-started.md) | The folders it needs, running the example, the keys, asking for motions in a scene, the tests. |
| 🌊 | [Motions](model.md) | What every rule moves (sun, wind, clouds, tide) and the properties entities use to ask for a motion (bob, spin, path), with pictures of each. |
| 🔍 | [How it works](how-it-works.md) | How the plugin fits go-simulator, how the rules are stepped over time, and how the pictures were made. |
| 📋 | [Rules reference](rules.md) | Every rule and parameter, generated from the plugin's own description. |

New here? Start with the [project README](../README.md).

## 🚧 Not done yet

- Pictures of the 3D viewer itself. The repository has no example scene with entities that ask for
  a motion, so the pictures in these pages are 2D drawings of what the plugin answered, not
  screenshots (see [how the pictures were made](how-it-works.md)).
- The heights and phases of the tide, the sway of plants and the look of day, twilight and night
  are worked out by the renderer from what this plugin writes; they are not shown here.
- An entity takes a single motion: a boat cannot both bob and follow a path.
