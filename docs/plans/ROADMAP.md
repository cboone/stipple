# stipple Roadmap

This document defines v1 priorities based on Maze Wars as the primary consumer, plus a roadmap for later major releases.

## v1.0: Maze Wars Core

### Goals

- Fast, predictable braille rendering for a first person maze view.
- Lightweight UI overlays for HUD, minimap, and status text.
- Keep core packages dependency free.

### Non-goals

- Animation framework or plotting utilities.
- HTML5 Canvas style 2D context.
- Image import and export.
- 3D object system.
- Thread safety guarantees for concurrent access.

### Public API sketch

```go
package canvas

type Canvas struct {
    // internal state
}

type Option func(*Canvas)

func New(width, height int, options ...Option) *Canvas

func (c *Canvas) Clear()
func (c *Canvas) Frame() string
func (c *Canvas) Get(x, y float64) bool
func (c *Canvas) Set(x, y float64)
func (c *Canvas) Toggle(x, y float64)
func (c *Canvas) Unset(x, y float64)

func (c *Canvas) Height() int
func (c *Canvas) Width() int
func (c *Canvas) Rows() int
func (c *Canvas) Cols() int

func WithInvertedY() Option
func WithColor() Option
func WithText() Option
```

```go
package draw

func Line(c *canvas.Canvas, startX, startY, endX, endY float64)
func Rectangle(c *canvas.Canvas, x, y, width, height float64)
func RectangleFilled(c *canvas.Canvas, x, y, width, height float64)
```

```go
package text

type Font struct {
    // internal state
}

func DefaultFont() *Font
func Draw(c *canvas.Canvas, x, y float64, content string, font *Font)
```

### Package layout

```
stipple/
  canvas/
    braille.go
    canvas.go
    color.go
    options.go
  draw/
    line.go
    rectangle.go
  text/
    font.go
    text.go
  internal/
    term/
  examples/
```

### Implementation steps

1. Implement braille encoding and decoding helpers.
2. Implement Canvas with Set, Unset, Toggle, Get, Clear, and Frame.
3. Add size helpers (Width, Height, Rows, Cols).
4. Add WithInvertedY option.
5. Add optional 8 color ANSI output per cell.
6. Implement Bresenham line algorithm.
7. Implement Rectangle and RectangleFilled.
8. Implement basic bitmap text rendering with a built in font.
9. Add examples for a maze frame, HUD overlay, and minimap.
10. Add core tests and a small golden output test.

### Testing scope

- Unit tests for braille mapping and Canvas operations.
- Line algorithm tests for horizontal, vertical, and diagonal lines.
- Text rendering tests for a few characters.
- One golden output test for a maze frame with HUD text.

### Release criteria

- Public API documented with examples.
- Tests green on supported Go versions.
- Zero external dependencies in the core packages.

## Roadmap for Major Version Releases

### v2.0: Expanded Drawing and Fills

- Add circles, ellipses, arcs, and Bezier curves.
- Add polygon drawing and polygon fill.
- Add flood fill for closed shapes.
- Add thick line support for bold walls and highlights.

### v3.0: Turtle and Canvas 2D Context

- Add turtle graphics with standard commands.
- Add Canvas 2D style path API with transforms and state stack.
- Define fill rules and stroke options.

### v4.0: Animation and Plotting

- Add an animation loop and drawable interface.
- Add basic plotting utilities for lines and scatter plots.
- Add timing utilities suitable for terminal frame rates.

### v5.0: Image and 3D Utilities

- Add optional image import and export behind build tags.
- Evaluate minimal 3D wireframe projection support.

## Versioning notes

- SemVer with major releases allowed to introduce breaking API changes.
- Minor releases within a major version add features without breaking the API.
