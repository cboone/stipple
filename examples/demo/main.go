// Demo program for brodot v0.3.0
package main

import (
	"fmt"

	"github.com/cboone/brodot/canvas"
	"github.com/cboone/brodot/draw"
)

func main() {
	fmt.Println("brodot v0.3.0 Demo")
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
}
