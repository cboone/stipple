# stipple

A braille graphics rendering library for Go, designed for terminal-based games and visualizations.

## Overview

stipple provides pixel-level graphics in the terminal using Unicode braille characters. Each terminal cell displays a 2x4 braille pattern, giving 8x resolution improvement over standard character graphics.

## Status

**In development** - Currently implementing v0.1.0 (Braille Canvas Foundation).

See [docs/plans/PLAN.md](docs/plans/PLAN.md) for the implementation roadmap.

## Features (Planned for v1.0)

- Core braille canvas with pixel-level control
- Line drawing (Bresenham algorithm)
- Rectangle drawing (outline and filled)
- Circle drawing (Bresenham midpoint algorithm)
- Optional per-cell ANSI color support
- Zero external dependencies

## Installation

```bash
go get github.com/cboone/stipple
```

## Quick Start

```go
package main

import (
    "fmt"

    "github.com/cboone/stipple/canvas"
    "github.com/cboone/stipple/draw"
)

func main() {
    // Create a 80x40 pixel canvas (40x10 terminal cells)
    c := canvas.New(80, 40)

    // Draw a line
    draw.Line(c, 0, 0, 79, 39)

    // Draw a rectangle
    draw.Rectangle(c, 10, 5, 20, 15)

    // Render to terminal
    fmt.Println(c.Frame())
}
```

## Documentation

- [Implementation Plan](docs/plans/PLAN.md) - Detailed v1 milestones
- [Roadmap](docs/plans/ROADMAP.md) - Long-term vision through v5.0
- [Showcase](docs/SHOWCASE.md) - Visual examples

## How It Works

Unicode braille patterns (U+2800 to U+28FF) encode a 2x4 dot grid per character:

```
Dot positions:     Bit values:
  0  3               0x01  0x08
  1  4               0x02  0x10
  2  5               0x04  0x20
  6  7               0x40  0x80
```

This means an 80-column terminal can display 160 horizontal pixels and a 24-row terminal can display 96 vertical pixels.

## License

[MIT](LICENSE)
