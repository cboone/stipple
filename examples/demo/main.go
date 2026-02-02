// Demo program for brodot v0.5.0
package main

import (
	"fmt"
	"math"

	"github.com/cboone/brodot/canvas"
	"github.com/cboone/brodot/draw"
)

func main() {
	fmt.Println("brodot v0.5.0 Demo")
	fmt.Println("==================")
	fmt.Println()

	// Show canvas dimensions
	c := canvas.New(40, 20)
	fmt.Printf("Canvas: %dx%d pixels (%d cols x %d rows)\n",
		c.Width(), c.Height(), c.Cols(), c.Rows())
	fmt.Println()

	// Demo 1: Individual pixels
	fmt.Println("1. Individual pixels:")
	c1 := canvas.New(4, 8)
	c1.Set(0, 0) // Top-left dot
	c1.Set(3, 0) // Top-right dot
	c1.Set(0, 7) // Bottom-left dot
	c1.Set(3, 7) // Bottom-right dot
	fmt.Println(c1.Frame())
	fmt.Println()

	// Demo 2: All 8 dots in a cell
	fmt.Println("2. All 8 dots in a cell:")
	c2 := canvas.New(2, 4)
	for y := 0; y < 4; y++ {
		for x := 0; x < 2; x++ {
			c2.Set(float64(x), float64(y))
		}
	}
	fmt.Println(c2.Frame())
	fmt.Println()
	fmt.Println("   Dot positions:    Bit values:")
	fmt.Println("     0  3              0x01  0x08")
	fmt.Println("     1  4              0x02  0x10")
	fmt.Println("     2  5              0x04  0x20")
	fmt.Println("     6  7              0x40  0x80")
	fmt.Println()

	// Demo 3: Diagonal pattern
	fmt.Println("3. Diagonal pattern:")
	c3 := canvas.New(20, 20)
	for index := 0; index < 20; index++ {
		c3.Set(float64(index), float64(index))
	}
	fmt.Println(c3.Frame())
	fmt.Println()

	// Demo 4: Box pattern
	fmt.Println("4. Box outline:")
	c4 := canvas.New(16, 12)
	// Top and bottom edges
	for x := 0; x < 16; x++ {
		c4.Set(float64(x), 0)
		c4.Set(float64(x), 11)
	}
	// Left and right edges
	for y := 0; y < 12; y++ {
		c4.Set(0, float64(y))
		c4.Set(15, float64(y))
	}
	fmt.Println(c4.Frame())
	fmt.Println()

	// Demo 5: Inverted Y-axis
	fmt.Println("5. Inverted Y-axis (mathematical coordinates):")
	c5 := canvas.New(8, 8, canvas.WithInvertedY())
	// Draw a simple "rising" line from (0,0) to (7,7)
	// With inverted Y, this goes from bottom-left to top-right
	for index := 0; index < 8; index++ {
		c5.Set(float64(index), float64(index))
	}
	fmt.Println(c5.Frame())
	fmt.Println("   (origin at bottom-left)")
	fmt.Println()

	// Demo 6: Horizontal line using draw.Line
	fmt.Println("6. Horizontal line (Bresenham):")
	c6 := canvas.New(20, 4)
	draw.Line(c6, 1, 1, 18, 1)
	fmt.Println(c6.Frame())
	fmt.Println()

	// Demo 7: Vertical line
	fmt.Println("7. Vertical line:")
	c7 := canvas.New(4, 16)
	draw.Line(c7, 1, 1, 1, 14)
	fmt.Println(c7.Frame())
	fmt.Println()

	// Demo 8: Diagonal line (45 degrees)
	fmt.Println("8. Diagonal line (45 degrees):")
	c8 := canvas.New(16, 16)
	draw.Line(c8, 0, 0, 15, 15)
	fmt.Println(c8.Frame())
	fmt.Println()

	// Demo 9: Shallow slope line
	fmt.Println("9. Shallow slope line:")
	c9 := canvas.New(20, 8)
	draw.Line(c9, 0, 1, 19, 5)
	fmt.Println(c9.Frame())
	fmt.Println()

	// Demo 10: Steep slope line
	fmt.Println("10. Steep slope line:")
	c10 := canvas.New(8, 20)
	draw.Line(c10, 1, 0, 5, 19)
	fmt.Println(c10.Frame())
	fmt.Println()

	// Demo 11: Star pattern with multiple lines
	fmt.Println("11. Star pattern (multiple lines):")
	c11 := canvas.New(20, 20)
	centerX, centerY := 9.0, 9.0
	// Draw 8 lines from center to edges
	draw.Line(c11, centerX, centerY, 9, 0)   // North
	draw.Line(c11, centerX, centerY, 19, 0)  // Northeast
	draw.Line(c11, centerX, centerY, 19, 9)  // East
	draw.Line(c11, centerX, centerY, 19, 19) // Southeast
	draw.Line(c11, centerX, centerY, 9, 19)  // South
	draw.Line(c11, centerX, centerY, 0, 19)  // Southwest
	draw.Line(c11, centerX, centerY, 0, 9)   // West
	draw.Line(c11, centerX, centerY, 0, 0)   // Northwest
	fmt.Println(c11.Frame())
	fmt.Println()

	// Demo 12: Rectangle outline
	fmt.Println("12. Rectangle outline:")
	c12 := canvas.New(40, 20)
	draw.Rectangle(c12, 5, 2, 30, 16)
	fmt.Println(c12.Frame())
	fmt.Println()

	// Demo 13: Filled rectangle
	fmt.Println("13. Filled rectangle:")
	c13 := canvas.New(40, 20)
	draw.RectangleFilled(c13, 5, 2, 30, 16)
	fmt.Println(c13.Frame())
	fmt.Println()

	// Demo 14: Multiple rectangles (nested)
	fmt.Println("14. Nested rectangles:")
	c14 := canvas.New(40, 24)
	draw.Rectangle(c14, 2, 2, 36, 20)
	draw.Rectangle(c14, 6, 4, 28, 16)
	draw.Rectangle(c14, 10, 6, 20, 12)
	draw.Rectangle(c14, 14, 8, 12, 8)
	fmt.Println(c14.Frame())
	fmt.Println()

	// Demo 15: Wall-like rectangles (Maze Wars preview)
	fmt.Println("15. Wall-like rectangles (Maze Wars preview):")
	c15 := canvas.New(60, 32)
	// Far wall (small, centered)
	draw.RectangleFilled(c15, 20, 8, 20, 16)
	// Side walls (trapezoid approximation with rectangles)
	draw.Rectangle(c15, 5, 2, 50, 28)
	fmt.Println(c15.Frame())
	fmt.Println()

	// Demo 16: Circle outline
	fmt.Println("16. Circle outline:")
	c16 := canvas.New(30, 28)
	draw.Circle(c16, 14, 13, 10)
	fmt.Println(c16.Frame())
	fmt.Println()

	// Demo 17: Filled circle
	fmt.Println("17. Filled circle:")
	c17 := canvas.New(30, 28)
	draw.CircleFilled(c17, 14, 13, 10)
	fmt.Println(c17.Frame())
	fmt.Println()

	// Demo 18: Multiple circle sizes
	fmt.Println("18. Multiple circle sizes:")
	c18 := canvas.New(60, 28)
	draw.Circle(c18, 8, 13, 5)
	draw.Circle(c18, 25, 13, 8)
	draw.Circle(c18, 47, 13, 10)
	fmt.Println(c18.Frame())
	fmt.Println()

	// Demo 19: Concentric circles
	fmt.Println("19. Concentric circles:")
	c19 := canvas.New(40, 36)
	for radius := 2; radius <= 14; radius += 3 {
		draw.Circle(c19, 19, 17, float64(radius))
	}
	fmt.Println(c19.Frame())
	fmt.Println()

	// Demo 20: Eyeball sprite (Maze Wars preview)
	fmt.Println("20. Eyeball sprite (Maze Wars preview):")
	c20 := canvas.New(30, 28)
	draw.CircleFilled(c20, 14, 13, 10) // outer eye
	draw.Circle(c20, 16, 12, 5)        // iris
	draw.CircleFilled(c20, 17, 11, 2)  // pupil
	fmt.Println(c20.Frame())
	fmt.Println()

	// Demo 21: Color basics - Vertical bars of each color
	fmt.Println("21. Color basics (vertical bars):")
	c21 := canvas.New(32, 8, canvas.WithColor())
	colors := []canvas.Color{
		canvas.ColorBlack,
		canvas.ColorBlue,
		canvas.ColorCyan,
		canvas.ColorGreen,
		canvas.ColorMagenta,
		canvas.ColorRed,
		canvas.ColorWhite,
		canvas.ColorYellow,
	}
	for index, color := range colors {
		startX := index * 4
		for x := startX; x < startX+4; x++ {
			for y := 0; y < 8; y++ {
				c21.SetColor(float64(x), float64(y), color)
			}
		}
	}
	fmt.Println(c21.Frame())
	fmt.Println()

	// Demo 22: Colored lines - Red horizontal, green diagonal, blue vertical
	fmt.Println("22. Colored lines:")
	c22 := canvas.New(20, 16, canvas.WithColor())
	// Red horizontal line
	for x := 0; x < 20; x++ {
		c22.SetColor(float64(x), 2, canvas.ColorRed)
	}
	// Green diagonal line
	for index := 0; index < 16; index++ {
		x := float64(index * 20 / 16)
		c22.SetColor(x, float64(index), canvas.ColorGreen)
	}
	// Blue vertical line
	for y := 0; y < 16; y++ {
		c22.SetColor(18, float64(y), canvas.ColorBlue)
	}
	fmt.Println(c22.Frame())
	fmt.Println()

	// Demo 23: Colored rectangles - Cyan outer, yellow inner nested rectangles
	fmt.Println("23. Colored rectangles:")
	c23 := canvas.New(40, 20, canvas.WithColor())
	// Cyan outer rectangle
	for x := 2; x < 38; x++ {
		c23.SetColor(float64(x), 1, canvas.ColorCyan)
		c23.SetColor(float64(x), 18, canvas.ColorCyan)
	}
	for y := 1; y < 19; y++ {
		c23.SetColor(2, float64(y), canvas.ColorCyan)
		c23.SetColor(37, float64(y), canvas.ColorCyan)
	}
	// Yellow inner rectangle
	for x := 10; x < 30; x++ {
		c23.SetColor(float64(x), 6, canvas.ColorYellow)
		c23.SetColor(float64(x), 13, canvas.ColorYellow)
	}
	for y := 6; y < 14; y++ {
		c23.SetColor(10, float64(y), canvas.ColorYellow)
		c23.SetColor(29, float64(y), canvas.ColorYellow)
	}
	fmt.Println(c23.Frame())
	fmt.Println()

	// Demo 24: Colored eyeball - White eye, blue iris, black pupil
	fmt.Println("24. Colored eyeball:")
	c24 := canvas.New(30, 28, canvas.WithColor())
	centerX24, centerY24 := 14.0, 13.0
	// White eye (filled circle)
	for y := 0; y < 28; y++ {
		for x := 0; x < 30; x++ {
			distanceX := float64(x) - centerX24
			distanceY := float64(y) - centerY24
			distance := math.Sqrt(distanceX*distanceX + distanceY*distanceY)
			if distance <= 10 {
				c24.SetColor(float64(x), float64(y), canvas.ColorWhite)
			}
		}
	}
	// Blue iris (circle outline)
	irisRadius := 5.0
	for angle := 0.0; angle < 2*math.Pi; angle += 0.05 {
		x := centerX24 + 2 + irisRadius*math.Cos(angle)
		y := centerY24 - 1 + irisRadius*math.Sin(angle)
		c24.SetColor(x, y, canvas.ColorBlue)
	}
	// Black pupil (filled circle)
	pupilCenterX, pupilCenterY := centerX24+3, centerY24-2
	for y := 0; y < 28; y++ {
		for x := 0; x < 30; x++ {
			distanceX := float64(x) - pupilCenterX
			distanceY := float64(y) - pupilCenterY
			distance := math.Sqrt(distanceX*distanceX + distanceY*distanceY)
			if distance <= 2 {
				c24.SetColor(float64(x), float64(y), canvas.ColorBlack)
			}
		}
	}
	fmt.Println(c24.Frame())
}
