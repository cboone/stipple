package draw

import (
	"testing"

	"github.com/cboone/stipple/canvas"
)

func TestRectangle(t *testing.T) {
	c := canvas.New(20, 16)
	Rectangle(c, 2, 2, 10, 8)

	// Verify all edge pixels are set
	// Top edge: y=2, x from 2 to 11
	for x := 2; x <= 11; x++ {
		if !c.Get(float64(x), 2) {
			t.Errorf("top edge pixel (%d, 2) not set", x)
		}
	}
	// Bottom edge: y=9, x from 2 to 11
	for x := 2; x <= 11; x++ {
		if !c.Get(float64(x), 9) {
			t.Errorf("bottom edge pixel (%d, 9) not set", x)
		}
	}
	// Left edge: x=2, y from 2 to 9
	for y := 2; y <= 9; y++ {
		if !c.Get(2, float64(y)) {
			t.Errorf("left edge pixel (2, %d) not set", y)
		}
	}
	// Right edge: x=11, y from 2 to 9
	for y := 2; y <= 9; y++ {
		if !c.Get(11, float64(y)) {
			t.Errorf("right edge pixel (11, %d) not set", y)
		}
	}

	// Verify interior pixels are NOT set
	for y := 3; y <= 8; y++ {
		for x := 3; x <= 10; x++ {
			if c.Get(float64(x), float64(y)) {
				t.Errorf("interior pixel (%d, %d) should not be set", x, y)
			}
		}
	}

	printVisual(t, "TestRectangle", c)
}

func TestRectangleFilled(t *testing.T) {
	c := canvas.New(20, 16)
	RectangleFilled(c, 2, 2, 10, 8)

	// Verify ALL pixels in range (2-11, 2-9) are set
	for y := 2; y <= 9; y++ {
		for x := 2; x <= 11; x++ {
			if !c.Get(float64(x), float64(y)) {
				t.Errorf("pixel (%d, %d) should be set", x, y)
			}
		}
	}

	// Verify pixels outside rectangle are NOT set
	// Check row above
	for x := 0; x < 20; x++ {
		if c.Get(float64(x), 1) {
			t.Errorf("pixel (%d, 1) should not be set", x)
		}
	}
	// Check row below
	for x := 0; x < 20; x++ {
		if c.Get(float64(x), 10) {
			t.Errorf("pixel (%d, 10) should not be set", x)
		}
	}
	// Check column left
	for y := 0; y < 16; y++ {
		if c.Get(1, float64(y)) {
			t.Errorf("pixel (1, %d) should not be set", y)
		}
	}
	// Check column right
	for y := 0; y < 16; y++ {
		if c.Get(12, float64(y)) {
			t.Errorf("pixel (12, %d) should not be set", y)
		}
	}

	printVisual(t, "TestRectangleFilled", c)
}

func TestRectangleZeroSize(t *testing.T) {
	tests := []struct {
		name          string
		width, height float64
	}{
		{"width zero", 0, 5},
		{"height zero", 5, 0},
		{"both zero", 0, 0},
		{"width negative", -5, 5},
		{"height negative", 5, -5},
		{"both negative", -5, -5},
	}

	for _, testCase := range tests {
		c := canvas.New(20, 16)
		Rectangle(c, 5, 5, testCase.width, testCase.height)

		// Count set pixels
		count := 0
		for y := 0; y < 16; y++ {
			for x := 0; x < 20; x++ {
				if c.Get(float64(x), float64(y)) {
					count++
				}
			}
		}
		if count != 0 {
			t.Errorf("%s: expected 0 pixels set, got %d", testCase.name, count)
		}
	}
}

func TestRectangleFilledZeroSize(t *testing.T) {
	tests := []struct {
		name          string
		width, height float64
	}{
		{"width zero", 0, 5},
		{"height zero", 5, 0},
		{"both zero", 0, 0},
		{"width negative", -5, 5},
		{"height negative", 5, -5},
		{"both negative", -5, -5},
	}

	for _, testCase := range tests {
		c := canvas.New(20, 16)
		RectangleFilled(c, 5, 5, testCase.width, testCase.height)

		// Count set pixels
		count := 0
		for y := 0; y < 16; y++ {
			for x := 0; x < 20; x++ {
				if c.Get(float64(x), float64(y)) {
					count++
				}
			}
		}
		if count != 0 {
			t.Errorf("%s: expected 0 pixels set, got %d", testCase.name, count)
		}
	}
}

func TestRectanglePartiallyOffCanvas(t *testing.T) {
	c := canvas.New(20, 16)
	Rectangle(c, -5, -5, 15, 15)

	// Only the on-canvas portion should be drawn
	// Rectangle from (-5,-5) to (9,9), but canvas is (0,0) to (19,15)
	// Visible edges: right edge at x=9 (y from 0 to 9), bottom edge at y=9 (x from 0 to 9)

	// Verify right edge (x=9, y=0 to 9)
	for y := 0; y <= 9; y++ {
		if !c.Get(9, float64(y)) {
			t.Errorf("right edge pixel (9, %d) not set", y)
		}
	}

	// Verify bottom edge (y=9, x=0 to 9)
	for x := 0; x <= 9; x++ {
		if !c.Get(float64(x), 9) {
			t.Errorf("bottom edge pixel (%d, 9) not set", x)
		}
	}

	printVisual(t, "TestRectanglePartiallyOffCanvas", c)
}

func TestRectangleFilledPartiallyOffCanvas(t *testing.T) {
	c := canvas.New(20, 16)
	RectangleFilled(c, -5, -5, 15, 15)

	// Only the on-canvas portion should be filled
	// Rectangle from (-5,-5) to (9,9), visible portion is (0,0) to (9,9)

	// Verify all visible pixels are set
	for y := 0; y <= 9; y++ {
		for x := 0; x <= 9; x++ {
			if !c.Get(float64(x), float64(y)) {
				t.Errorf("pixel (%d, %d) should be set", x, y)
			}
		}
	}

	// Verify pixels outside are not set
	for y := 0; y < 16; y++ {
		if c.Get(10, float64(y)) {
			t.Errorf("pixel (10, %d) should not be set", y)
		}
	}
	for x := 0; x < 20; x++ {
		if c.Get(float64(x), 10) {
			t.Errorf("pixel (%d, 10) should not be set", x)
		}
	}

	printVisual(t, "TestRectangleFilledPartiallyOffCanvas", c)
}

func TestRectangleSmall(t *testing.T) {
	// 1x1 rectangle: single pixel
	c1 := canvas.New(10, 8)
	Rectangle(c1, 3, 3, 1, 1)
	if !c1.Get(3, 3) {
		t.Error("1x1 rectangle: pixel (3, 3) not set")
	}
	count := 0
	for y := 0; y < 8; y++ {
		for x := 0; x < 10; x++ {
			if c1.Get(float64(x), float64(y)) {
				count++
			}
		}
	}
	if count != 1 {
		t.Errorf("1x1 rectangle: expected 1 pixel, got %d", count)
	}
	printVisual(t, "TestRectangleSmall 1x1", c1)

	// 2x2 outline: 4 corner pixels
	c2 := canvas.New(10, 8)
	Rectangle(c2, 3, 3, 2, 2)
	corners := []struct{ x, y float64 }{{3, 3}, {4, 3}, {3, 4}, {4, 4}}
	for _, corner := range corners {
		if !c2.Get(corner.x, corner.y) {
			t.Errorf("2x2 rectangle: corner pixel (%.0f, %.0f) not set", corner.x, corner.y)
		}
	}
	count = 0
	for y := 0; y < 8; y++ {
		for x := 0; x < 10; x++ {
			if c2.Get(float64(x), float64(y)) {
				count++
			}
		}
	}
	if count != 4 {
		t.Errorf("2x2 rectangle: expected 4 pixels, got %d", count)
	}
	printVisual(t, "TestRectangleSmall 2x2", c2)
}

func TestRectangleFilledSmall(t *testing.T) {
	// 1x1 filled: single pixel
	c1 := canvas.New(10, 8)
	RectangleFilled(c1, 3, 3, 1, 1)
	if !c1.Get(3, 3) {
		t.Error("1x1 filled: pixel (3, 3) not set")
	}
	count := 0
	for y := 0; y < 8; y++ {
		for x := 0; x < 10; x++ {
			if c1.Get(float64(x), float64(y)) {
				count++
			}
		}
	}
	if count != 1 {
		t.Errorf("1x1 filled: expected 1 pixel, got %d", count)
	}
	printVisual(t, "TestRectangleFilledSmall 1x1", c1)

	// 2x2 filled: all 4 pixels
	c2 := canvas.New(10, 8)
	RectangleFilled(c2, 3, 3, 2, 2)
	pixels := []struct{ x, y float64 }{{3, 3}, {4, 3}, {3, 4}, {4, 4}}
	for _, pixel := range pixels {
		if !c2.Get(pixel.x, pixel.y) {
			t.Errorf("2x2 filled: pixel (%.0f, %.0f) not set", pixel.x, pixel.y)
		}
	}
	count = 0
	for y := 0; y < 8; y++ {
		for x := 0; x < 10; x++ {
			if c2.Get(float64(x), float64(y)) {
				count++
			}
		}
	}
	if count != 4 {
		t.Errorf("2x2 filled: expected 4 pixels, got %d", count)
	}
	printVisual(t, "TestRectangleFilledSmall 2x2", c2)
}

func TestRectangleFilledFloatCoordinates(t *testing.T) {
	c := canvas.New(20, 16)
	// Rectangle at (2.7, 3.9) should start at (2, 3)
	RectangleFilled(c, 2.7, 3.9, 5, 4)

	// Verify floored start position
	if !c.Get(2, 3) {
		t.Error("floored start pixel (2, 3) not set")
	}

	// Verify end position: floor(2.7 + 5 - 1) = floor(6.7) = 6
	// floor(3.9 + 4 - 1) = floor(6.9) = 6
	if !c.Get(6, 6) {
		t.Error("floored end pixel (6, 6) not set")
	}

	// Count total pixels: should be 5 * 4 = 20
	count := 0
	for y := 0; y < 16; y++ {
		for x := 0; x < 20; x++ {
			if c.Get(float64(x), float64(y)) {
				count++
			}
		}
	}
	if count != 20 {
		t.Errorf("expected 20 pixels, got %d", count)
	}

	printVisual(t, "TestRectangleFilledFloatCoordinates", c)
}
