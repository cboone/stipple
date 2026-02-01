package draw

import (
	"math"

	"github.com/cboone/brodot/canvas"
)

// Rectangle draws a rectangle outline from (x, y) with the given width and height.
// The rectangle's top-left corner is at (x, y), extending to (x+width-1, y+height-1).
// Width or height of 0 or negative draws nothing.
func Rectangle(c *canvas.Canvas, x, y, width, height float64) {
	if width <= 0 || height <= 0 {
		return
	}

	// Calculate corner coordinates
	// Top-left: (x, y)
	// Top-right: (x + width - 1, y)
	// Bottom-left: (x, y + height - 1)
	// Bottom-right: (x + width - 1, y + height - 1)
	right := x + width - 1
	bottom := y + height - 1

	// Draw four edges using Line
	Line(c, x, y, right, y)           // Top edge
	Line(c, right, y, right, bottom)  // Right edge
	Line(c, right, bottom, x, bottom) // Bottom edge
	Line(c, x, bottom, x, y)          // Left edge
}

// RectangleFilled draws a filled rectangle from (x, y) with the given width and height.
// The rectangle's top-left corner is at (x, y), extending to (x+width-1, y+height-1).
// Width or height of 0 or negative draws nothing.
func RectangleFilled(c *canvas.Canvas, x, y, width, height float64) {
	if width <= 0 || height <= 0 {
		return
	}

	// Convert to integers for iteration
	startX := int(math.Floor(x))
	startY := int(math.Floor(y))
	endX := int(math.Floor(x + width - 1))  // inclusive
	endY := int(math.Floor(y + height - 1)) // inclusive

	// Set all pixels in the rectangle
	for pixelY := startY; pixelY <= endY; pixelY++ {
		for pixelX := startX; pixelX <= endX; pixelX++ {
			c.Set(float64(pixelX), float64(pixelY))
		}
	}
}
