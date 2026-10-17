# Welcome to go-simulator-motion 👋

**The motion plugin makes a scene feel alive**: the sun crosses the sky, the wind gusts, clouds drift, the tide rises, boats bob and sails turn.

<p align="center"><img alt="Strips showing the sun, the wind and the tide changing over two minutes" src="images/hero.svg" width="720"></p>

It works while you watch: it moves the sun, the wind, the clouds and the tide, and anything in the scene that asks to move (*bob*, *spin*, or follow a *path*).

## Is it for me?

- ✅ You use [go-simulator](https://github.com/akrck02/go-simulator) and want **a scene that moves**.
- ✅ You want to **set it up by writing a rules file**: no programming needed.
- ✅ You are curious about **the science** behind it: the pages explain the model with pictures.

## Where to start

Pick the line that sounds like you.

| If you… | Read | Time |
|---|---|---|
| are **new** and want to see it work | 🏴‍☠️ [The quest](2-the-quest.md): a harbour that never keeps still, in 5 levels | 45 min |
| want to **run its example** | [Getting started](3-getting-started.md) | 10 min |
| want to know **what each motion does** | [The model](4-the-model.md) | 10 min |
| are **looking up** a parameter | [Rules reference](6-rules-reference.md) | — |

## All the pages, in reading order

The menu on the left (or the files in `docs/`) follows this order.

**Start**

- [Overview](1-overview.md): this page
- [The quest](2-the-quest.md)
- [Getting started](3-getting-started.md)

**Understand**

- [The model](4-the-model.md)
- [How it works](5-how-it-works.md)

**Reference**

- [Rules reference](6-rules-reference.md)

## 🚧 Not done yet

- Pictures of the 3D viewer itself. The repository has no example scene with entities that ask for
  a motion, so the pictures in these pages are 2D drawings of what the plugin answered, not
  screenshots (see [how the pictures were made](5-how-it-works.md)).
- The heights and phases of the tide, the sway of plants and the look of day, twilight and night
  are worked out by the renderer from what this plugin writes; they are not shown here.
- An entity takes a single motion: a boat cannot both bob and follow a path.

## Words you will meet

| Word | What it means |
|---|---|
| Terminal | The window where you type commands: *Terminal* on macOS and Linux, *PowerShell* on Windows. |
| Repository | A project's folder of files, kept on GitHub. `git clone <address>` copies it to your computer. |
| Plugin | A small program that adds one ability to another: here, one kind of simulation each. |
| Rules file | A text file (YAML) that says which plugins to use, with which numbers, for how long: no programming needed. |
| Field | A grid of numbers laid over the world: the height of every spot, its temperature, its rain. |
| Quest | The tutorial of these docs: short levels, each ending with something that works and an achievement. |

## The family

go-simulator-motion is one of a family of small projects that work together:

| Project | What it does |
|---|---|
| [go-renderer](https://github.com/akrck02/go-renderer) | draws 3D scenes and flat interfaces |
| [go-ui](https://github.com/akrck02/go-ui) | buttons, panels and whole applications |
| [go-ui-android](https://github.com/akrck02/go-ui-android) | Android apps that look like Android's own |
| [go-sound](https://github.com/akrck02/go-sound) | sound and music |
| [go-scene-builder](https://github.com/akrck02/go-scene-builder) | 3D scenes from a flat map |
| [go-simulator](https://github.com/akrck02/go-simulator) | simulations that change a scene |
| [go-simulator-terrain](https://github.com/akrck02/go-simulator-terrain) | ground shaped by rivers |
| [go-simulator-climate](https://github.com/akrck02/go-simulator-climate) | weather and biomes |
| [go-simulator-flora](https://github.com/akrck02/go-simulator-flora) | forests that grow |
| [go-simulator-urban](https://github.com/akrck02/go-simulator-urban) | towns that grow |
| **go-simulator-motion** (this one) | a scene that moves |
| [go-3ds](https://github.com/akrck02/go-3ds) | Go on the Nintendo 3DS |

> ⚠️ All of them are **educational experiments**, written with an AI assistant and pending review by their author: good for learning and playing, not for production.
