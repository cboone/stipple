# brodot - Claude Code Instructions

A braille graphics rendering library in Go, designed for the Maze Wars TUI game.

## Project Overview

brodot provides pixel-level braille graphics rendering for terminal applications. Each terminal cell displays a 2x4 braille pattern, giving 8x resolution improvement over standard characters.

## Key Documentation

- `docs/plans/PLAN.md` - Detailed v1 implementation plan with milestones v0.1 through v1.0
- `docs/plans/ROADMAP.md` - Long-term vision through v5.0
- `docs/plans/INITIAL-PLAN.md` - Original planning discussion

## Code Style

### Go Conventions

- Use descriptive variable names (no abbreviations)
- Zero external dependencies in core packages (standard library only)
- All public types and functions require doc comments
- Use functional options pattern for configuration

### Naming

- Coordinates: use `x`, `y`, `startX`, `startY`, `endX`, `endY`, `centerX`, `centerY`
- Dimensions: use `width`, `height`, `radius`
- Cell vs pixel: "pixel" for braille dots, "cell" for terminal characters

### Error Handling

- Out-of-bounds coordinates are silently ignored (no panics, no errors)
- This is intentional - enables partial rendering of off-canvas shapes

## Package Structure

```
brodot/
├── canvas/          # Core braille canvas
│   ├── braille.go   # Braille encoding constants
│   ├── canvas.go    # Canvas struct and methods
│   ├── color.go     # ANSI color support
│   └── options.go   # Functional options
├── draw/            # Drawing primitives
│   ├── circle.go    # Bresenham midpoint circle
│   ├── line.go      # Bresenham line algorithm
│   └── rectangle.go # Rectangle outline and fill
└── examples/        # Demo programs
    ├── demo/        # Feature showcase
    └── maze/        # Maze Wars example
```

## Implementation Details

### Braille Encoding

Unicode braille patterns (U+2800 to U+28FF) encode a 2x4 dot grid:

```
Dot positions:     Bit values:
  0  3               0x01  0x08
  1  4               0x02  0x10
  2  5               0x04  0x20
  6  7               0x40  0x80
```

### Coordinate System

- All public API coordinates use `float64` for sub-pixel precision
- Convert to `int` internally using `math.Floor`
- `New(width, height)` expects pixel units
- `Rows()` = height / 4, `Cols()` = width / 2

### Color Model

- Optional per-cell ANSI colors (8 basic colors)
- Enabled via `WithColor()` option
- Per-cell, not per-pixel (matches braille character granularity)
- Last write wins when multiple pixels in same cell use different colors

## Testing

### Visual Testing

Tests support a `-visual` flag for visual output:

```bash
go test ./... -v -args -visual
```

### Running Tests

```bash
go test ./...
```

### Demo Program

```bash
go run ./examples/demo/
```

## Current Implementation Status

Check `docs/plans/PLAN.md` for the current milestone deliverables and their completion status.
