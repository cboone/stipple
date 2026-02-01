# brodot v1 Implementation Plan

A focused braille graphics rendering library for the Maze Wars TUI game.

## Scope

**In scope for v1:**
- Core braille canvas with pixel-level control
- Line drawing (Bresenham algorithm)
- Rectangle drawing (outline and filled)
- Circle drawing (Bresenham midpoint algorithm) for eyeball sprites
- Optional per-cell ANSI color support

**Out of scope (handled by game layer):**
- Text rendering and fonts (HUD handled elsewhere)
- Animation framework
- Input handling
- 3D math/projection (game provides 2D coordinates)

---

## Package Structure

```
brodot/
├── canvas/
│   ├── braille.go      # Braille encoding constants and helpers
│   ├── canvas.go       # Canvas struct and core methods
│   ├── color.go        # ANSI color type and escape codes
│   └── options.go      # Functional options pattern
├── draw/
│   ├── circle.go       # Bresenham midpoint circle algorithm
│   ├── line.go         # Bresenham line algorithm
│   └── rectangle.go    # Rectangle outline and fill
├── internal/
│   └── term/
│       └── size.go     # Terminal size detection (optional utility)
└── examples/
    ├── demo/
    │   └── main.go     # Visual demo (grows with each version)
    └── maze/
        └── main.go     # Maze Wars-style rendering example
```

---

## Visual Testing Strategy

Two complementary approaches for visual feedback during development:

### 1. Demo Program (`examples/demo/main.go`)

A growing example program that exercises new features as they're added:

```bash
go run ./examples/demo/
```

**v0.1.0**: Draw individual pixels, show braille encoding
**v0.2.0**: Add line demonstrations (horizontal, vertical, diagonal)
**v0.3.0**: Add rectangle demonstrations (outline and filled)
**v0.4.0**: Add circle demonstrations (outline and filled)
**v0.5.0+**: Add color demonstrations

### 2. Test Visual Flag

Tests can print rendered output when run with `-visual`:

```bash
go test ./... -v -args -visual
```

Implementation in test files:

```go
var visualFlag = flag.Bool("visual", false, "print visual output")

func TestSomething(t *testing.T) {
    c := canvas.New(40, 20)
    // ... draw something ...

    if *visualFlag {
        fmt.Println("\n=== TestSomething ===")
        fmt.Println(c.Frame())
    }

    // ... assertions ...
}
```

Add `flag.Parse()` in `TestMain` for each test package:

```go
func TestMain(m *testing.M) {
    flag.Parse()
    os.Exit(m.Run())
}
```

---

## v0.1.0: Braille Canvas Foundation

### Goal
Establish the core canvas with braille encoding. A user can set individual pixels and render to a string.

### Files to Create

#### `canvas/braille.go`

```go
package canvas

// BrailleOffset is the Unicode code point for the empty braille pattern.
const BrailleOffset = 0x2800

// pixelMap maps (row, column) within a 4x2 cell to the braille dot bit.
// Rows 0-3, Columns 0-1.
var pixelMap = [4][2]rune{
    {0x01, 0x08}, // row 0: dots 1, 4
    {0x02, 0x10}, // row 1: dots 2, 5
    {0x04, 0x20}, // row 2: dots 3, 6
    {0x40, 0x80}, // row 3: dots 7, 8
}
```

#### `canvas/canvas.go`

Core struct and methods:

```go
package canvas

type Canvas struct {
    width    int       // pixel width
    height   int       // pixel height
    cells    [][]rune  // braille character grid [row][col]
    invertY  bool      // Y-axis direction
}

func New(width, height int, options ...Option) *Canvas
func (c *Canvas) Set(x, y float64)
func (c *Canvas) Unset(x, y float64)
func (c *Canvas) Toggle(x, y float64)
func (c *Canvas) Get(x, y float64) bool
func (c *Canvas) Clear()
func (c *Canvas) Frame() string
func (c *Canvas) Width() int
func (c *Canvas) Height() int
func (c *Canvas) Rows() int   // height / 4 (terminal rows)
func (c *Canvas) Cols() int   // width / 2 (terminal columns)
```

Implementation notes:
- `cells` is allocated as `[Rows()][Cols()]rune`, each initialized to `BrailleOffset`
- Pixel coordinates use `float64` for sub-pixel precision; convert to `int` internally
- Out-of-bounds coordinates are silently ignored (no error, no panic)
- `Frame()` joins rows with newlines, each row joins cells into a string

#### `canvas/options.go`

```go
package canvas

type Option func(*Canvas)

func WithInvertedY() Option {
    return func(c *Canvas) {
        c.invertY = true
    }
}
```

### Tests to Create

#### `canvas/braille_test.go`
- Test that each pixel position maps to the correct braille pattern
- Test all 8 dot positions individually
- Test full cell (all dots set) equals `\u28FF`

#### `canvas/canvas_test.go`
- `TestNew`: verify dimensions and initial state (all cells are `BrailleOffset`)
- `TestSetGet`: set a pixel, verify Get returns true
- `TestUnset`: set then unset, verify Get returns false
- `TestToggle`: toggle twice returns to original state
- `TestClear`: set pixels, clear, verify all Get return false
- `TestFrame`: set specific pixels, verify exact braille output
- `TestOutOfBounds`: setting out-of-bounds coordinates does not panic
- `TestInvertedY`: verify Y-axis inversion works correctly

### Deliverables
- [ ] `canvas/braille.go` with constants and pixel map
- [ ] `canvas/canvas.go` with Canvas struct and all core methods
- [ ] `canvas/options.go` with `WithInvertedY`
- [ ] `canvas/braille_test.go`
- [ ] `canvas/canvas_test.go` with `-visual` flag support
- [ ] `examples/demo/main.go` showing pixel operations
- [ ] `go.mod` initialized

### Verification
```bash
# Run tests
go test ./canvas/...

# Visual verification
go test ./canvas/... -v -args -visual
go run ./examples/demo/
```

---

## v0.2.0: Line Drawing

### Goal
Add Bresenham's line algorithm for drawing straight lines between two points.

### Files to Create

#### `draw/line.go`

```go
package draw

import "github.com/brodot/brodot/canvas"

// Line draws a line from (startX, startY) to (endX, endY) using Bresenham's algorithm.
func Line(c *canvas.Canvas, startX, startY, endX, endY float64)
```

Implementation notes:
- Use integer Bresenham algorithm (no floating point in inner loop)
- Convert float64 coordinates to int at function entry
- Handle all octants (steep/shallow, positive/negative slopes)
- Set each pixel along the line using `c.Set()`

### Tests to Create

#### `draw/line_test.go`
- `TestLineHorizontal`: verify pixels along a horizontal line
- `TestLineVertical`: verify pixels along a vertical line
- `TestLineDiagonalPositive`: 45-degree line with positive slope
- `TestLineDiagonalNegative`: 45-degree line with negative slope
- `TestLineSymmetry`: `Line(a, b, c, d)` produces same pixels as `Line(c, d, a, b)`
- `TestLineSinglePoint`: start equals end draws one pixel
- `TestLineShallowSlope`: slope < 1
- `TestLineSteepSlope`: slope > 1

### Deliverables
- [ ] `draw/line.go` with Bresenham implementation
- [ ] `draw/line_test.go` with `-visual` flag support
- [ ] Update `examples/demo/main.go` with line demonstrations

### Verification
```bash
go test ./draw/... -v -args -visual
go run ./examples/demo/
```

---

## v0.3.0: Rectangle Drawing

### Goal
Add rectangle drawing (outline and filled) for rendering wall faces.

### Files to Create

#### `draw/rectangle.go`

```go
package draw

import "github.com/brodot/brodot/canvas"

// Rectangle draws a rectangle outline from (x, y) with the given width and height.
func Rectangle(c *canvas.Canvas, x, y, width, height float64)

// RectangleFilled draws a filled rectangle from (x, y) with the given width and height.
func RectangleFilled(c *canvas.Canvas, x, y, width, height float64)
```

Implementation notes:
- `Rectangle` calls `Line` four times for the edges
- `RectangleFilled` iterates row by row, setting all pixels in each row
- Width and height of 0 or negative draw nothing
- Coordinates can be negative (partially off-canvas rectangles are clipped)

### Tests to Create

#### `draw/rectangle_test.go`
- `TestRectangle`: verify outline pixels are set, interior is not
- `TestRectangleFilled`: verify all interior pixels are set
- `TestRectangleZeroSize`: zero width or height draws nothing
- `TestRectanglePartiallyOffCanvas`: rectangle extending beyond bounds is clipped

### Deliverables
- [ ] `draw/rectangle.go`
- [ ] `draw/rectangle_test.go` with `-visual` flag support
- [ ] Update `examples/demo/main.go` with rectangle demonstrations

### Verification
```bash
go test ./draw/... -v -args -visual
go run ./examples/demo/
```

---

## v0.4.0: Circle Drawing

### Goal
Add Bresenham's midpoint circle algorithm for drawing circles (eyeball sprites).

### Files to Create

#### `draw/circle.go`

```go
package draw

import "github.com/brodot/brodot/canvas"

// Circle draws a circle outline centered at (centerX, centerY) with the given radius.
func Circle(c *canvas.Canvas, centerX, centerY, radius float64)

// CircleFilled draws a filled circle centered at (centerX, centerY) with the given radius.
func CircleFilled(c *canvas.Canvas, centerX, centerY, radius float64)
```

Implementation notes:
- Use Bresenham's midpoint circle algorithm (integer arithmetic in inner loop)
- Convert float64 coordinates to int at function entry
- Leverage 8-way symmetry: compute one octant, reflect to all 8
- `Circle` sets outline pixels only
- `CircleFilled` draws horizontal lines between symmetric points for each y level
- Radius of 0 draws a single pixel at center
- Negative radius draws nothing

### Tests to Create

#### `draw/circle_test.go`
- `TestCircleSymmetry`: verify 8-way symmetry (circle looks round, not skewed)
- `TestCircleRadius0`: radius 0 draws single pixel
- `TestCircleRadius1`: verify small circle pixels
- `TestCircleFilled`: verify interior pixels are set
- `TestCircleFilledNoGaps`: verify no gaps in filled circle (scan each row)
- `TestCircleOutlineOnly`: verify outline circle has empty interior

### Deliverables
- [ ] `draw/circle.go` with Bresenham midpoint implementation
- [ ] `draw/circle_test.go` with `-visual` flag support
- [ ] Update `examples/demo/main.go` with circle demonstrations

### Verification
```bash
go test ./draw/... -v -args -visual
go run ./examples/demo/
```

---

## v0.5.0: Color Support (Canvas)

### Goal
Add optional per-cell ANSI color support for visual distinction.

### Files to Modify/Create

#### `canvas/color.go`

```go
package canvas

type Color uint8

const (
    ColorDefault Color = iota
    ColorBlack
    ColorRed
    ColorGreen
    ColorYellow
    ColorBlue
    ColorMagenta
    ColorCyan
    ColorWhite
)

// ANSI returns the ANSI escape sequence for this color.
func (c Color) ANSI() string

// ANSIReset returns the reset escape sequence.
func ANSIReset() string
```

#### Modify `canvas/canvas.go`

Add color grid and methods:

```go
type Canvas struct {
    // ... existing fields
    colors      [][]Color  // per-cell color, nil if colors disabled
    colorEnabled bool
}

// SetColor sets a pixel and its cell color.
func (c *Canvas) SetColor(x, y float64, color Color)

// Frame() updated to include ANSI escape codes when colors are enabled
```

#### `canvas/options.go`

Add color option:

```go
func WithColor() Option {
    return func(c *Canvas) {
        c.colorEnabled = true
        c.colors = make([][]Color, c.Rows())
        for i := range c.colors {
            c.colors[i] = make([]Color, c.Cols())
        }
    }
}
```

### Tests to Create

#### `canvas/color_test.go`
- `TestColorANSI`: verify each color produces correct escape sequence
- `TestSetColor`: set colored pixel, verify Frame contains escape codes
- `TestColorDisabled`: without `WithColor()`, SetColor still sets pixel but no escape codes
- `TestColorReset`: verify colors reset between cells

### Deliverables
- [ ] `canvas/color.go`
- [ ] Updated `canvas/canvas.go` with color support
- [ ] Updated `canvas/options.go` with `WithColor()`
- [ ] `canvas/color_test.go` with `-visual` flag support
- [ ] Update `examples/demo/main.go` with color demonstrations

### Verification
```bash
go test ./canvas/... -v -args -visual
go run ./examples/demo/
```

---

## v0.6.0: Draw Package Color Support

### Goal
Extend draw functions to accept optional color parameters.

### Files to Modify

#### `draw/line.go`

Add color variant:

```go
// LineWithColor draws a colored line.
func LineWithColor(c *canvas.Canvas, startX, startY, endX, endY float64, color canvas.Color)
```

#### `draw/rectangle.go`

Add color variants:

```go
// RectangleWithColor draws a colored rectangle outline.
func RectangleWithColor(c *canvas.Canvas, x, y, width, height float64, color canvas.Color)

// RectangleFilledWithColor draws a colored filled rectangle.
func RectangleFilledWithColor(c *canvas.Canvas, x, y, width, height float64, color canvas.Color)
```

#### `draw/circle.go`

Add color variants:

```go
// CircleWithColor draws a colored circle outline.
func CircleWithColor(c *canvas.Canvas, centerX, centerY, radius float64, color canvas.Color)

// CircleFilledWithColor draws a colored filled circle.
func CircleFilledWithColor(c *canvas.Canvas, centerX, centerY, radius float64, color canvas.Color)
```

### Tests to Create
- `TestLineWithColor`: verify colored line output
- `TestRectangleWithColor`: verify colored rectangle output
- `TestCircleWithColor`: verify colored circle output

### Deliverables
- [ ] Updated `draw/line.go` with color variant
- [ ] Updated `draw/rectangle.go` with color variants
- [ ] Updated `draw/circle.go` with color variants
- [ ] Updated tests with `-visual` flag support
- [ ] Update `examples/demo/main.go` with colored shape demonstrations

### Verification
```bash
go test ./... -v -args -visual
go run ./examples/demo/
```

---

## v0.7.0: Example and Documentation

### Goal
Create a Maze Wars-style rendering example and documentation.

### Files to Create

#### `examples/maze/main.go`

Demonstrate:
- Creating a canvas sized to terminal
- Drawing a simple first-person corridor view with rectangles
- Using colors to distinguish wall segments at different depths
- Rendering an "eyeball" sprite using circles (outline and filled)

#### `README.md`

Document:
- Installation
- Quick start example
- API overview
- Link to examples

### Deliverables
- [ ] `examples/maze/main.go`
- [ ] `README.md` with documentation

### Verification
```bash
go run ./examples/maze/
```
Visual inspection of output.

---

## v1.0.0: Release Polish

### Goal
Final polish, golden tests, and release preparation.

### Tasks

1. **Golden Output Test**
   - Create `testdata/` directory with expected output files
   - Add golden test that renders a known scene and compares to expected output

2. **Code Review**
   - Ensure all public types and functions have doc comments
   - Verify error handling is consistent (silent ignore for out-of-bounds)
   - Check for any panics on edge cases

3. **CI Setup**
   - Add GitHub Actions workflow for `go test ./...`
   - Test on Go 1.21 and latest

4. **Release**
   - Tag v1.0.0
   - Ensure `go.mod` module path is correct for import

### Deliverables
- [ ] `canvas/golden_test.go` with golden output tests
- [ ] `testdata/*.golden` files
- [ ] `.github/workflows/test.yml`
- [ ] All doc comments complete
- [ ] v1.0.0 tag

### Verification
```bash
go test ./...
```
All tests pass, including golden tests.

---

## Summary

| Version | Focus                        | Key Deliverables                         |
|---------|------------------------------|------------------------------------------|
| v0.1.0  | Braille canvas foundation    | Canvas, Set/Get/Frame, options           |
| v0.2.0  | Line drawing                 | Bresenham line algorithm                 |
| v0.3.0  | Rectangle drawing            | Outline and filled rectangles            |
| v0.4.0  | Circle drawing               | Bresenham midpoint circle algorithm      |
| v0.5.0  | Color support (canvas)       | ANSI colors, per-cell coloring           |
| v0.6.0  | Colored draw functions       | Color variants for all draw functions    |
| v0.7.0  | Example and docs             | Maze example, README                     |
| v1.0.0  | Release polish               | Golden tests, CI, release                |

---

## Design Decisions

### Coordinate System
- Use `float64` for all public API coordinates
- Allows sub-pixel precision for smooth rendering
- Convert to `int` internally for pixel operations

### Error Handling
- Out-of-bounds coordinates are silently ignored
- No panics, no errors returned
- Simplifies calling code, enables partial rendering of off-canvas shapes

### Color Model
- Optional via `WithColor()` to avoid allocation when not needed
- 8 basic ANSI colors (sufficient for Maze Wars)
- Per-cell coloring (not per-pixel, matches braille character granularity)

### Dependencies
- Zero external dependencies in core packages
- Standard library only (`strings`, `math`)

---

## Braille Reference

Unicode Braille patterns (U+2800 to U+28FF) encode a 2x4 dot grid:

```
Dot positions:     Bit values:
  0  1               0x01  0x08
  2  3               0x02  0x10
  4  5               0x04  0x20
  6  7               0x40  0x80
```

Each terminal cell represents 2 pixels wide by 4 pixels tall, providing 8x resolution improvement over standard characters.

Conversion formulas:
- Cell column = pixel_x / 2
- Cell row = pixel_y / 4
- Dot column = pixel_x % 2
- Dot row = pixel_y % 4
- Braille char = BrailleOffset | pixelMap[dot_row][dot_column]
