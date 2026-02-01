package canvas

import (
	"flag"
	"os"
	"testing"
)

var visualFlag = flag.Bool("visual", false, "print visual output")

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

func printVisual(t *testing.T, name string, canvas *Canvas) {
	if *visualFlag {
		t.Logf("\n=== %s ===\n%s", name, canvas.Frame())
	}
}

func TestNew(t *testing.T) {
	canvas := New(4, 8)

	// Verify dimensions
	if canvas.Width() != 4 {
		t.Errorf("Width() = %d, want 4", canvas.Width())
	}
	if canvas.Height() != 8 {
		t.Errorf("Height() = %d, want 8", canvas.Height())
	}
	if canvas.Cols() != 2 {
		t.Errorf("Cols() = %d, want 2", canvas.Cols())
	}
	if canvas.Rows() != 2 {
		t.Errorf("Rows() = %d, want 2", canvas.Rows())
	}

	// Verify all cells are initialized to BrailleOffset
	for row := 0; row < canvas.Rows(); row++ {
		for column := 0; column < canvas.Cols(); column++ {
			if canvas.cells[row][column] != BrailleOffset {
				t.Errorf("cells[%d][%d] = %#x, want %#x", row, column, canvas.cells[row][column], BrailleOffset)
			}
		}
	}

	printVisual(t, "TestNew", canvas)
}

func TestSetGet(t *testing.T) {
	canvas := New(4, 8)

	// Initially pixel should be off
	if canvas.Get(0, 0) {
		t.Error("Get(0, 0) = true before Set, want false")
	}

	// Set the pixel
	canvas.Set(0, 0)

	// Now it should be on
	if !canvas.Get(0, 0) {
		t.Error("Get(0, 0) = false after Set, want true")
	}

	printVisual(t, "TestSetGet", canvas)
}

func TestUnset(t *testing.T) {
	canvas := New(4, 8)

	// Set then unset
	canvas.Set(0, 0)
	canvas.Unset(0, 0)

	if canvas.Get(0, 0) {
		t.Error("Get(0, 0) = true after Unset, want false")
	}

	printVisual(t, "TestUnset", canvas)
}

func TestToggle(t *testing.T) {
	canvas := New(4, 8)

	// Toggle on
	canvas.Toggle(0, 0)
	if !canvas.Get(0, 0) {
		t.Error("Get(0, 0) = false after first Toggle, want true")
	}

	// Toggle off
	canvas.Toggle(0, 0)
	if canvas.Get(0, 0) {
		t.Error("Get(0, 0) = true after second Toggle, want false")
	}

	printVisual(t, "TestToggle", canvas)
}

func TestClear(t *testing.T) {
	canvas := New(4, 8)

	// Set some pixels
	canvas.Set(0, 0)
	canvas.Set(1, 1)
	canvas.Set(2, 2)

	// Clear
	canvas.Clear()

	// Verify all pixels are off
	for x := 0; x < canvas.Width(); x++ {
		for y := 0; y < canvas.Height(); y++ {
			if canvas.Get(float64(x), float64(y)) {
				t.Errorf("Get(%d, %d) = true after Clear, want false", x, y)
			}
		}
	}

	printVisual(t, "TestClear", canvas)
}

func TestFrame(t *testing.T) {
	canvas := New(4, 8)

	// Set pixel (0,0) -> cell (0,0), dot (0,0) -> bit 0x01
	canvas.Set(0, 0)
	// Set pixel (3,0) -> cell (0,1), dot (0,1) -> bit 0x08
	canvas.Set(3, 0)
	// Set pixel (0,5) -> cell (1,0), dot (1,0) -> bit 0x02
	canvas.Set(0, 5)
	// Set pixel (3,5) -> cell (1,1), dot (1,1) -> bit 0x10
	canvas.Set(3, 5)

	frame := canvas.Frame()

	// First row: cells with dots 0 and 3 set
	// Second row: cells with dots 1 and 4 set
	expected := string([]rune{
		BrailleOffset | 0x01, BrailleOffset | 0x08,
	}) + "\n" + string([]rune{
		BrailleOffset | 0x02, BrailleOffset | 0x10,
	})

	if frame != expected {
		t.Errorf("Frame() =\n%s\nwant:\n%s", frame, expected)
	}

	printVisual(t, "TestFrame", canvas)
}

func TestOutOfBounds(t *testing.T) {
	canvas := New(4, 8)

	// These should not panic
	canvas.Set(100, 100)
	canvas.Set(float64(canvas.Width()), 0)
	canvas.Set(0, float64(canvas.Height()))

	// Get should return false for out of bounds
	if canvas.Get(100, 100) {
		t.Error("Get(100, 100) = true for out of bounds, want false")
	}

	printVisual(t, "TestOutOfBounds", canvas)
}

func TestOutOfBoundsNegative(t *testing.T) {
	canvas := New(4, 8)

	// Negative coordinates should not panic
	canvas.Set(-1, 0)
	canvas.Set(0, -1)
	canvas.Set(-1, -1)
	canvas.Set(-0.5, -0.5)

	// Get should return false for negative coordinates
	if canvas.Get(-1, 0) {
		t.Error("Get(-1, 0) = true for negative, want false")
	}
	if canvas.Get(0, -1) {
		t.Error("Get(0, -1) = true for negative, want false")
	}

	printVisual(t, "TestOutOfBoundsNegative", canvas)
}

func TestInvertedY(t *testing.T) {
	canvas := New(4, 8, WithInvertedY())

	// With inverted Y, pixel (0, 0) should map to bottom-left
	// In a 4x8 canvas with 2 rows, pixel (0, 0) with inverted Y
	// becomes pixel (0, 7) in screen coords
	canvas.Set(0, 0)

	// The pixel should appear in the bottom row (row 1), left cell (col 0)
	// Position (0, 7) -> cell row 1, dot row 3, col 0, dot col 0
	// That's dot 6 (bit 0x40)
	expectedCell := BrailleOffset | 0x40
	if canvas.cells[1][0] != expectedCell {
		t.Errorf("cells[1][0] = %#x, want %#x", canvas.cells[1][0], expectedCell)
	}

	printVisual(t, "TestInvertedY", canvas)
}

func TestDimensions(t *testing.T) {
	tests := []struct {
		width  int
		height int
		rows   int
		cols   int
	}{
		{4, 8, 2, 2},
		{10, 20, 5, 5},
		{80, 40, 10, 40},
		{5, 9, 2, 2},   // Tests truncation: 9/4 = 2, 5/2 = 2
		{100, 100, 25, 50},
	}

	for _, testCase := range tests {
		canvas := New(testCase.width, testCase.height)

		if canvas.Width() != testCase.width {
			t.Errorf("New(%d, %d).Width() = %d, want %d",
				testCase.width, testCase.height, canvas.Width(), testCase.width)
		}
		if canvas.Height() != testCase.height {
			t.Errorf("New(%d, %d).Height() = %d, want %d",
				testCase.width, testCase.height, canvas.Height(), testCase.height)
		}
		if canvas.Rows() != testCase.rows {
			t.Errorf("New(%d, %d).Rows() = %d, want %d",
				testCase.width, testCase.height, canvas.Rows(), testCase.rows)
		}
		if canvas.Cols() != testCase.cols {
			t.Errorf("New(%d, %d).Cols() = %d, want %d",
				testCase.width, testCase.height, canvas.Cols(), testCase.cols)
		}
	}
}

func TestFrameMultipleRows(t *testing.T) {
	canvas := New(4, 16) // 2 cols x 4 rows

	// Set one pixel in each row
	canvas.Set(0, 0)   // row 0
	canvas.Set(0, 4)   // row 1
	canvas.Set(0, 8)   // row 2
	canvas.Set(0, 12)  // row 3

	frame := canvas.Frame()

	// Each row should have the first cell with dot 0 (bit 0x01)
	lines := []string{
		string([]rune{BrailleOffset | 0x01, BrailleOffset}),
		string([]rune{BrailleOffset | 0x01, BrailleOffset}),
		string([]rune{BrailleOffset | 0x01, BrailleOffset}),
		string([]rune{BrailleOffset | 0x01, BrailleOffset}),
	}
	expected := lines[0] + "\n" + lines[1] + "\n" + lines[2] + "\n" + lines[3]

	if frame != expected {
		t.Errorf("Frame() =\n%s\nwant:\n%s", frame, expected)
	}

	printVisual(t, "TestFrameMultipleRows", canvas)
}

func TestFloatCoordinates(t *testing.T) {
	canvas := New(4, 8)

	// Test that fractional coordinates floor correctly
	canvas.Set(0.9, 0.9)  // Should map to (0, 0)
	if !canvas.Get(0, 0) {
		t.Error("Set(0.9, 0.9) should set pixel (0, 0)")
	}

	canvas.Set(1.5, 2.7)  // Should map to (1, 2)
	if !canvas.Get(1, 2) {
		t.Error("Set(1.5, 2.7) should set pixel (1, 2)")
	}

	printVisual(t, "TestFloatCoordinates", canvas)
}

func TestDimensionTruncationBounds(t *testing.T) {
	// Test that pixels in truncated regions are correctly rejected.
	// With width=5, height=9:
	// - Cols() = 5/2 = 2 (valid column indices: 0, 1)
	// - Rows() = 9/4 = 2 (valid row indices: 0, 1)
	// - Pixel (4, 8) passes pixel bounds (4 < 5, 8 < 9)
	// - But cellCol = 4/2 = 2, cellRow = 8/4 = 2 are out of cell bounds
	canvas := New(5, 9)

	// These should not panic - they're in valid pixel space but invalid cell space
	canvas.Set(4, 8)
	canvas.Set(4, 0)
	canvas.Set(0, 8)

	// Get should return false for pixels that map to truncated cells
	if canvas.Get(4, 8) {
		t.Error("Get(4, 8) = true for truncated cell region, want false")
	}
	if canvas.Get(4, 0) {
		t.Error("Get(4, 0) = true for truncated cell region, want false")
	}
	if canvas.Get(0, 8) {
		t.Error("Get(0, 8) = true for truncated cell region, want false")
	}

	// But valid pixels should still work
	canvas.Set(3, 7)
	if !canvas.Get(3, 7) {
		t.Error("Get(3, 7) = false for valid pixel, want true")
	}

	printVisual(t, "TestDimensionTruncationBounds", canvas)
}
