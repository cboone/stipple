package draw

import (
	"math"

	"github.com/cboone/stipple/canvas"
)

// Circle draws a circle outline centered at (centerX, centerY) with the given radius.
// Radius of 0 draws a single pixel at the center.
// Negative radius draws nothing.
func Circle(c *canvas.Canvas, centerX, centerY, radius float64) {
	if radius < 0 {
		return
	}

	intCenterX := int(math.Floor(centerX))
	intCenterY := int(math.Floor(centerY))
	intRadius := int(math.Floor(radius))

	if intRadius == 0 {
		c.Set(float64(intCenterX), float64(intCenterY))
		return
	}

	x := 0
	y := intRadius
	d := 1 - intRadius

	plotCirclePoints(c, intCenterX, intCenterY, x, y)

	for x <= y {
		x++
		if d < 0 {
			d = d + 2*x + 1 // move east
		} else {
			y--
			d = d + 2*(x-y) + 1 // move southeast
		}
		plotCirclePoints(c, intCenterX, intCenterY, x, y)
	}
}

// CircleFilled draws a filled circle centered at (centerX, centerY) with the given radius.
// Radius of 0 draws a single pixel at the center.
// Negative radius draws nothing.
func CircleFilled(c *canvas.Canvas, centerX, centerY, radius float64) {
	if radius < 0 {
		return
	}

	intCenterX := int(math.Floor(centerX))
	intCenterY := int(math.Floor(centerY))
	intRadius := int(math.Floor(radius))

	if intRadius == 0 {
		c.Set(float64(intCenterX), float64(intCenterY))
		return
	}

	x := 0
	y := intRadius
	d := 1 - intRadius

	drawCircleSpans(c, intCenterX, intCenterY, x, y)

	for x <= y {
		x++
		if d < 0 {
			d = d + 2*x + 1 // move east
		} else {
			y--
			d = d + 2*(x-y) + 1 // move southeast
		}
		drawCircleSpans(c, intCenterX, intCenterY, x, y)
	}
}

// plotCirclePoints plots all 8 symmetric points for the circle outline.
func plotCirclePoints(c *canvas.Canvas, centerX, centerY, x, y int) {
	c.Set(float64(centerX+x), float64(centerY+y))
	c.Set(float64(centerX-x), float64(centerY+y))
	c.Set(float64(centerX+x), float64(centerY-y))
	c.Set(float64(centerX-x), float64(centerY-y))
	c.Set(float64(centerX+y), float64(centerY+x))
	c.Set(float64(centerX-y), float64(centerY+x))
	c.Set(float64(centerX+y), float64(centerY-x))
	c.Set(float64(centerX-y), float64(centerY-x))
}

// drawCircleSpans draws 4 horizontal spans covering all octants for filled circles.
func drawCircleSpans(c *canvas.Canvas, centerX, centerY, x, y int) {
	drawHorizontalSpan(c, centerX-x, centerX+x, centerY+y)
	drawHorizontalSpan(c, centerX-x, centerX+x, centerY-y)
	drawHorizontalSpan(c, centerX-y, centerX+y, centerY+x)
	drawHorizontalSpan(c, centerX-y, centerX+y, centerY-x)
}

// drawHorizontalSpan draws a horizontal line from startX to endX at the given y.
func drawHorizontalSpan(c *canvas.Canvas, startX, endX, y int) {
	for pixelX := startX; pixelX <= endX; pixelX++ {
		c.Set(float64(pixelX), float64(y))
	}
}
