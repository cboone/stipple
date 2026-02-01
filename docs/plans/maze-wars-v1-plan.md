# brodot v1.0 Implementation Plan

A focused implementation plan for brodot v1.0, targeting the Maze Wars TUI game as the primary consumer.

## Overview

This plan synthesizes INITIAL-PLAN.md and ROADMAP.md into actionable minor version releases leading to v1.0. Each minor version is a releasable milestone with tested, documented functionality.

### Target Use Case: Maze Wars

Maze Wars requires:
- **First-person maze view**: Braille-rendered 3D perspective of corridors and walls
- **HUD overlay**: Score, health, ammo, and status text
- **Minimap**: Top-down view of maze layout
- **Status messages**: Game events and player info

### Design Principles

1. **Zero external dependencies** in core packages
2. **Functional options pattern** for flexible configuration
3. **float64 coordinates** for sub-pixel positioning and smooth rendering
4. **Graceful degradation**: out-of-bounds coordinates silently ignored
5. **Idiomatic Go**: clear interfaces, proper error handling where needed

---

## Package Structure

```
brodot/
├── go.mod
├── canvas/
│   ├── braille.go      # Braille encoding/decoding
│   ├── canvas.go       # Canvas struct and core methods
│   ├── color.go        # ANSI color support
│   └── options.go      # Functional options
├── draw/
│   ├── line.go         # Bresenham line algorithm
│   └── rectangle.go    # Rectangle stroke and fill
├── text/
│   ├── font.go         # Bitmap font definition
│   └── text.go         # Text rendering
├── internal/
│   └── term/
│       └── size.go     # Terminal size detection
└── examples/
    ├── basic/          # Basic canvas operations
    ├── maze/           # Maze frame rendering
    └── hud/            # HUD overlay demo
```

---

## v0.1.0: Core Canvas Foundation

### Goal
Establish the foundational braille canvas with basic pixel operations.

### Files to Create

#### `go.mod`
```go
module github.com/brodot/brodot

go 1.21
```

#### `canvas/braille.go`

Braille encoding utilities:

```go
package canvas

// BrailleOffset is the Unicode code point for the empty braille character
const BrailleOffset = 0x2800

// pixelMap maps (row, col) within a 2x4 cell to braille bit values
// Row 0-3 (top to bottom), Col 0-1 (left to right)
var pixelMap = [4][2]rune{
    {0x01, 0x08}, // row 0: dots 1, 4
    {0x02, 0x10}, // row 1: dots 2, 5
    {0x04, 0x20}, // row 2: dots 3, 6
    {0x40, 0x80}, // row 3: dots 7, 8
}

// PixelToBrailleBit returns the braille bit for a pixel position within a cell
func PixelToBrailleBit(pixelX, pixelY int) rune {
    return pixelMap[pixelY%4][pixelX%2]
}
```

#### `canvas/canvas.go`

Core canvas implementation:

```go
package canvas

import (
    "strings"
)

// Canvas represents a braille-based drawing surface
type Canvas struct {
    width   int         // pixel width
    height  int         // pixel height
    cells   [][]rune    // braille character grid [row][col]
    invertY bool        // if true, Y increases downward
}

// New creates a new Canvas with the given pixel dimensions
func New(width, height int, options ...Option) *Canvas {
    // Ensure minimum size of one cell (2x4 pixels)
    if width < 2 {
        width = 2
    }
    if height < 4 {
        height = 4
    }

    canvas := &Canvas{
        width:  width,
        height: height,
    }

    // Apply options
    for _, option := range options {
        option(canvas)
    }

    // Initialize cell grid
    canvas.initCells()

    return canvas
}

func (canvas *Canvas) initCells() {
    rows := canvas.Rows()
    cols := canvas.Cols()
    canvas.cells = make([][]rune, rows)
    for row := range canvas.cells {
        canvas.cells[row] = make([]rune, cols)
        for col := range canvas.cells[row] {
            canvas.cells[row][col] = BrailleOffset
        }
    }
}

// Width returns the pixel width of the canvas
func (canvas *Canvas) Width() int {
    return canvas.width
}

// Height returns the pixel height of the canvas
func (canvas *Canvas) Height() int {
    return canvas.height
}

// Cols returns the number of terminal columns (width / 2)
func (canvas *Canvas) Cols() int {
    return (canvas.width + 1) / 2
}

// Rows returns the number of terminal rows (height / 4)
func (canvas *Canvas) Rows() int {
    return (canvas.height + 3) / 4
}

// Set turns on the pixel at (x, y)
func (canvas *Canvas) Set(x, y float64) {
    pixelX, pixelY := canvas.toPixel(x, y)
    if !canvas.inBounds(pixelX, pixelY) {
        return
    }

    cellX, cellY := pixelX/2, pixelY/4
    canvas.cells[cellY][cellX] |= pixelMap[pixelY%4][pixelX%2]
}

// Unset turns off the pixel at (x, y)
func (canvas *Canvas) Unset(x, y float64) {
    pixelX, pixelY := canvas.toPixel(x, y)
    if !canvas.inBounds(pixelX, pixelY) {
        return
    }

    cellX, cellY := pixelX/2, pixelY/4
    canvas.cells[cellY][cellX] &^= pixelMap[pixelY%4][pixelX%2]
}

// Toggle flips the pixel at (x, y)
func (canvas *Canvas) Toggle(x, y float64) {
    pixelX, pixelY := canvas.toPixel(x, y)
    if !canvas.inBounds(pixelX, pixelY) {
        return
    }

    cellX, cellY := pixelX/2, pixelY/4
    canvas.cells[cellY][cellX] ^= pixelMap[pixelY%4][pixelX%2]
}

// Get returns true if the pixel at (x, y) is set
func (canvas *Canvas) Get(x, y float64) bool {
    pixelX, pixelY := canvas.toPixel(x, y)
    if !canvas.inBounds(pixelX, pixelY) {
        return false
    }

    cellX, cellY := pixelX/2, pixelY/4
    bit := pixelMap[pixelY%4][pixelX%2]
    return (canvas.cells[cellY][cellX] & bit) != 0
}

// Clear resets all pixels to off
func (canvas *Canvas) Clear() {
    for row := range canvas.cells {
        for col := range canvas.cells[row] {
            canvas.cells[row][col] = BrailleOffset
        }
    }
}

// Frame returns the canvas as a string with newlines between rows
func (canvas *Canvas) Frame() string {
    var builder strings.Builder
    for rowIndex, row := range canvas.cells {
        for _, cell := range row {
            builder.WriteRune(cell)
        }
        if rowIndex < len(canvas.cells)-1 {
            builder.WriteByte('\n')
        }
    }
    return builder.String()
}

// toPixel converts float coordinates to integer pixel coordinates
func (canvas *Canvas) toPixel(x, y float64) (int, int) {
    pixelX := int(x)
    pixelY := int(y)

    if canvas.invertY {
        pixelY = canvas.height - 1 - pixelY
    }

    return pixelX, pixelY
}

// inBounds checks if pixel coordinates are within the canvas
func (canvas *Canvas) inBounds(pixelX, pixelY int) bool {
    return pixelX >= 0 && pixelX < canvas.width && pixelY >= 0 && pixelY < canvas.height
}
```

#### `canvas/options.go`

Functional options:

```go
package canvas

// Option configures a Canvas
type Option func(*Canvas)

// WithInvertedY makes Y coordinates increase downward (standard screen coordinates)
// By default, Y increases upward (mathematical coordinates)
func WithInvertedY() Option {
    return func(canvas *Canvas) {
        canvas.invertY = true
    }
}
```

### Tests for v0.1.0

#### `canvas/braille_test.go`

```go
package canvas

import "testing"

func TestPixelToBrailleBit(t *testing.T) {
    tests := []struct {
        name         string
        pixelX       int
        pixelY       int
        expectedBit  rune
    }{
        {"top-left dot 1", 0, 0, 0x01},
        {"top-right dot 4", 1, 0, 0x08},
        {"row 1 left dot 2", 0, 1, 0x02},
        {"row 1 right dot 5", 1, 1, 0x10},
        {"row 2 left dot 3", 0, 2, 0x04},
        {"row 2 right dot 6", 1, 2, 0x20},
        {"bottom-left dot 7", 0, 3, 0x40},
        {"bottom-right dot 8", 1, 3, 0x80},
    }

    for _, testCase := range tests {
        t.Run(testCase.name, func(t *testing.T) {
            got := PixelToBrailleBit(testCase.pixelX, testCase.pixelY)
            if got != testCase.expectedBit {
                t.Errorf("PixelToBrailleBit(%d, %d) = %#x, want %#x",
                    testCase.pixelX, testCase.pixelY, got, testCase.expectedBit)
            }
        })
    }
}

func TestFullBrailleCharacter(t *testing.T) {
    // All 8 dots set should produce U+28FF
    var fullChar rune = BrailleOffset
    for row := 0; row < 4; row++ {
        for col := 0; col < 2; col++ {
            fullChar |= pixelMap[row][col]
        }
    }

    if fullChar != '\u28FF' {
        t.Errorf("full braille character = %#x, want 0x28FF", fullChar)
    }
}
```

#### `canvas/canvas_test.go`

```go
package canvas

import "testing"

func TestNewCanvas(t *testing.T) {
    canvas := New(80, 40)

    if canvas.Width() != 80 {
        t.Errorf("Width() = %d, want 80", canvas.Width())
    }
    if canvas.Height() != 40 {
        t.Errorf("Height() = %d, want 40", canvas.Height())
    }
    if canvas.Cols() != 40 {
        t.Errorf("Cols() = %d, want 40", canvas.Cols())
    }
    if canvas.Rows() != 10 {
        t.Errorf("Rows() = %d, want 10", canvas.Rows())
    }
}

func TestMinimumCanvasSize(t *testing.T) {
    canvas := New(0, 0)

    if canvas.Width() < 2 {
        t.Errorf("Width() = %d, want >= 2", canvas.Width())
    }
    if canvas.Height() < 4 {
        t.Errorf("Height() = %d, want >= 4", canvas.Height())
    }
}

func TestSetAndGet(t *testing.T) {
    canvas := New(10, 10)

    // Initially unset
    if canvas.Get(5, 5) {
        t.Error("pixel should be unset initially")
    }

    // Set and verify
    canvas.Set(5, 5)
    if !canvas.Get(5, 5) {
        t.Error("pixel should be set after Set()")
    }

    // Other pixels should remain unset
    if canvas.Get(0, 0) {
        t.Error("other pixels should remain unset")
    }
}

func TestUnset(t *testing.T) {
    canvas := New(10, 10)

    canvas.Set(3, 3)
    canvas.Unset(3, 3)

    if canvas.Get(3, 3) {
        t.Error("pixel should be unset after Unset()")
    }
}

func TestToggle(t *testing.T) {
    canvas := New(10, 10)

    canvas.Toggle(2, 2)
    if !canvas.Get(2, 2) {
        t.Error("pixel should be set after first Toggle()")
    }

    canvas.Toggle(2, 2)
    if canvas.Get(2, 2) {
        t.Error("pixel should be unset after second Toggle()")
    }
}

func TestClear(t *testing.T) {
    canvas := New(10, 10)

    canvas.Set(0, 0)
    canvas.Set(5, 5)
    canvas.Set(9, 9)
    canvas.Clear()

    if canvas.Get(0, 0) || canvas.Get(5, 5) || canvas.Get(9, 9) {
        t.Error("all pixels should be unset after Clear()")
    }
}

func TestOutOfBoundsIgnored(t *testing.T) {
    canvas := New(10, 10)

    // These should not panic
    canvas.Set(-1, 0)
    canvas.Set(0, -1)
    canvas.Set(100, 0)
    canvas.Set(0, 100)

    // Out of bounds Get returns false
    if canvas.Get(-1, 0) {
        t.Error("out of bounds Get should return false")
    }
}

func TestFrameSingleCell(t *testing.T) {
    canvas := New(2, 4) // Single cell

    // Set top-left and bottom-right dots
    canvas.Set(0, 0) // dot 1
    canvas.Set(1, 3) // dot 8

    frame := canvas.Frame()
    expected := "\u2881" // dots 1 and 8

    if frame != expected {
        t.Errorf("Frame() = %q (%#x), want %q (%#x)",
            frame, []rune(frame)[0], expected, []rune(expected)[0])
    }
}

func TestInvertedY(t *testing.T) {
    canvas := New(2, 8, WithInvertedY())

    // With inverted Y, y=0 is at the top
    canvas.Set(0, 0) // Should map to top-left

    // In standard mode, y=0 would be bottom
    // With inverted Y, y=0 is top, so this should set the first row's first dot
    frame := canvas.Frame()

    // First character should have top-left dot set
    if len(frame) == 0 || ([]rune(frame)[0]&0x01) == 0 {
        t.Error("inverted Y should map y=0 to top row")
    }
}
```

### Acceptance Criteria for v0.1.0

- [ ] `go mod init` creates valid module
- [ ] `go build ./...` succeeds with no errors
- [ ] `go test ./...` passes all tests
- [ ] Canvas can be created with arbitrary dimensions
- [ ] Set, Unset, Toggle, Get work correctly
- [ ] Frame() produces valid braille Unicode string
- [ ] Out-of-bounds coordinates are silently ignored
- [ ] WithInvertedY option works correctly

---

## v0.2.0: Line and Rectangle Drawing

### Goal
Add drawing primitives needed for maze walls and HUD elements.

### Files to Create

#### `draw/line.go`

Bresenham's line algorithm:

```go
package draw

import "github.com/brodot/brodot/canvas"

// Line draws a line from (startX, startY) to (endX, endY)
func Line(c *canvas.Canvas, startX, startY, endX, endY float64) {
    x0 := int(startX)
    y0 := int(startY)
    x1 := int(endX)
    y1 := int(endY)

    dx := abs(x1 - x0)
    dy := -abs(y1 - y0)
    sx := sign(x1 - x0)
    sy := sign(y1 - y0)
    err := dx + dy

    for {
        c.Set(float64(x0), float64(y0))

        if x0 == x1 && y0 == y1 {
            break
        }

        e2 := 2 * err

        if e2 >= dy {
            err += dy
            x0 += sx
        }

        if e2 <= dx {
            err += dx
            y0 += sy
        }
    }
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func sign(x int) int {
    if x > 0 {
        return 1
    }
    if x < 0 {
        return -1
    }
    return 0
}
```

#### `draw/rectangle.go`

Rectangle drawing:

```go
package draw

import "github.com/brodot/brodot/canvas"

// Rectangle draws the outline of a rectangle
func Rectangle(c *canvas.Canvas, x, y, width, height float64) {
    x2 := x + width - 1
    y2 := y + height - 1

    // Four sides
    Line(c, x, y, x2, y)     // top
    Line(c, x, y2, x2, y2)   // bottom
    Line(c, x, y, x, y2)     // left
    Line(c, x2, y, x2, y2)   // right
}

// RectangleFilled draws a filled rectangle
func RectangleFilled(c *canvas.Canvas, x, y, width, height float64) {
    x1 := int(x)
    y1 := int(y)
    x2 := int(x + width)
    y2 := int(y + height)

    for py := y1; py < y2; py++ {
        for px := x1; px < x2; px++ {
            c.Set(float64(px), float64(py))
        }
    }
}
```

### Tests for v0.2.0

#### `draw/line_test.go`

```go
package draw

import (
    "testing"

    "github.com/brodot/brodot/canvas"
)

func TestLineHorizontal(t *testing.T) {
    c := canvas.New(20, 10)

    Line(c, 0, 5, 19, 5)

    // Check all points on the line are set
    for x := 0; x < 20; x++ {
        if !c.Get(float64(x), 5) {
            t.Errorf("pixel at (%d, 5) should be set", x)
        }
    }

    // Check points off the line are not set
    if c.Get(0, 0) {
        t.Error("pixel at (0, 0) should not be set")
    }
}

func TestLineVertical(t *testing.T) {
    c := canvas.New(10, 20)

    Line(c, 5, 0, 5, 19)

    for y := 0; y < 20; y++ {
        if !c.Get(5, float64(y)) {
            t.Errorf("pixel at (5, %d) should be set", y)
        }
    }
}

func TestLineDiagonal(t *testing.T) {
    c := canvas.New(10, 10)

    Line(c, 0, 0, 9, 9)

    // Diagonal should have no gaps
    for i := 0; i < 10; i++ {
        if !c.Get(float64(i), float64(i)) {
            t.Errorf("pixel at (%d, %d) should be set", i, i)
        }
    }
}

func TestLineSymmetry(t *testing.T) {
    c1 := canvas.New(20, 20)
    c2 := canvas.New(20, 20)

    Line(c1, 3, 7, 15, 12)
    Line(c2, 15, 12, 3, 7)

    // Both canvases should have the same pixels set
    for y := 0; y < 20; y++ {
        for x := 0; x < 20; x++ {
            if c1.Get(float64(x), float64(y)) != c2.Get(float64(x), float64(y)) {
                t.Errorf("line not symmetric at (%d, %d)", x, y)
            }
        }
    }
}

func TestLineSinglePoint(t *testing.T) {
    c := canvas.New(10, 10)

    Line(c, 5, 5, 5, 5)

    if !c.Get(5, 5) {
        t.Error("single point line should set the point")
    }
}
```

#### `draw/rectangle_test.go`

```go
package draw

import (
    "testing"

    "github.com/brodot/brodot/canvas"
)

func TestRectangle(t *testing.T) {
    c := canvas.New(20, 20)

    Rectangle(c, 5, 5, 10, 8)

    // Check corners are set
    corners := [][2]float64{
        {5, 5}, {14, 5}, {5, 12}, {14, 12},
    }
    for _, corner := range corners {
        if !c.Get(corner[0], corner[1]) {
            t.Errorf("corner at (%.0f, %.0f) should be set", corner[0], corner[1])
        }
    }

    // Check middle of rectangle is not set (outline only)
    if c.Get(9, 8) {
        t.Error("interior point should not be set for outline rectangle")
    }
}

func TestRectangleFilled(t *testing.T) {
    c := canvas.New(20, 20)

    RectangleFilled(c, 5, 5, 10, 8)

    // Check interior is filled
    for y := 5; y < 13; y++ {
        for x := 5; x < 15; x++ {
            if !c.Get(float64(x), float64(y)) {
                t.Errorf("pixel at (%d, %d) should be set in filled rectangle", x, y)
            }
        }
    }

    // Check outside is not set
    if c.Get(4, 5) || c.Get(5, 4) {
        t.Error("pixels outside rectangle should not be set")
    }
}
```

### Acceptance Criteria for v0.2.0

- [ ] Line draws correctly in all directions (horizontal, vertical, diagonal)
- [ ] Line is symmetric (same pixels regardless of direction)
- [ ] Rectangle draws proper outline
- [ ] RectangleFilled fills entire interior
- [ ] All drawing functions handle edge cases (single point, zero size)

---

## v0.3.0: Color Support

### Goal
Add 8-color ANSI support for differentiating maze elements (walls, floor, enemies, HUD).

### Files to Create/Modify

#### `canvas/color.go`

```go
package canvas

import "fmt"

// Color represents a terminal color
type Color uint8

// Standard 8 ANSI colors
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

// ANSI returns the ANSI escape code for the foreground color
func (color Color) ANSI() string {
    if color == ColorDefault {
        return "\033[39m"
    }
    return fmt.Sprintf("\033[%dm", 30+color-1)
}

// ANSIBright returns the ANSI escape code for the bright foreground color
func (color Color) ANSIBright() string {
    if color == ColorDefault {
        return "\033[39m"
    }
    return fmt.Sprintf("\033[%dm", 90+color-1)
}

// ANSIReset returns the code to reset all attributes
func ANSIReset() string {
    return "\033[0m"
}
```

#### `canvas/canvas.go` (modifications)

Add color support to the Canvas struct:

```go
// Add to Canvas struct:
type Canvas struct {
    width   int
    height  int
    cells   [][]rune
    colors  [][]Color  // per-cell colors (nil if color disabled)
    invertY bool
}

// Add SetColor method:
func (canvas *Canvas) SetColor(x, y float64, color Color) {
    if canvas.colors == nil {
        return
    }

    pixelX, pixelY := canvas.toPixel(x, y)
    if !canvas.inBounds(pixelX, pixelY) {
        return
    }

    cellX, cellY := pixelX/2, pixelY/4
    canvas.colors[cellY][cellX] = color
}

// Modify Frame() to include color codes:
func (canvas *Canvas) Frame() string {
    var builder strings.Builder

    for rowIndex, row := range canvas.cells {
        if canvas.colors != nil {
            currentColor := ColorDefault
            for colIndex, cell := range row {
                cellColor := canvas.colors[rowIndex][colIndex]
                if cellColor != currentColor {
                    builder.WriteString(cellColor.ANSI())
                    currentColor = cellColor
                }
                builder.WriteRune(cell)
            }
            if currentColor != ColorDefault {
                builder.WriteString(ANSIReset())
            }
        } else {
            for _, cell := range row {
                builder.WriteRune(cell)
            }
        }

        if rowIndex < len(canvas.cells)-1 {
            builder.WriteByte('\n')
        }
    }

    return builder.String()
}
```

#### `canvas/options.go` (add color option)

```go
// WithColor enables per-cell color support
func WithColor() Option {
    return func(canvas *Canvas) {
        // Mark that colors should be initialized
        // Actual initialization happens in initCells
    }
}
```

#### `draw/line.go` and `draw/rectangle.go` (add color variants)

```go
// LineWithColor draws a colored line
func LineWithColor(c *canvas.Canvas, startX, startY, endX, endY float64, color canvas.Color) {
    // Same algorithm but also set color
    x0, y0, x1, y1 := int(startX), int(startY), int(endX), int(endY)
    // ... Bresenham loop with c.SetColor(x, y, color) added
}
```

### Tests for v0.3.0

#### `canvas/color_test.go`

```go
package canvas

import "testing"

func TestColorANSI(t *testing.T) {
    tests := []struct {
        color    Color
        expected string
    }{
        {ColorDefault, "\033[39m"},
        {ColorRed, "\033[31m"},
        {ColorGreen, "\033[32m"},
        {ColorBlue, "\033[34m"},
    }

    for _, testCase := range tests {
        got := testCase.color.ANSI()
        if got != testCase.expected {
            t.Errorf("Color(%d).ANSI() = %q, want %q", testCase.color, got, testCase.expected)
        }
    }
}

func TestCanvasWithColor(t *testing.T) {
    c := New(10, 10, WithColor())

    c.Set(0, 0)
    c.SetColor(0, 0, ColorRed)

    frame := c.Frame()

    // Frame should contain ANSI escape sequences
    if len(frame) < 10 {
        t.Error("colored frame should contain ANSI codes")
    }
}
```

### Acceptance Criteria for v0.3.0

- [ ] 8 standard colors available as constants
- [ ] Color.ANSI() returns correct escape codes
- [ ] WithColor() option enables per-cell coloring
- [ ] SetColor() sets color for the cell containing (x, y)
- [ ] Frame() includes ANSI codes when colors enabled
- [ ] Frame() resets color at end of each line

---

## v0.4.0: Text Rendering

### Goal
Add bitmap text rendering for HUD text, score display, and status messages.

### Files to Create

#### `text/font.go`

Define a simple 4x6 bitmap font (fits well in braille cells):

```go
package text

// Font represents a bitmap font
type Font struct {
    Width  int
    Height int
    Glyphs map[rune][]uint8 // bitmap data per character
}

// DefaultFont returns the built-in 4x6 font
func DefaultFont() *Font {
    return &defaultFont
}

var defaultFont = Font{
    Width:  4,
    Height: 6,
    Glyphs: map[rune][]uint8{
        'A': {
            0b0110,
            0b1001,
            0b1111,
            0b1001,
            0b1001,
            0b0000,
        },
        'B': {
            0b1110,
            0b1001,
            0b1110,
            0b1001,
            0b1110,
            0b0000,
        },
        // ... complete alphabet, numbers, punctuation
        '0': {
            0b0110,
            0b1001,
            0b1001,
            0b1001,
            0b0110,
            0b0000,
        },
        // ... etc
    },
}
```

#### `text/text.go`

```go
package text

import "github.com/brodot/brodot/canvas"

// Draw renders text at the given position
func Draw(c *canvas.Canvas, x, y float64, content string, font *Font) {
    if font == nil {
        font = DefaultFont()
    }

    cursorX := int(x)
    cursorY := int(y)

    for _, char := range content {
        glyph, ok := font.Glyphs[char]
        if !ok {
            // Skip unknown characters, advance cursor
            cursorX += font.Width + 1
            continue
        }

        // Draw glyph bitmap
        for row, bits := range glyph {
            for col := 0; col < font.Width; col++ {
                if (bits>>(font.Width-1-col))&1 == 1 {
                    c.Set(float64(cursorX+col), float64(cursorY+row))
                }
            }
        }

        cursorX += font.Width + 1 // 1 pixel spacing
    }
}

// DrawWithColor renders colored text
func DrawWithColor(c *canvas.Canvas, x, y float64, content string, font *Font, color canvas.Color) {
    // Similar to Draw but also sets color
}
```

### Tests for v0.4.0

#### `text/text_test.go`

```go
package text

import (
    "testing"

    "github.com/brodot/brodot/canvas"
)

func TestDrawSingleCharacter(t *testing.T) {
    c := canvas.New(10, 10)
    font := DefaultFont()

    Draw(c, 0, 0, "A", font)

    // Check that some pixels are set (glyph was drawn)
    pixelsSet := 0
    for y := 0; y < 10; y++ {
        for x := 0; x < 10; x++ {
            if c.Get(float64(x), float64(y)) {
                pixelsSet++
            }
        }
    }

    if pixelsSet == 0 {
        t.Error("Draw should set some pixels")
    }
}

func TestDrawMultipleCharacters(t *testing.T) {
    c := canvas.New(50, 10)
    font := DefaultFont()

    Draw(c, 0, 0, "AB", font)

    // Second character should be offset from first
    // Check that there are pixels in both character regions
}
```

### Acceptance Criteria for v0.4.0

- [ ] DefaultFont() returns a usable bitmap font
- [ ] Draw() renders text at specified position
- [ ] Unknown characters are skipped gracefully
- [ ] Character spacing is consistent
- [ ] DrawWithColor() colors the rendered text

---

## v0.5.0: Terminal Utilities and Examples

### Goal
Add terminal size detection and create example programs demonstrating Maze Wars use cases.

### Files to Create

#### `internal/term/size.go`

```go
package term

import (
    "os"

    "golang.org/x/term"
)

// Size returns the terminal dimensions in columns and rows
func Size() (cols, rows int, err error) {
    return term.GetSize(int(os.Stdout.Fd()))
}

// PixelSize returns terminal dimensions in braille pixels
func PixelSize() (width, height int, err error) {
    cols, rows, err := Size()
    if err != nil {
        return 0, 0, err
    }
    return cols * 2, rows * 4, nil
}
```

Note: This is the one allowed external dependency (golang.org/x/term), kept in internal.

#### `canvas/options.go` (add terminal size option)

```go
// WithTerminalSize sets canvas dimensions to match terminal
func WithTerminalSize() Option {
    return func(canvas *Canvas) {
        width, height, err := term.PixelSize()
        if err != nil {
            return // Keep default size on error
        }
        canvas.width = width
        canvas.height = height
    }
}
```

#### `examples/basic/main.go`

```go
package main

import (
    "fmt"

    "github.com/brodot/brodot/canvas"
    "github.com/brodot/brodot/draw"
)

func main() {
    c := canvas.New(80, 40)

    // Draw a box
    draw.Rectangle(c, 5, 5, 70, 30)

    // Draw diagonal lines
    draw.Line(c, 5, 5, 74, 34)
    draw.Line(c, 74, 5, 5, 34)

    fmt.Println(c.Frame())
}
```

#### `examples/maze/main.go`

First-person maze frame rendering example:

```go
package main

import (
    "fmt"
    "math"

    "github.com/brodot/brodot/canvas"
    "github.com/brodot/brodot/draw"
)

func main() {
    c := canvas.New(160, 80, canvas.WithInvertedY())

    // Simulate a simple corridor view
    drawCorridor(c, 0.5) // 0.5 = halfway down corridor

    fmt.Println(c.Frame())
}

func drawCorridor(c *canvas.Canvas, depth float64) {
    width := float64(c.Width())
    height := float64(c.Height())

    // Calculate perspective vanishing point
    vanishX := width / 2
    vanishY := height / 2

    // Calculate wall edges based on depth
    perspective := 0.2 + (0.8 * depth)

    leftNear := width * 0.1
    leftFar := vanishX - (vanishX-leftNear)*perspective
    rightNear := width * 0.9
    rightFar := vanishX + (rightNear-vanishX)*perspective

    topNear := height * 0.1
    topFar := vanishY - (vanishY-topNear)*perspective
    bottomNear := height * 0.9
    bottomFar := vanishY + (bottomNear-vanishY)*perspective

    // Draw corridor walls
    // Left wall
    draw.Line(c, leftNear, topNear, leftFar, topFar)
    draw.Line(c, leftNear, bottomNear, leftFar, bottomFar)
    draw.Line(c, leftNear, topNear, leftNear, bottomNear)

    // Right wall
    draw.Line(c, rightNear, topNear, rightFar, topFar)
    draw.Line(c, rightNear, bottomNear, rightFar, bottomFar)
    draw.Line(c, rightNear, topNear, rightNear, bottomNear)

    // Back wall
    draw.Line(c, leftFar, topFar, rightFar, topFar)
    draw.Line(c, leftFar, bottomFar, rightFar, bottomFar)
    draw.Line(c, leftFar, topFar, leftFar, bottomFar)
    draw.Line(c, rightFar, topFar, rightFar, bottomFar)
}
```

#### `examples/hud/main.go`

HUD overlay example:

```go
package main

import (
    "fmt"

    "github.com/brodot/brodot/canvas"
    "github.com/brodot/brodot/draw"
    "github.com/brodot/brodot/text"
)

func main() {
    c := canvas.New(160, 80, canvas.WithColor(), canvas.WithInvertedY())

    // Draw HUD background
    draw.RectangleFilled(c, 0, 0, 160, 12)

    // Draw score
    text.DrawWithColor(c, 5, 2, "SCORE: 1250", nil, canvas.ColorYellow)

    // Draw health bar outline
    text.DrawWithColor(c, 100, 2, "HEALTH:", nil, canvas.ColorWhite)
    draw.Rectangle(c, 140, 2, 18, 8)

    fmt.Println(c.Frame())
}
```

### Acceptance Criteria for v0.5.0

- [ ] term.Size() returns correct terminal dimensions
- [ ] WithTerminalSize() option works
- [ ] examples/basic runs and produces output
- [ ] examples/maze demonstrates corridor rendering
- [ ] examples/hud demonstrates text and color overlay
- [ ] All examples can be run with `go run`

---

## v1.0.0: Polish and Documentation

### Goal
Finalize API, add comprehensive documentation, and create release.

### Tasks

1. **API Review**
   - Ensure all public functions have doc comments
   - Review naming consistency
   - Add any missing convenience methods

2. **Documentation**
   - Write README.md with quick start guide
   - Add package-level doc comments
   - Create godoc examples for all major functions

3. **Golden File Test**
   - Create golden output test for maze frame with HUD

4. **CI/CD**
   - Add GitHub Actions workflow for testing
   - Test on multiple Go versions (1.21, 1.22)

5. **Release**
   - Tag v1.0.0
   - Create GitHub release with changelog

### Golden File Test

#### `canvas/golden_test.go`

```go
package canvas_test

import (
    "os"
    "path/filepath"
    "testing"

    "github.com/brodot/brodot/canvas"
    "github.com/brodot/brodot/draw"
    "github.com/brodot/brodot/text"
)

func TestGoldenMazeFrame(t *testing.T) {
    c := canvas.New(80, 40, canvas.WithInvertedY())

    // Draw a simple maze frame
    draw.Rectangle(c, 5, 5, 70, 30)
    draw.Line(c, 5, 5, 40, 20)
    draw.Line(c, 75, 5, 40, 20)

    // Add some text
    text.Draw(c, 10, 32, "MAZE WARS", nil)

    got := c.Frame()

    goldenPath := filepath.Join("testdata", "maze_frame.golden")

    if os.Getenv("UPDATE_GOLDEN") != "" {
        os.MkdirAll("testdata", 0755)
        os.WriteFile(goldenPath, []byte(got), 0644)
        return
    }

    expected, err := os.ReadFile(goldenPath)
    if err != nil {
        t.Fatalf("failed to read golden file: %v", err)
    }

    if got != string(expected) {
        t.Errorf("output doesn't match golden file\ngot:\n%s\nwant:\n%s", got, expected)
    }
}
```

### Acceptance Criteria for v1.0.0

- [ ] All public API documented with godoc comments
- [ ] README.md contains quick start and examples
- [ ] Golden file test passes
- [ ] `go test ./...` passes on Go 1.21+
- [ ] No external dependencies in canvas/, draw/, text/ packages
- [ ] GitHub release created with v1.0.0 tag

---

## Verification Plan

### Per-Release Testing

Each minor version should pass:

```bash
# Build check
go build ./...

# Test suite
go test ./...

# Lint (optional but recommended)
go vet ./...

# Run examples (manual verification)
go run ./examples/basic
go run ./examples/maze
go run ./examples/hud
```

### Visual Verification

For each release, manually verify:

1. **v0.1.0**: Create a canvas, set pixels, verify Frame() output in terminal shows braille characters
2. **v0.2.0**: Draw lines and rectangles, verify they appear correctly
3. **v0.3.0**: Draw with colors, verify ANSI codes render correctly in terminal
4. **v0.4.0**: Render text, verify characters are readable
5. **v0.5.0**: Run all examples, verify they produce expected visual output
6. **v1.0.0**: Full integration test with maze frame + HUD + color

### Integration Test

Create a small test program that exercises all v1.0 features:

```go
func TestFullIntegration(t *testing.T) {
    c := canvas.New(160, 80, canvas.WithColor(), canvas.WithInvertedY())

    // Draw maze frame
    draw.Rectangle(c, 10, 10, 140, 60)
    draw.Line(c, 10, 10, 80, 40)
    draw.Line(c, 150, 10, 80, 40)

    // Draw HUD
    draw.RectangleFilled(c, 0, 70, 160, 10)
    text.DrawWithColor(c, 5, 72, "SCORE: 0", nil, canvas.ColorYellow)

    frame := c.Frame()

    // Verify frame is non-empty and contains expected content
    if len(frame) < 100 {
        t.Error("frame too short")
    }
}
```

---

## Summary

| Version | Focus | Key Deliverables |
|---------|-------|------------------|
| v0.1.0 | Core Canvas | Braille encoding, Set/Unset/Toggle/Get/Clear/Frame |
| v0.2.0 | Drawing | Bresenham lines, Rectangle stroke/fill |
| v0.3.0 | Color | 8 ANSI colors, per-cell coloring |
| v0.4.0 | Text | Bitmap font, text rendering |
| v0.5.0 | Integration | Terminal size, examples |
| v1.0.0 | Release | Documentation, polish, golden tests |

Each version is independently testable and builds incrementally toward the full Maze Wars rendering capability.
