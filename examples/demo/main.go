// Demo program for brodot v0.1.0
package main

import (
	"fmt"

	"github.com/cboone/brodot/canvas"
)

func main() {
	fmt.Println("brodot v0.1.0 Demo")
	fmt.Println("==================")
	fmt.Println()

	// Demo 1: Show canvas dimensions
	c := canvas.New(40, 20)
	fmt.Printf("Canvas: %dx%d pixels (%d cols x %d rows)\n",
		c.Width(), c.Height(), c.Cols(), c.Rows())
	fmt.Println()

	// Demo 2: Individual pixels
	fmt.Println("1. Individual pixels:")
	c1 := canvas.New(4, 8)
	c1.Set(0, 0) // Top-left dot
	c1.Set(3, 0) // Top-right dot
	c1.Set(0, 7) // Bottom-left dot
	c1.Set(3, 7) // Bottom-right dot
	fmt.Println(c1.Frame())
	fmt.Println()

	// Demo 3: All 8 dots in a cell
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

	// Demo 4: Diagonal pattern
	fmt.Println("3. Diagonal pattern:")
	c3 := canvas.New(20, 20)
	for index := 0; index < 20; index++ {
		c3.Set(float64(index), float64(index))
	}
	fmt.Println(c3.Frame())
	fmt.Println()

	// Demo 5: Box pattern
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

	// Demo 6: Inverted Y-axis
	fmt.Println("5. Inverted Y-axis (mathematical coordinates):")
	c5 := canvas.New(8, 8, canvas.WithInvertedY())
	// Draw a simple "rising" line from (0,0) to (7,7)
	// With inverted Y, this goes from bottom-left to top-right
	for index := 0; index < 8; index++ {
		c5.Set(float64(index), float64(index))
	}
	fmt.Println(c5.Frame())
	fmt.Println("   (origin at bottom-left)")
}
