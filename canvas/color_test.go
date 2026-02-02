package canvas

import (
	"strings"
	"testing"
)

func TestColorANSI(t *testing.T) {
	tests := []struct {
		color    Color
		expected string
	}{
		{ColorDefault, ""},
		{ColorBlack, "\x1b[30m"},
		{ColorBlue, "\x1b[34m"},
		{ColorCyan, "\x1b[36m"},
		{ColorGreen, "\x1b[32m"},
		{ColorMagenta, "\x1b[35m"},
		{ColorRed, "\x1b[31m"},
		{ColorWhite, "\x1b[37m"},
		{ColorYellow, "\x1b[33m"},
	}

	for _, testCase := range tests {
		result := testCase.color.ANSI()
		if result != testCase.expected {
			t.Errorf("Color(%d).ANSI() = %q, want %q", testCase.color, result, testCase.expected)
		}
	}
}

func TestColorANSIOutOfBounds(t *testing.T) {
	// Invalid color values should return empty string
	invalidColor := Color(255)
	result := invalidColor.ANSI()
	if result != "" {
		t.Errorf("Color(255).ANSI() = %q, want empty string", result)
	}
}

func TestANSIReset(t *testing.T) {
	expected := "\x1b[0m"
	result := ANSIReset()
	if result != expected {
		t.Errorf("ANSIReset() = %q, want %q", result, expected)
	}
}

func TestSetColorWithColorEnabled(t *testing.T) {
	canvas := New(4, 8, WithColor())

	canvas.SetColor(0, 0, ColorRed)

	// Verify pixel is set
	if !canvas.Get(0, 0) {
		t.Error("Get(0, 0) = false after SetColor, want true")
	}

	// Verify color is stored
	if canvas.colors[0][0] != ColorRed {
		t.Errorf("colors[0][0] = %d, want %d (ColorRed)", canvas.colors[0][0], ColorRed)
	}

	printVisual(t, "TestSetColorWithColorEnabled", canvas)
}

func TestSetColorWithoutColorEnabled(t *testing.T) {
	canvas := New(4, 8) // No WithColor()

	canvas.SetColor(0, 0, ColorRed)

	// Verify pixel is set
	if !canvas.Get(0, 0) {
		t.Error("Get(0, 0) = false after SetColor, want true")
	}

	// Verify colors grid is nil
	if canvas.colors != nil {
		t.Error("colors should be nil when WithColor() not used")
	}

	printVisual(t, "TestSetColorWithoutColorEnabled", canvas)
}

func TestColorLastWriteWins(t *testing.T) {
	canvas := New(4, 8, WithColor())

	// Set multiple pixels in the same cell with different colors
	canvas.SetColor(0, 0, ColorRed)
	canvas.SetColor(1, 0, ColorGreen)
	canvas.SetColor(0, 1, ColorBlue)

	// The last write should win
	if canvas.colors[0][0] != ColorBlue {
		t.Errorf("colors[0][0] = %d, want %d (ColorBlue) - last write wins",
			canvas.colors[0][0], ColorBlue)
	}

	printVisual(t, "TestColorLastWriteWins", canvas)
}

func TestColorReset(t *testing.T) {
	canvas := New(4, 8, WithColor())
	canvas.SetColor(0, 0, ColorRed)
	canvas.SetColor(2, 0, ColorGreen)

	frame := canvas.Frame()

	// Count resets - should have one reset per colored cell
	resetCount := strings.Count(frame, ANSIReset())
	if resetCount != 2 {
		t.Errorf("Frame has %d resets, want 2", resetCount)
	}

	printVisual(t, "TestColorReset", canvas)
}

func TestColorClear(t *testing.T) {
	canvas := New(4, 8, WithColor())

	canvas.SetColor(0, 0, ColorRed)
	canvas.SetColor(2, 4, ColorBlue)

	canvas.Clear()

	// Verify pixels are cleared
	if canvas.Get(0, 0) {
		t.Error("Get(0, 0) = true after Clear, want false")
	}

	// Verify colors are reset to default
	if canvas.colors[0][0] != ColorDefault {
		t.Errorf("colors[0][0] = %d after Clear, want %d (ColorDefault)",
			canvas.colors[0][0], ColorDefault)
	}
	if canvas.colors[1][1] != ColorDefault {
		t.Errorf("colors[1][1] = %d after Clear, want %d (ColorDefault)",
			canvas.colors[1][1], ColorDefault)
	}

	printVisual(t, "TestColorClear", canvas)
}

func TestColorOutOfBounds(t *testing.T) {
	canvas := New(4, 8, WithColor())

	// These should not panic
	canvas.SetColor(100, 100, ColorRed)
	canvas.SetColor(-1, -1, ColorBlue)
	canvas.SetColor(float64(canvas.Width()), 0, ColorGreen)

	printVisual(t, "TestColorOutOfBounds", canvas)
}

func TestFrameWithMixedColors(t *testing.T) {
	canvas := New(4, 8, WithColor())

	canvas.SetColor(0, 0, ColorRed)
	canvas.SetColor(2, 0, ColorGreen)
	canvas.Set(0, 4) // No color - just set pixel

	frame := canvas.Frame()

	// Check that red escape code is present
	if !strings.Contains(frame, ColorRed.ANSI()) {
		t.Error("Frame should contain red ANSI code")
	}

	// Check that green escape code is present
	if !strings.Contains(frame, ColorGreen.ANSI()) {
		t.Error("Frame should contain green ANSI code")
	}

	printVisual(t, "TestFrameWithMixedColors", canvas)
}

func TestFrameNoColorsWhenDisabled(t *testing.T) {
	canvas := New(4, 8) // No WithColor()

	canvas.Set(0, 0)
	canvas.Set(2, 0)

	frame := canvas.Frame()

	// Should contain no escape codes
	if strings.Contains(frame, "\x1b[") {
		t.Error("Frame should not contain ANSI codes when color is disabled")
	}

	printVisual(t, "TestFrameNoColorsWhenDisabled", canvas)
}

func TestWithColorAndWithInvertedY(t *testing.T) {
	canvas := New(4, 8, WithColor(), WithInvertedY())

	// With inverted Y, pixel (0, 0) maps to bottom-left
	canvas.SetColor(0, 0, ColorRed)

	// The color should be in the bottom row
	if canvas.colors[1][0] != ColorRed {
		t.Errorf("colors[1][0] = %d, want %d (ColorRed) with inverted Y",
			canvas.colors[1][0], ColorRed)
	}

	// Verify pixel is also set correctly
	if !canvas.Get(0, 0) {
		t.Error("Get(0, 0) = false after SetColor with inverted Y, want true")
	}

	printVisual(t, "TestWithColorAndWithInvertedY", canvas)
}

func TestColorDefaultNotRendered(t *testing.T) {
	canvas := New(4, 8, WithColor())

	// Set a pixel without color (uses default)
	canvas.Set(0, 0)

	frame := canvas.Frame()

	// Default color should not emit any ANSI codes
	if strings.Contains(frame, "\x1b[") {
		t.Error("Frame should not contain ANSI codes for ColorDefault cells")
	}

	printVisual(t, "TestColorDefaultNotRendered", canvas)
}

func TestColorGridAllocation(t *testing.T) {
	// Without WithColor, colors should be nil
	canvasNoColor := New(4, 8)
	if canvasNoColor.colors != nil {
		t.Error("colors should be nil without WithColor()")
	}

	// With WithColor, colors should be allocated
	canvasWithColor := New(4, 8, WithColor())
	if canvasWithColor.colors == nil {
		t.Error("colors should not be nil with WithColor()")
	}

	// Verify dimensions match cells
	if len(canvasWithColor.colors) != canvasWithColor.Rows() {
		t.Errorf("colors rows = %d, want %d",
			len(canvasWithColor.colors), canvasWithColor.Rows())
	}
	if len(canvasWithColor.colors[0]) != canvasWithColor.Cols() {
		t.Errorf("colors cols = %d, want %d",
			len(canvasWithColor.colors[0]), canvasWithColor.Cols())
	}
}

func TestColorVisualDemo(t *testing.T) {
	if !*visualFlag {
		t.Skip("Skipping visual demo (use -visual flag)")
	}

	canvas := New(20, 8, WithColor())

	// Draw vertical bars of different colors
	colors := []Color{ColorRed, ColorGreen, ColorBlue, ColorYellow, ColorCyan, ColorMagenta}
	for index, color := range colors {
		x := float64(index * 3)
		for y := 0; y < 8; y++ {
			canvas.SetColor(x, float64(y), color)
			canvas.SetColor(x+1, float64(y), color)
		}
	}

	t.Logf("\n=== Color Visual Demo ===\n%s", canvas.Frame())
}
