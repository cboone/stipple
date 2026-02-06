// Package draw provides drawing primitives for braille canvases.
package draw

import (
	"math"

	"github.com/cboone/stipple/canvas"
)

// Line draws a line from (startX, startY) to (endX, endY) using Bresenham's algorithm.
func Line(c *canvas.Canvas, startX, startY, endX, endY float64) {
	// Convert float coordinates to int using floor
	x0 := int(math.Floor(startX))
	y0 := int(math.Floor(startY))
	x1 := int(math.Floor(endX))
	y1 := int(math.Floor(endY))

	// Calculate absolute deltas
	dx := x1 - x0
	dy := y1 - y0
	if dx < 0 {
		dx = -dx
	}
	if dy < 0 {
		dy = -dy
	}

	// Determine step directions
	var stepX, stepY int
	if x0 < x1 {
		stepX = 1
	} else {
		stepX = -1
	}
	if y0 < y1 {
		stepY = 1
	} else {
		stepY = -1
	}

	// Determine if line is steep (dy > dx)
	steep := dy > dx

	// Initialize error term
	var err int
	if steep {
		err = dy / 2
	} else {
		err = dx / 2
	}

	// Draw the line
	x, y := x0, y0
	for {
		c.Set(float64(x), float64(y))

		// Check if we've reached the end
		if x == x1 && y == y1 {
			break
		}

		if steep {
			// Steep line: always step in y, sometimes step in x
			y += stepY
			err -= dx
			if err < 0 {
				x += stepX
				err += dy
			}
		} else {
			// Shallow line: always step in x, sometimes step in y
			x += stepX
			err -= dy
			if err < 0 {
				y += stepY
				err += dx
			}
		}
	}
}
