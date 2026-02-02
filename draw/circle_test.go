package draw

import (
	"testing"

	"github.com/cboone/brodot/canvas"
)

func countSetPixels(c *canvas.Canvas, width, height int) int {
	count := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if c.Get(float64(x), float64(y)) {
				count++
			}
		}
	}
	return count
}

func TestCircleSymmetry(t *testing.T) {
	c := canvas.New(30, 30)
	Circle(c, 14, 14, 10)

	// Verify 8-way symmetry by checking that symmetric points are all set
	// For a point (cx+dx, cy+dy), check all 8 symmetric positions
	symmetryPoints := []struct {
		deltaX, deltaY int
	}{
		{10, 0},  // East point
		{0, 10},  // South point
		{7, 7},   // Diagonal (approximate for radius 10)
	}

	for _, point := range symmetryPoints {
		// All 8 symmetric positions should have the same state
		positions := []struct{ x, y int }{
			{14 + point.deltaX, 14 + point.deltaY},
			{14 - point.deltaX, 14 + point.deltaY},
			{14 + point.deltaX, 14 - point.deltaY},
			{14 - point.deltaX, 14 - point.deltaY},
			{14 + point.deltaY, 14 + point.deltaX},
			{14 - point.deltaY, 14 + point.deltaX},
			{14 + point.deltaY, 14 - point.deltaX},
			{14 - point.deltaY, 14 - point.deltaX},
		}

		for _, position := range positions {
			if !c.Get(float64(position.x), float64(position.y)) {
				t.Errorf("symmetric point (%d, %d) for delta (%d, %d) not set",
					position.x, position.y, point.deltaX, point.deltaY)
			}
		}
	}

	printVisual(t, "TestCircleSymmetry", c)
}

func TestCircleRadius0(t *testing.T) {
	c := canvas.New(10, 10)
	Circle(c, 5, 5, 0)

	// Should draw exactly 1 pixel at center
	if !c.Get(5, 5) {
		t.Error("center pixel (5, 5) not set")
	}

	count := countSetPixels(c, 10, 10)
	if count != 1 {
		t.Errorf("expected 1 pixel for radius 0, got %d", count)
	}

	printVisual(t, "TestCircleRadius0", c)
}

func TestCircleRadius1(t *testing.T) {
	c := canvas.New(10, 10)
	Circle(c, 5, 5, 1)

	// Radius 1 in Bresenham produces exactly 4 cardinal points
	// Cardinal points: (5+1,5), (5-1,5), (5,5+1), (5,5-1)
	cardinalPoints := []struct{ x, y float64 }{
		{6, 5}, // East
		{4, 5}, // West
		{5, 6}, // South
		{5, 4}, // North
	}

	for _, point := range cardinalPoints {
		if !c.Get(point.x, point.y) {
			t.Errorf("cardinal point (%.0f, %.0f) not set", point.x, point.y)
		}
	}

	count := countSetPixels(c, 10, 10)
	if count != 4 {
		t.Errorf("radius 1 circle: expected 4 pixels, got %d", count)
	}

	printVisual(t, "TestCircleRadius1", c)
}

func TestCircleDiagonalHandling(t *testing.T) {
	// Radius 7 exercises the x == y case (x=5, y=5 iteration)
	c := canvas.New(20, 20)
	Circle(c, 9, 9, 7)

	// Cardinal points
	if !c.Get(16, 9) {
		t.Error("east point (16, 9) not set")
	}
	if !c.Get(2, 9) {
		t.Error("west point (2, 9) not set")
	}
	if !c.Get(9, 16) {
		t.Error("south point (9, 16) not set")
	}
	if !c.Get(9, 2) {
		t.Error("north point (9, 2) not set")
	}

	// Diagonal points at x==y==5 offset from center (9,9)
	// These are at approximately 45 degrees
	diagonalPoints := []struct{ x, y float64 }{
		{14, 14}, // Southeast (9+5, 9+5)
		{4, 14},  // Southwest (9-5, 9+5)
		{14, 4},  // Northeast (9+5, 9-5)
		{4, 4},   // Northwest (9-5, 9-5)
	}

	for _, point := range diagonalPoints {
		if !c.Get(point.x, point.y) {
			t.Errorf("diagonal point (%.0f, %.0f) not set", point.x, point.y)
		}
	}

	printVisual(t, "TestCircleDiagonalHandling", c)
}

func TestCircleNegativeRadius(t *testing.T) {
	c := canvas.New(20, 20)
	Circle(c, 10, 10, -5)

	count := countSetPixels(c, 20, 20)
	if count != 0 {
		t.Errorf("expected 0 pixels for negative radius, got %d", count)
	}
}

func TestCircleFloatCoordinates(t *testing.T) {
	c := canvas.New(20, 20)
	// Circle at (5.7, 6.9) should have center at (5, 6) after flooring
	Circle(c, 5.7, 6.9, 3)

	// East point should be at floored center + radius = (5+3, 6) = (8, 6)
	if !c.Get(8, 6) {
		t.Error("east point (8, 6) not set with floored coordinates")
	}

	// West point should be at (5-3, 6) = (2, 6)
	if !c.Get(2, 6) {
		t.Error("west point (2, 6) not set with floored coordinates")
	}

	printVisual(t, "TestCircleFloatCoordinates", c)
}

func TestCirclePartiallyOffCanvas(t *testing.T) {
	c := canvas.New(20, 20)
	// Circle centered at corner, partially off canvas
	Circle(c, 0, 0, 10)

	// Points on the visible portion should be set
	// East point at (10, 0) should be visible
	if !c.Get(10, 0) {
		t.Error("east point (10, 0) not set")
	}

	// South point at (0, 10) should be visible
	if !c.Get(0, 10) {
		t.Error("south point (0, 10) not set")
	}

	// Points off canvas should be silently clipped (no panic)
	// The circle should still be partially drawn

	printVisual(t, "TestCirclePartiallyOffCanvas", c)
}

func TestCircleFilled(t *testing.T) {
	c := canvas.New(30, 30)
	CircleFilled(c, 14, 14, 10)

	// Center should be set
	if !c.Get(14, 14) {
		t.Error("center pixel (14, 14) not set in filled circle")
	}

	// Cardinal points (on the edge) should be set
	if !c.Get(24, 14) {
		t.Error("east edge (24, 14) not set")
	}
	if !c.Get(4, 14) {
		t.Error("west edge (4, 14) not set")
	}
	if !c.Get(14, 24) {
		t.Error("south edge (14, 24) not set")
	}
	if !c.Get(14, 4) {
		t.Error("north edge (14, 4) not set")
	}

	// Interior points should be set
	if !c.Get(14, 10) {
		t.Error("interior point (14, 10) not set")
	}
	if !c.Get(10, 14) {
		t.Error("interior point (10, 14) not set")
	}

	printVisual(t, "TestCircleFilled", c)
}

func TestCircleFilledNoGaps(t *testing.T) {
	c := canvas.New(30, 30)
	CircleFilled(c, 14, 14, 10)

	// Scan each row to verify no gaps (contiguous fill)
	for y := 4; y <= 24; y++ {
		firstX := -1
		lastX := -1

		for x := 0; x < 30; x++ {
			if c.Get(float64(x), float64(y)) {
				if firstX == -1 {
					firstX = x
				}
				lastX = x
			}
		}

		// If we found any pixels on this row, verify all pixels between first and last are set
		if firstX != -1 {
			for x := firstX; x <= lastX; x++ {
				if !c.Get(float64(x), float64(y)) {
					t.Errorf("gap found at (%d, %d) between first=%d and last=%d", x, y, firstX, lastX)
				}
			}
		}
	}

	printVisual(t, "TestCircleFilledNoGaps", c)
}

func TestCircleOutlineOnly(t *testing.T) {
	c := canvas.New(30, 30)
	Circle(c, 14, 14, 10)

	// Interior should NOT be set
	if c.Get(14, 14) {
		t.Error("center should not be set in outline")
	}

	// More interior points should not be set
	if c.Get(14, 10) {
		t.Error("interior point (14, 10) should not be set in outline")
	}
	if c.Get(10, 14) {
		t.Error("interior point (10, 14) should not be set in outline")
	}

	printVisual(t, "TestCircleOutlineOnly", c)
}

func TestCircleFilledNegativeRadius(t *testing.T) {
	c := canvas.New(20, 20)
	CircleFilled(c, 10, 10, -5)

	count := countSetPixels(c, 20, 20)
	if count != 0 {
		t.Errorf("expected 0 pixels for negative radius, got %d", count)
	}
}

func TestCircleFilledRadius0(t *testing.T) {
	c := canvas.New(10, 10)
	CircleFilled(c, 5, 5, 0)

	// Should draw exactly 1 pixel at center
	if !c.Get(5, 5) {
		t.Error("center pixel (5, 5) not set")
	}

	count := countSetPixels(c, 10, 10)
	if count != 1 {
		t.Errorf("expected 1 pixel for radius 0, got %d", count)
	}

	printVisual(t, "TestCircleFilledRadius0", c)
}

func TestCircleFilledFloatCoordinates(t *testing.T) {
	c := canvas.New(20, 20)
	// Circle at (5.7, 6.9) should have center at (5, 6) after flooring
	CircleFilled(c, 5.7, 6.9, 3)

	// Center should be at floored position
	if !c.Get(5, 6) {
		t.Error("center (5, 6) not set with floored coordinates")
	}

	// Edge points should also be floored
	if !c.Get(8, 6) {
		t.Error("east edge (8, 6) not set with floored coordinates")
	}

	printVisual(t, "TestCircleFilledFloatCoordinates", c)
}
