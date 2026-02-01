package draw

import (
	"flag"
	"os"
	"testing"

	"github.com/cboone/brodot/canvas"
)

var visualFlag = flag.Bool("visual", false, "print visual output")

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

func printVisual(t *testing.T, name string, c *canvas.Canvas) {
	if *visualFlag {
		t.Logf("\n=== %s ===\n%s", name, c.Frame())
	}
}

func TestLineHorizontal(t *testing.T) {
	c := canvas.New(10, 4)
	Line(c, 1, 1, 8, 1)

	// Verify all pixels along the horizontal line are set
	for x := 1; x <= 8; x++ {
		if !c.Get(float64(x), 1) {
			t.Errorf("pixel (%d, 1) not set for horizontal line", x)
		}
	}

	// Verify pixels outside the line are not set
	if c.Get(0, 1) {
		t.Error("pixel (0, 1) should not be set")
	}
	if c.Get(9, 1) {
		t.Error("pixel (9, 1) should not be set")
	}

	printVisual(t, "TestLineHorizontal", c)
}

func TestLineVertical(t *testing.T) {
	c := canvas.New(4, 12)
	Line(c, 1, 1, 1, 10)

	// Verify all pixels along the vertical line are set
	for y := 1; y <= 10; y++ {
		if !c.Get(1, float64(y)) {
			t.Errorf("pixel (1, %d) not set for vertical line", y)
		}
	}

	// Verify pixels outside the line are not set
	if c.Get(1, 0) {
		t.Error("pixel (1, 0) should not be set")
	}
	if c.Get(1, 11) {
		t.Error("pixel (1, 11) should not be set")
	}

	printVisual(t, "TestLineVertical", c)
}

func TestLineDiagonalPositive(t *testing.T) {
	c := canvas.New(10, 12)
	Line(c, 0, 0, 7, 7)

	// Verify pixels along the 45-degree diagonal
	for index := 0; index <= 7; index++ {
		if !c.Get(float64(index), float64(index)) {
			t.Errorf("pixel (%d, %d) not set for diagonal line", index, index)
		}
	}

	printVisual(t, "TestLineDiagonalPositive", c)
}

func TestLineDiagonalNegative(t *testing.T) {
	c := canvas.New(10, 12)
	Line(c, 7, 0, 0, 7)

	// Verify pixels along the 45-degree diagonal with negative slope
	for index := 0; index <= 7; index++ {
		if !c.Get(float64(7-index), float64(index)) {
			t.Errorf("pixel (%d, %d) not set for negative diagonal line", 7-index, index)
		}
	}

	printVisual(t, "TestLineDiagonalNegative", c)
}

func TestLineSymmetry(t *testing.T) {
	// Line(a, b, c, d) should produce the same pixels as Line(c, d, a, b)
	c1 := canvas.New(20, 16)
	c2 := canvas.New(20, 16)

	Line(c1, 2, 3, 15, 11)
	Line(c2, 15, 11, 2, 3)

	// Compare all pixels
	for x := 0; x < 20; x++ {
		for y := 0; y < 16; y++ {
			pixel1 := c1.Get(float64(x), float64(y))
			pixel2 := c2.Get(float64(x), float64(y))
			if pixel1 != pixel2 {
				t.Errorf("asymmetric pixel at (%d, %d): forward=%v, reverse=%v", x, y, pixel1, pixel2)
			}
		}
	}

	printVisual(t, "TestLineSymmetry (forward)", c1)
	printVisual(t, "TestLineSymmetry (reverse)", c2)
}

func TestLineSinglePoint(t *testing.T) {
	c := canvas.New(4, 4)
	Line(c, 1, 1, 1, 1)

	// Verify the single point is set
	if !c.Get(1, 1) {
		t.Error("single point (1, 1) not set")
	}

	// Count total set pixels
	count := 0
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			if c.Get(float64(x), float64(y)) {
				count++
			}
		}
	}
	if count != 1 {
		t.Errorf("expected 1 pixel set, got %d", count)
	}

	printVisual(t, "TestLineSinglePoint", c)
}

func TestLineShallowSlope(t *testing.T) {
	// Slope < 1 (more horizontal than vertical)
	c := canvas.New(20, 8)
	Line(c, 0, 0, 15, 3)

	// Verify start and end points
	if !c.Get(0, 0) {
		t.Error("start point (0, 0) not set")
	}
	if !c.Get(15, 3) {
		t.Error("end point (15, 3) not set")
	}

	// Count total pixels (should be 16 for a line of length 16 in x)
	count := 0
	for x := 0; x < 20; x++ {
		for y := 0; y < 8; y++ {
			if c.Get(float64(x), float64(y)) {
				count++
			}
		}
	}
	// For a shallow line, the number of pixels should equal dx + 1
	expectedCount := 16
	if count != expectedCount {
		t.Errorf("expected %d pixels, got %d", expectedCount, count)
	}

	printVisual(t, "TestLineShallowSlope", c)
}

func TestLineSteepSlope(t *testing.T) {
	// Slope > 1 (more vertical than horizontal)
	c := canvas.New(8, 20)
	Line(c, 0, 0, 3, 15)

	// Verify start and end points
	if !c.Get(0, 0) {
		t.Error("start point (0, 0) not set")
	}
	if !c.Get(3, 15) {
		t.Error("end point (3, 15) not set")
	}

	// Count total pixels (should be 16 for a line of length 16 in y)
	count := 0
	for x := 0; x < 8; x++ {
		for y := 0; y < 20; y++ {
			if c.Get(float64(x), float64(y)) {
				count++
			}
		}
	}
	// For a steep line, the number of pixels should equal dy + 1
	expectedCount := 16
	if count != expectedCount {
		t.Errorf("expected %d pixels, got %d", expectedCount, count)
	}

	printVisual(t, "TestLineSteepSlope", c)
}

func TestLineReversedCoordinates(t *testing.T) {
	// Test drawing lines with reversed start/end coordinates
	tests := []struct {
		name                           string
		startX, startY, endX, endY     float64
		checkX, checkY                 int
	}{
		{"horizontal reversed", 8, 1, 1, 1, 4, 1},
		{"vertical reversed", 1, 10, 1, 1, 1, 5},
		{"diagonal reversed", 7, 7, 0, 0, 3, 3},
	}

	for _, testCase := range tests {
		c := canvas.New(12, 12)
		Line(c, testCase.startX, testCase.startY, testCase.endX, testCase.endY)

		// Check that a midpoint pixel is set
		if !c.Get(float64(testCase.checkX), float64(testCase.checkY)) {
			t.Errorf("%s: midpoint pixel (%d, %d) not set", testCase.name, testCase.checkX, testCase.checkY)
		}

		printVisual(t, "TestLineReversedCoordinates_"+testCase.name, c)
	}
}

func TestLineFloatCoordinates(t *testing.T) {
	// Test that float coordinates are properly floored
	c := canvas.New(10, 8)
	Line(c, 0.9, 0.9, 5.9, 3.9)

	// Should be equivalent to Line(0, 0, 5, 3)
	if !c.Get(0, 0) {
		t.Error("floored start point (0, 0) not set")
	}
	if !c.Get(5, 3) {
		t.Error("floored end point (5, 3) not set")
	}

	printVisual(t, "TestLineFloatCoordinates", c)
}
