# go-drawille Implementation Plan

A comprehensive Go port of drawille that incorporates the best features from 20+ implementations across different languages.

## Executive Summary

This plan synthesizes learnings from the original Python drawille and its ports to Java, Node.js, Go, Ruby, Julia, Rust, C, C++, Perl, and others. The goal is to create a feature-rich, idiomatic Go library that combines:

- **Core simplicity** from the original Python version
- **Drawing algorithms** from Kerrigan29a/drawille-go (Bresenham, Bezier curves)
- **HTML5 Canvas-like API** from node-drawille-canvas
- **Animation framework** from rsille (Rust)
- **Advanced features** from Term-Graille (Perl): thick lines, text overlay, scrolling
- **Plotting capabilities** from TextPlots.jl (Julia)
- **Idiomatic Go design** with proper interfaces and options patterns

---

## 1. Architecture Overview

### 1.1 Package Structure

```
go-drawille/
├── canvas/           # Core canvas implementation
│   ├── canvas.go     # Main Canvas struct and methods
│   ├── braille.go    # Braille character encoding/decoding
│   ├── color.go      # Terminal color support
│   └── options.go    # Functional options pattern
├── draw/             # Drawing primitives
│   ├── line.go       # Line algorithms (Bresenham)
│   ├── circle.go     # Circle/ellipse (Bresenham)
│   ├── bezier.go     # Quadratic and cubic Bezier curves
│   ├── polygon.go    # Polygon fill/stroke
│   ├── thick.go      # Murphy's thick line algorithm
│   └── text.go       # Text rendering with bitmap fonts
├── turtle/           # Turtle graphics
│   ├── turtle.go     # Logo-style turtle
│   └── commands.go   # Extended turtle commands
├── animation/        # Animation framework
│   ├── animator.go   # Animation loop and timing
│   ├── drawable.go   # Drawable interface
│   └── easing.go     # Easing functions
├── plot/             # Plotting capabilities
│   ├── plot.go       # High-level plotting API
│   ├── axis.go       # Axis rendering and labels
│   ├── series.go     # Data series handling
│   └── functions.go  # Mathematical function plotting
├── transform/        # Coordinate transformations
│   ├── matrix.go     # 2D transformation matrices
│   └── 3d.go         # 3D projection support
├── io/               # Import/Export
│   ├── image.go      # Image conversion (with build tag)
│   └── export.go     # Export to string/file
├── examples/         # Example programs
└── internal/         # Internal utilities
    └── term/         # Terminal detection and sizing
```

### 1.2 Core Design Principles

1. **Zero dependencies for core** - Image processing optional via build tags
2. **Functional options pattern** - For flexible, backwards-compatible configuration
3. **Interface-driven design** - `Drawable`, `Transformable` interfaces
4. **Coordinate system flexibility** - Standard (Y-up) and inverted (Y-down) modes
5. **Thread-safe canvas** - Safe for concurrent pixel operations

---

## 2. Core Components

### 2.1 Braille Encoding (from all implementations)

Unicode Braille characters (U+2800 to U+28FF) encode an 8-dot pattern in a 2×4 grid:

```
Dot positions:     Bit values:
  0  3               0x01  0x08
  1  4               0x02  0x10
  2  5               0x04  0x20
  6  7               0x40  0x80
```

Each terminal character cell represents 2×4 pixels, giving 8× resolution improvement.

```go
// Braille offset for Unicode braille block
const BrailleOffset = 0x2800

// Pixel to braille bit mapping
var pixelMap = [4][2]rune{
    {0x01, 0x08},
    {0x02, 0x10},
    {0x04, 0x20},
    {0x40, 0x80},
}

func (c *Canvas) Set(x, y float64) {
    px, py := int(x), int(y)
    cellX, cellY := px/2, py/4
    dotX, dotY := px%2, py%4
    c.cells[cellY][cellX] |= pixelMap[dotY][dotX]
}
```

### 2.2 Canvas API

Combining the best from Python, Node.js, and Rust implementations:

```go
type Canvas struct {
    width, height int           // Pixel dimensions
    cells         [][]rune      // Braille character grid
    colors        [][]Color     // Per-cell color (optional)
    invertY       bool          // Inverted Y-axis mode
    mu            sync.RWMutex  // Thread safety
}

// Core methods (from original drawille)
func New(width, height int, opts ...Option) *Canvas
func (c *Canvas) Set(x, y float64)
func (c *Canvas) Unset(x, y float64)
func (c *Canvas) Toggle(x, y float64)
func (c *Canvas) Get(x, y float64) bool
func (c *Canvas) Clear()
func (c *Canvas) Frame() string

// Extended methods (from various ports)
func (c *Canvas) Width() int
func (c *Canvas) Height() int
func (c *Canvas) Rows() int      // Height in terminal rows
func (c *Canvas) Cols() int      // Width in terminal columns
func (c *Canvas) SetColor(x, y float64, color Color)
func (c *Canvas) Fill(x, y float64)  // Flood fill
```

### 2.3 Functional Options (idiomatic Go)

```go
type Option func(*Canvas)

func WithInvertedY() Option {
    return func(c *Canvas) { c.invertY = true }
}

func WithColor() Option {
    return func(c *Canvas) { c.colors = make([][]Color, c.Rows()) }
}

func WithTerminalSize() Option {
    return func(c *Canvas) {
        cols, rows := term.Size()
        c.width, c.height = cols*2, rows*4
    }
}
```

---

## 3. Drawing Primitives

### 3.1 Line Drawing (from Kerrigan29a/drawille-go)

Bresenham's line algorithm for efficient integer-based line drawing:

```go
func (c *Canvas) Line(x1, y1, x2, y2 float64)
func (c *Canvas) LineThick(x1, y1, x2, y2, thickness float64)  // Murphy's algorithm
```

### 3.2 Circles and Ellipses (from Kerrigan29a/drawille-go)

```go
func (c *Canvas) Circle(cx, cy, r float64)
func (c *Canvas) CircleFilled(cx, cy, r float64)
func (c *Canvas) Ellipse(cx, cy, rx, ry float64)
func (c *Canvas) Arc(cx, cy, r, startAngle, endAngle float64)
```

### 3.3 Bezier Curves (from Kerrigan29a/drawille-go)

```go
func (c *Canvas) QuadraticBezier(x0, y0, x1, y1, x2, y2 float64)
func (c *Canvas) CubicBezier(x0, y0, x1, y1, x2, y2, x3, y3 float64)
```

### 3.4 Polygons (enhanced from multiple implementations)

```go
func (c *Canvas) Polygon(points []Point)
func (c *Canvas) PolygonFilled(points []Point)
func (c *Canvas) Rectangle(x, y, w, h float64)
func (c *Canvas) RectangleFilled(x, y, w, h float64)
```

### 3.5 Text Rendering (from Term-Graille)

Convert bitmap fonts to braille-compatible 4×2 character cells:

```go
func (c *Canvas) Text(x, y float64, text string, opts ...TextOption)

type TextOption func(*textConfig)
func WithFont(font *BitmapFont) TextOption
func WithScale(scale int) TextOption
```

---

## 4. Turtle Graphics

### 4.1 Core Turtle (from original Python)

```go
type Turtle struct {
    canvas  *Canvas
    x, y    float64
    heading float64  // Degrees, 0 = right
    penDown bool
}

func NewTurtle(canvas *Canvas) *Turtle
func (t *Turtle) Forward(distance float64)
func (t *Turtle) Backward(distance float64)
func (t *Turtle) Right(angle float64)
func (t *Turtle) Left(angle float64)
func (t *Turtle) PenUp()
func (t *Turtle) PenDown()
func (t *Turtle) Goto(x, y float64)
func (t *Turtle) Position() (float64, float64)
func (t *Turtle) Heading() float64
func (t *Turtle) SetHeading(angle float64)
```

### 4.2 Extended Turtle Commands (from various ports)

```go
func (t *Turtle) Circle(radius float64)
func (t *Turtle) Arc(radius, angle float64)
func (t *Turtle) Dot(size float64)
func (t *Turtle) Home()  // Return to center, heading 0
func (t *Turtle) SetColor(color Color)
```

---

## 5. HTML5 Canvas-like API (from node-drawille-canvas)

### 5.1 2D Context

Provides familiar API for those coming from web development:

```go
type Context2D struct {
    canvas     *Canvas
    path       []pathOp
    transform  Matrix3x3
    stack      []state
    fillColor  Color
    strokeColor Color
}

func (c *Canvas) GetContext2D() *Context2D

// State management
func (ctx *Context2D) Save()
func (ctx *Context2D) Restore()

// Transformations
func (ctx *Context2D) Translate(x, y float64)
func (ctx *Context2D) Rotate(angle float64)
func (ctx *Context2D) Scale(x, y float64)
func (ctx *Context2D) SetTransform(a, b, c, d, e, f float64)

// Path operations
func (ctx *Context2D) BeginPath()
func (ctx *Context2D) ClosePath()
func (ctx *Context2D) MoveTo(x, y float64)
func (ctx *Context2D) LineTo(x, y float64)
func (ctx *Context2D) Arc(x, y, r, startAngle, endAngle float64, counterclockwise bool)
func (ctx *Context2D) QuadraticCurveTo(cpx, cpy, x, y float64)
func (ctx *Context2D) BezierCurveTo(cp1x, cp1y, cp2x, cp2y, x, y float64)
func (ctx *Context2D) Rect(x, y, w, h float64)

// Drawing
func (ctx *Context2D) Fill()
func (ctx *Context2D) Stroke()
func (ctx *Context2D) FillRect(x, y, w, h float64)
func (ctx *Context2D) StrokeRect(x, y, w, h float64)
func (ctx *Context2D) ClearRect(x, y, w, h float64)
```

---

## 6. Animation Framework (from rsille)

### 6.1 Animation Loop

```go
type Animator struct {
    canvas    *Canvas
    objects   []Drawable
    fps       int
    running   bool
}

type Drawable interface {
    Draw(c *Canvas)
    Update(dt float64) bool  // Returns false to remove
}

func NewAnimator(canvas *Canvas, fps int) *Animator
func (a *Animator) Add(d Drawable)
func (a *Animator) AddWithPosition(d Drawable, x, y float64)
func (a *Animator) Run(ctx context.Context)
func (a *Animator) Stop()
```

### 6.2 Easing Functions

```go
type EaseFunc func(t float64) float64

var (
    EaseLinear     EaseFunc
    EaseInQuad     EaseFunc
    EaseOutQuad    EaseFunc
    EaseInOutQuad  EaseFunc
    EaseInCubic    EaseFunc
    EaseOutCubic   EaseFunc
    EaseInOutCubic EaseFunc
    // ... more easing functions
)
```

---

## 7. Plotting Capabilities (from TextPlots.jl)

### 7.1 High-Level Plotting API

```go
type Plot struct {
    canvas     *Canvas
    title      string
    xLabel     string
    yLabel     string
    showBorder bool
    series     []Series
}

func NewPlot(opts ...PlotOption) *Plot

// Plot functions
func (p *Plot) Function(f func(float64) float64, xMin, xMax float64)
func (p *Plot) Functions(fs []func(float64) float64, xMin, xMax float64)

// Plot data
func (p *Plot) Scatter(x, y []float64)
func (p *Plot) Line(x, y []float64)
func (p *Plot) Bar(values []float64)  // Uses inverted Y mode

// Render
func (p *Plot) String() string
func (p *Plot) Render() string
```

### 7.2 Plot Options

```go
type PlotOption func(*Plot)

func WithTitle(title string) PlotOption
func WithXLabel(label string) PlotOption
func WithYLabel(label string) PlotOption
func WithBorder(show bool) PlotOption
func WithSize(cols, rows int) PlotOption
func WithXRange(min, max float64) PlotOption
func WithYRange(min, max float64) PlotOption
```

---

## 8. 3D Support (from examples in multiple implementations)

### 8.1 3D Primitives

```go
type Point3D struct {
    X, Y, Z float64
}

type Object3D interface {
    Vertices() []Point3D
    Edges() [][2]int  // Pairs of vertex indices
    Transform(m Matrix4x4) Object3D
}

// Built-in shapes
func NewCube(size float64) Object3D
func NewPyramid(base, height float64) Object3D
func NewSphere(radius float64, segments int) Object3D
```

### 8.2 3D Projection

```go
type Camera struct {
    Position Point3D
    Target   Point3D
    Up       Point3D
    FOV      float64
}

func (cam *Camera) Project(p Point3D, screenWidth, screenHeight int) (x, y float64)

func (c *Canvas) DrawObject3D(obj Object3D, cam *Camera)
```

---

## 9. Color Support

### 9.1 Terminal Colors (from Rust drawille)

```go
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

// ANSI escape code generation
func (clr Color) ANSI() string
func (clr Color) ANSIBright() string
```

### 9.2 RGB/256 Color Support (extended)

```go
type Color256 uint8
type ColorRGB struct{ R, G, B uint8 }

func (c Color256) ANSI() string
func (c ColorRGB) ANSI() string
```

---

## 10. Image Support (optional, with build tags)

### 10.1 Image Conversion (from rsille and jpverkamp)

The image conversion approach originates from JP Verkamp's Racket implementation:
- Convert each 2×4 pixel block to grayscale
- Compare against a threshold (default ~0.75) to determine dot state
- Out-of-bounds pixels default to 0 (off) for graceful edge handling

```go
//go:build with_image

import "image"

func FromImage(img image.Image, opts ...ImageOption) *Canvas
func (c *Canvas) ToImage() image.Image

type ImageOption func(*imageConfig)
func WithDithering(d DitheringMethod) ImageOption
func WithThreshold(t float64) ImageOption  // 0.0-1.0, default 0.75
func WithGrayscale() ImageOption
func WithInvert() ImageOption
func WithAutoThreshold() ImageOption  // Otsu's method
```

---

## 11. Testing Strategy

### 11.1 Unit Tests

**Braille Encoding Tests:**
```go
func TestPixelToBraille(t *testing.T) {
    tests := []struct{
        pixels []Point
        want   rune
    }{
        {[]Point{{0,0}}, '\u2801'},           // Single dot top-left
        {[]Point{{1,0}}, '\u2808'},           // Single dot top-right
        {[]Point{{0,0},{1,0},{0,1},{1,1},{0,2},{1,2},{0,3},{1,3}}, '\u28FF'}, // Full block
    }
    // ...
}
```

**Canvas Operations:**
```go
func TestCanvasSetGet(t *testing.T)
func TestCanvasToggle(t *testing.T)
func TestCanvasClear(t *testing.T)
func TestCanvasFrame(t *testing.T)
func TestCanvasInvertedY(t *testing.T)
func TestCanvasConcurrentAccess(t *testing.T)
```

### 11.2 Drawing Algorithm Tests

**Line Drawing (Bresenham):**
```go
func TestLineHorizontal(t *testing.T)
func TestLineVertical(t *testing.T)
func TestLineDiagonal(t *testing.T)
func TestLineSymmetry(t *testing.T)  // Line(a,b) == Line(b,a)
```

**Circle Drawing:**
```go
func TestCircleSymmetry(t *testing.T)  // 8-way symmetry
func TestCircleCompleteness(t *testing.T)  // No gaps
```

**Bezier Curves:**
```go
func TestQuadraticBezierEndpoints(t *testing.T)
func TestCubicBezierEndpoints(t *testing.T)
func TestBezierSmoothness(t *testing.T)  // No gaps between segments
```

### 11.3 Golden File Tests

Compare rendered output against known-good outputs:

```go
func TestGoldenSineWave(t *testing.T) {
    c := canvas.New(80, 40)
    for x := 0.0; x < 80; x++ {
        y := 20 + 15*math.Sin(x*0.1)
        c.Set(x, y)
    }
    golden := readGoldenFile("testdata/sine_wave.golden")
    if c.Frame() != golden {
        t.Errorf("sine wave output doesn't match golden file")
    }
}
```

**Golden files to create:**
- `sine_wave.golden` - Basic sine wave
- `circle.golden` - Circle at various sizes
- `turtle_spiral.golden` - Turtle-drawn spiral
- `text_hello.golden` - Rendered text
- `cube_wireframe.golden` - 3D cube projection

### 11.4 Property-Based Tests

Using a library like `gopter`:

```go
func TestCanvasPropertySetGet(t *testing.T) {
    properties := gopter.NewProperties(nil)

    properties.Property("Set then Get returns true", prop.ForAll(
        func(x, y int) bool {
            c := canvas.New(100, 100)
            c.Set(float64(x%100), float64(y%100))
            return c.Get(float64(x%100), float64(y%100))
        },
        gen.Int(), gen.Int(),
    ))

    properties.TestingRun(t)
}
```

### 11.5 Benchmark Tests

```go
func BenchmarkCanvasSet(b *testing.B)
func BenchmarkCanvasFrame(b *testing.B)
func BenchmarkLineBresenham(b *testing.B)
func BenchmarkCircle(b *testing.B)
func BenchmarkCubicBezier(b *testing.B)
func BenchmarkAnimation60FPS(b *testing.B)
```

### 11.6 Visual Regression Tests

Generate PNG from Canvas output and compare:

```go
func TestVisualRegression(t *testing.T) {
    if os.Getenv("VISUAL_TESTS") == "" {
        t.Skip("Set VISUAL_TESTS=1 to run visual regression tests")
    }

    c := canvas.New(160, 80)
    // Draw complex scene

    actual := renderToPNG(c)
    expected := loadPNG("testdata/complex_scene.png")

    diff := imageDiff(actual, expected)
    if diff > 0.01 {  // 1% tolerance
        t.Errorf("Visual regression: %f%% difference", diff*100)
    }
}
```

### 11.7 Example-Based Tests

Every public function should have an example that also serves as a test:

```go
func ExampleCanvas_Set() {
    c := canvas.New(10, 10)
    c.Set(0, 0)
    c.Set(1, 0)
    fmt.Println(c.Frame())
    // Output: ⠉
}

func ExampleTurtle_spiral() {
    c := canvas.New(80, 40)
    t := turtle.New(c)
    t.PenDown()
    for i := 0; i < 100; i++ {
        t.Forward(float64(i) * 0.5)
        t.Right(91)
    }
    // Output is a spiral shape
}
```

### 11.8 Fuzz Testing (Go 1.18+)

```go
func FuzzCanvasSet(f *testing.F) {
    f.Add(0.0, 0.0)
    f.Add(100.0, 100.0)
    f.Add(-1.0, -1.0)
    f.Add(math.MaxFloat64, math.MaxFloat64)

    f.Fuzz(func(t *testing.T, x, y float64) {
        c := canvas.New(100, 100)
        // Should not panic
        c.Set(x, y)
        c.Frame()
    })
}
```

### 11.9 Integration Tests

**End-to-end animation test:**
```go
func TestAnimationIntegration(t *testing.T) {
    c := canvas.New(80, 40)
    a := animation.NewAnimator(c, 30)

    ball := &BouncingBall{x: 40, y: 20, vx: 2, vy: 1}
    a.Add(ball)

    ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
    defer cancel()

    a.Run(ctx)
    // Verify ball moved
    if ball.x == 40 && ball.y == 20 {
        t.Error("Ball should have moved")
    }
}
```

---

## 12. Implementation Phases

### Phase 1: Core Foundation
1. Braille encoding/decoding
2. Basic Canvas (Set, Unset, Toggle, Get, Clear, Frame)
3. Functional options
4. Unit tests for core

### Phase 2: Drawing Primitives
1. Bresenham line algorithm
2. Bresenham circle algorithm
3. Bezier curves (quadratic and cubic)
4. Rectangle and polygon
5. Golden file tests

### Phase 3: Turtle Graphics
1. Basic turtle (Forward, Backward, Left, Right)
2. Pen up/down
3. Extended commands
4. Example programs

### Phase 4: Canvas2D API
1. Path operations
2. Transformations (translate, rotate, scale)
3. State stack (save/restore)
4. Fill and stroke

### Phase 5: Animation Framework
1. Animation loop
2. Drawable interface
3. Easing functions
4. Timing control

### Phase 6: Plotting
1. Function plotting
2. Data series (scatter, line, bar)
3. Axes and labels
4. Multiple series support

### Phase 7: Advanced Features
1. 3D projection
2. Color support
3. Text rendering
4. Image conversion (optional)

### Phase 8: Polish
1. Documentation
2. Performance optimization
3. Additional examples
4. CI/CD setup

---

## 13. API Design Decisions

### 13.1 Coordinate Types

Use `float64` for all coordinates (like rsille) rather than `int`:
- Allows sub-pixel positioning
- Enables smooth animations
- Natural for mathematical functions
- Conversion to int happens internally

### 13.2 Error Handling

Most drawing operations silently ignore out-of-bounds coordinates (like most implementations). This simplifies code and follows the principle of graceful degradation.

### 13.3 Thread Safety

Canvas uses `sync.RWMutex`:
- Multiple goroutines can read simultaneously
- Write operations are serialized
- Animation framework handles locking

### 13.4 Zero Values

```go
var c canvas.Canvas  // Valid, but zero-sized
c := canvas.New(0, 0)  // Creates minimum 2x4 canvas (one cell)
```

---

## 14. Compatibility Notes

### 14.1 Terminal Requirements

- Must support Unicode (UTF-8)
- Braille characters: U+2800 to U+28FF
- Recommended fonts: DejaVu Sans Mono, Hack, JetBrains Mono

### 14.2 Go Version

- Minimum: Go 1.18 (for generics if used, fuzz testing)
- Recommended: Go 1.21+

### 14.3 Build Tags

```go
//go:build with_image
// Enables image import/export functionality
// Adds dependency on image packages
```

---

## 15. Historical Context & Attribution

### Terminal Graphics Evolution

Terminal graphics have a long history predating braille-based approaches:

- **Sixel Graphics (1980s)** - DEC introduced "six pixel" graphics for dot matrix printers and later VT240/VT340 terminals. Each sixel encodes a 6-pixel-high, 1-pixel-wide column. Still supported by modern terminals like xterm, mlterm, and foot.

- **Block Characters** - Unicode includes various block drawing characters (U+2580-U+259F) offering 2×2 quadrant resolution. The lower half block (U+2584) technique uses foreground/background colors for 2 vertical pixels per cell.

- **Teletext/Videotex (1970s-80s)** - European broadcast systems used 2×3 block graphics (64 patterns) per character cell.

### Braille Graphics Timeline

The technique of using Unicode Braille characters (U+2800-U+28FF) for terminal pixel graphics has earlier origins than commonly cited:

- **[aempirei/Chat-Art](https://github.com/aempirei/Chat-Art)** (C/C++) - Created **February 16, 2013**. Contains `pgmtobraille.c` for converting images to braille. This appears to be the earliest known implementation, predating drawille by over a year.

- **[asciimoo/drawille](https://github.com/asciimoo/drawille)** (Python) - Created **April 22, 2014**. Popularized the technique with a clean Canvas + Turtle API, spawning 20+ ports across languages.

- **[JP Verkamp's Racket implementation](https://blog.jverkamp.com/2014/05/30/braille-unicode-pixelation/)** ([source](https://github.com/jpverkamp/small-projects/blob/master/blog/braille-images.rkt)) - Published **May 30, 2014**. Provided clear explanation of the braille encoding and threshold-based image conversion.

All implementations recognized the key insight: the Braille Unicode block encodes all 256 combinations of an 8-dot pattern in a 2×4 grid, giving 8× pixel density improvement over standard characters.

### Modern Comprehensive Tools

- **[hpjansson/chafa](https://github.com/hpjansson/chafa)** (C) - Created 2018. Comprehensive terminal graphics supporting ASCII, braille, sixel, Kitty protocol, and iTerm2 protocol. The gold standard for image-to-terminal conversion.

### Historical Note on Braille Dot Ordering

As JP Verkamp noted, the dot numbering (1,4,2,5,3,6,7,8 rather than sequential) exists because the original 6-dot Braille system predates the 8-dot extension. Dots 7 and 8 were added later at the bottom:

```
Classic 6-dot:    Extended 8-dot:
  1  4              1  4
  2  5              2  5
  3  6              3  6
                    7  8
```

This historical quirk means the bit mapping isn't intuitive, but all implementations must follow it for correct Unicode encoding.

---

## 16. References

### Notable Implementations Reviewed:

| Implementation | Language | Year | Notable Features |
|----------------|----------|------|------------------|
| [aempirei/Chat-Art](https://github.com/aempirei/Chat-Art) | C/C++ | 2013 | **Earliest known**, pgmtobraille image converter |
| [asciimoo/drawille](https://github.com/asciimoo/drawille) | Python | 2014 | Popularized technique, Canvas + Turtle |
| [jpverkamp/braille-images.rkt](https://github.com/jpverkamp/small-projects/blob/master/blog/braille-images.rkt) | Racket | 2014 | Clear algorithm explanation, threshold-based |
| [exrook/drawille-go](https://github.com/exrook/drawille-go) | Go | 2014 | Basic Go port |
| [Kerrigan29a/drawille-go](https://github.com/Kerrigan29a/drawille-go) | Go | 2015 | Bezier, inverse Y, Bresenham algorithms |
| [madbence/node-drawille](https://github.com/madbence/node-drawille) | Node.js | 2014 | Simple API, powers vtop |
| [madbence/node-drawille-canvas](https://github.com/madbence/node-drawille-canvas) | Node.js | 2014 | HTML5 Canvas 2D API |
| [ftxqxd/drawille-rs](https://github.com/ftxqxd/drawille-rs) | Rust | 2015 | Turtle, 8 terminal colors |
| [nidhoggfgg/rsille](https://github.com/nidhoggfgg/rsille) | Rust | 2023 | Animation framework, 3D, image, Game of Life |
| [asciimoo/lua-drawille](https://github.com/asciimoo/lua-drawille) | Lua | 2014 | 3D support, L-systems, Game of Life, DDL |
| [saiftynet/Term-Graille](https://github.com/saiftynet/Term-Graille) | Perl | 2021 | Thick lines, menus, audio, sprites |
| [sunetos/TextPlots.jl](https://github.com/sunetos/TextPlots.jl) | Julia | 2014 | Elegant function/data plotting |
| [dheera/python-termgraphics](https://github.com/dheera/python-termgraphics) | Python | 2017 | ANSI colors, ASCII fallback, numpy |
| [null93/drawille](https://github.com/null93/drawille) | Java | 2017 | Clean OO design, Maven build |
| [Huulivoide/libdrawille](https://github.com/Huulivoide/libdrawille) | C | 2016 | Low-level, benchmarks, CMake |
| [Nirei/vrawille](https://github.com/Nirei/vrawille) | V | 2020 | V language port, stbi image support |
| [hpjansson/chafa](https://github.com/hpjansson/chafa) | C | 2018 | Comprehensive: braille, sixel, kitty, iTerm2 |

### Additional Implementations (not reviewed in detail)

| Implementation | Language | Notes |
|----------------|----------|-------|
| [hoelzro/term-drawille](https://github.com/hoelzro/term-drawille) | Perl 5 | CPAN: Term::Drawille |
| [yamadapc/haskell-drawille](https://github.com/yamadapc/haskell-drawille) | Haskell | Functional approach |
| [mkremins/drawille-clj](https://github.com/mkremins/drawille-clj) | Clojure | Lisp variant |
| [Goheeca/cl-drawille](https://github.com/Goheeca/cl-drawille) | Common Lisp | Full CL implementation |
| [PMunch/drawille-nim](https://github.com/PMunch/drawille-nim) | Nim | Systems language port |
| [mydzor/bash-drawille](https://github.com/mydzor/bash-drawille) | Bash | Pure shell implementation |
| [whatthejeff/php-drawille](https://github.com/whatthejeff/php-drawille) | PHP | Web server compatible |
| [liam-middlebrook/drawille-sharp](https://github.com/liam-middlebrook/drawille-sharp) | C# | .NET implementation |
| [l-a-i-n/drawille-plusplus](https://github.com/l-a-i-n/drawille-plusplus) | C++ | Object-oriented C++ |
| [massn/elixir-drawille](https://github.com/massn/elixir-drawille) | Elixir | BEAM/Erlang ecosystem |

### Related Tools & Prior Art

| Tool | Description |
|------|-------------|
| [Sixel](https://en.wikipedia.org/wiki/Sixel) | DEC's 1980s bitmap graphics for terminals (6 pixels high) |
| [libsixel](https://saitoha.github.io/libsixel/) | Modern sixel encoder/decoder library |
| [Kitty Graphics Protocol](https://sw.kovidgoyal.net/kitty/graphics-protocol/) | Modern terminal graphics, true color, animations |
| [timg](https://github.com/hzeller/timg) | Terminal image/video viewer using various methods |
