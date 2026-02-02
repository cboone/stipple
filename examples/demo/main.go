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
	canvasDemo := canvas.New(40, 20)
	fmt.Printf("Canvas: %dx%d pixels (%d cols x %d rows)\n",
		canvasDemo.Width(), canvasDemo.Height(), canvasDemo.Cols(), canvasDemo.Rows())
	fmt.Println()

	demoIndividualPixels()
	demoAllDotsInCell()
	demoDiagonalPattern()
	demoBoxOutline()
	demoInvertedY()
	demoHorizontalLine()
	demoVerticalLine()
	demoDiagonalLine()
	demoShallowSlopeLine()
	demoSteepSlopeLine()
	demoStarPattern()
	demoRectangleOutline()
	demoFilledRectangle()
	demoNestedRectangles()
	demoWallRectangles()
	demoCircleOutline()
	demoFilledCircle()
	demoMultipleCircleSizes()
	demoConcentricCircles()
	demoEyeballSprite()
	demoColorBars()
	demoColoredLines()
	demoColoredRectangles()
	demoColoredEyeball()
}

func demoIndividualPixels() {
	fmt.Println("1. Individual pixels:")
	canvasDemo := canvas.New(4, 8)
	canvasDemo.Set(0, 0) // Top-left dot
	canvasDemo.Set(3, 0) // Top-right dot
	canvasDemo.Set(0, 7) // Bottom-left dot
	canvasDemo.Set(3, 7) // Bottom-right dot
	fmt.Println(canvasDemo.Frame())
	fmt.Println()
}

func demoAllDotsInCell() {
	fmt.Println("2. All 8 dots in a cell:")
	canvasDemo := canvas.New(2, 4)
	for y := 0; y < 4; y++ {
		for x := 0; x < 2; x++ {
			canvasDemo.Set(float64(x), float64(y))
		}
	}
	fmt.Println(canvasDemo.Frame())
	fmt.Println()
	fmt.Println("   Dot positions:    Bit values:")
	fmt.Println("     0  3              0x01  0x08")
	fmt.Println("     1  4              0x02  0x10")
	fmt.Println("     2  5              0x04  0x20")
	fmt.Println("     6  7              0x40  0x80")
	fmt.Println()
}

func demoDiagonalPattern() {
	fmt.Println("3. Diagonal pattern:")
	canvasDemo := canvas.New(20, 20)
	for index := 0; index < 20; index++ {
		canvasDemo.Set(float64(index), float64(index))
	}
	fmt.Println(canvasDemo.Frame())
	fmt.Println()
}

func demoBoxOutline() {
	fmt.Println("4. Box outline:")
	canvasDemo := canvas.New(16, 12)
	// Top and bottom edges
	for x := 0; x < 16; x++ {
		canvasDemo.Set(float64(x), 0)
		canvasDemo.Set(float64(x), 11)
	}
	// Left and right edges
	for y := 0; y < 12; y++ {
		canvasDemo.Set(0, float64(y))
		canvasDemo.Set(15, float64(y))
	}
	fmt.Println(canvasDemo.Frame())
	fmt.Println()
}

func demoInvertedY() {
	fmt.Println("5. Inverted Y-axis (mathematical coordinates):")
	canvasDemo := canvas.New(8, 8, canvas.WithInvertedY())
	// Draw a simple "rising" line from (0,0) to (7,7)
	// With inverted Y, this goes from bottom-left to top-right
	for index := 0; index < 8; index++ {
		canvasDemo.Set(float64(index), float64(index))
	}
	fmt.Println(canvasDemo.Frame())
	fmt.Println("   (origin at bottom-left)")
	fmt.Println()
}

func demoHorizontalLine() {
	fmt.Println("6. Horizontal line (Bresenham):")
	canvasDemo := canvas.New(20, 4)
	draw.Line(canvasDemo, 1, 1, 18, 1)
	fmt.Println(canvasDemo.Frame())
	fmt.Println()
}

func demoVerticalLine() {
	fmt.Println("7. Vertical line:")
	canvasDemo := canvas.New(4, 16)
	draw.Line(canvasDemo, 1, 1, 1, 14)
	fmt.Println(canvasDemo.Frame())
	fmt.Println()
}

func demoDiagonalLine() {
	fmt.Println("8. Diagonal line (45 degrees):")
	canvasDemo := canvas.New(16, 16)
	draw.Line(canvasDemo, 0, 0, 15, 15)
	fmt.Println(canvasDemo.Frame())
	fmt.Println()
}

func demoShallowSlopeLine() {
	fmt.Println("9. Shallow slope line:")
	canvasDemo := canvas.New(20, 8)
	draw.Line(canvasDemo, 0, 1, 19, 5)
	fmt.Println(canvasDemo.Frame())
	fmt.Println()
}

func demoSteepSlopeLine() {
	fmt.Println("10. Steep slope line:")
	canvasDemo := canvas.New(8, 20)
	draw.Line(canvasDemo, 1, 0, 5, 19)
	fmt.Println(canvasDemo.Frame())
	fmt.Println()
}

func demoStarPattern() {
	fmt.Println("11. Star pattern (multiple lines):")
	canvasDemo := canvas.New(20, 20)
	centerX, centerY := 9.0, 9.0
	// Draw 8 lines from center to edges
	draw.Line(canvasDemo, centerX, centerY, 9, 0)   // North
	draw.Line(canvasDemo, centerX, centerY, 19, 0)  // Northeast
	draw.Line(canvasDemo, centerX, centerY, 19, 9)  // East
	draw.Line(canvasDemo, centerX, centerY, 19, 19) // Southeast
	draw.Line(canvasDemo, centerX, centerY, 9, 19)  // South
	draw.Line(canvasDemo, centerX, centerY, 0, 19)  // Southwest
	draw.Line(canvasDemo, centerX, centerY, 0, 9)   // West
	draw.Line(canvasDemo, centerX, centerY, 0, 0)   // Northwest
	fmt.Println(canvasDemo.Frame())
	fmt.Println()
}

func demoRectangleOutline() {
	fmt.Println("12. Rectangle outline:")
	canvasDemo := canvas.New(40, 20)
	draw.Rectangle(canvasDemo, 5, 2, 30, 16)
	fmt.Println(canvasDemo.Frame())
	fmt.Println()
}

func demoFilledRectangle() {
	fmt.Println("13. Filled rectangle:")
	canvasDemo := canvas.New(40, 20)
	draw.RectangleFilled(canvasDemo, 5, 2, 30, 16)
	fmt.Println(canvasDemo.Frame())
	fmt.Println()
}

func demoNestedRectangles() {
	fmt.Println("14. Nested rectangles:")
	canvasDemo := canvas.New(40, 24)
	draw.Rectangle(canvasDemo, 2, 2, 36, 20)
	draw.Rectangle(canvasDemo, 6, 4, 28, 16)
	draw.Rectangle(canvasDemo, 10, 6, 20, 12)
	draw.Rectangle(canvasDemo, 14, 8, 12, 8)
	fmt.Println(canvasDemo.Frame())
	fmt.Println()
}

func demoWallRectangles() {
	fmt.Println("15. Wall-like rectangles (Maze Wars preview):")
	canvasDemo := canvas.New(60, 32)
	// Far wall (small, centered)
	draw.RectangleFilled(canvasDemo, 20, 8, 20, 16)
	// Side walls (trapezoid approximation with rectangles)
	draw.Rectangle(canvasDemo, 5, 2, 50, 28)
	fmt.Println(canvasDemo.Frame())
	fmt.Println()
}

func demoCircleOutline() {
	fmt.Println("16. Circle outline:")
	canvasDemo := canvas.New(30, 28)
	draw.Circle(canvasDemo, 14, 13, 10)
	fmt.Println(canvasDemo.Frame())
	fmt.Println()
}

func demoFilledCircle() {
	fmt.Println("17. Filled circle:")
	canvasDemo := canvas.New(30, 28)
	draw.CircleFilled(canvasDemo, 14, 13, 10)
	fmt.Println(canvasDemo.Frame())
	fmt.Println()
}

func demoMultipleCircleSizes() {
	fmt.Println("18. Multiple circle sizes:")
	canvasDemo := canvas.New(60, 28)
	draw.Circle(canvasDemo, 8, 13, 5)
	draw.Circle(canvasDemo, 25, 13, 8)
	draw.Circle(canvasDemo, 47, 13, 10)
	fmt.Println(canvasDemo.Frame())
	fmt.Println()
}

func demoConcentricCircles() {
	fmt.Println("19. Concentric circles:")
	canvasDemo := canvas.New(40, 36)
	for radius := 2; radius <= 14; radius += 3 {
		draw.Circle(canvasDemo, 19, 17, float64(radius))
	}
	fmt.Println(canvasDemo.Frame())
	fmt.Println()
}

func demoEyeballSprite() {
	fmt.Println("20. Eyeball sprite (Maze Wars preview):")
	canvasDemo := canvas.New(30, 28)
	draw.CircleFilled(canvasDemo, 14, 13, 10) // outer eye
	draw.Circle(canvasDemo, 16, 12, 5)        // iris
	draw.CircleFilled(canvasDemo, 17, 11, 2)  // pupil
	fmt.Println(canvasDemo.Frame())
	fmt.Println()
}

func demoColorBars() {
	fmt.Println("21. Color basics (vertical bars):")
	canvasDemo := canvas.New(32, 8, canvas.WithColor())
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
				canvasDemo.SetColor(float64(x), float64(y), color)
			}
		}
	}
	fmt.Println(canvasDemo.Frame())
	fmt.Println()
}

func demoColoredLines() {
	fmt.Println("22. Colored lines:")
	canvasDemo := canvas.New(20, 16, canvas.WithColor())
	// Red horizontal line
	for x := 0; x < 20; x++ {
		canvasDemo.SetColor(float64(x), 2, canvas.ColorRed)
	}
	// Green diagonal line
	for index := 0; index < 16; index++ {
		x := float64(index * 20 / 16)
		canvasDemo.SetColor(x, float64(index), canvas.ColorGreen)
	}
	// Blue vertical line
	for y := 0; y < 16; y++ {
		canvasDemo.SetColor(18, float64(y), canvas.ColorBlue)
	}
	fmt.Println(canvasDemo.Frame())
	fmt.Println()
}

func demoColoredRectangles() {
	fmt.Println("23. Colored rectangles:")
	canvasDemo := canvas.New(40, 20, canvas.WithColor())
	// Cyan outer rectangle
	for x := 2; x < 38; x++ {
		canvasDemo.SetColor(float64(x), 1, canvas.ColorCyan)
		canvasDemo.SetColor(float64(x), 18, canvas.ColorCyan)
	}
	for y := 1; y < 19; y++ {
		canvasDemo.SetColor(2, float64(y), canvas.ColorCyan)
		canvasDemo.SetColor(37, float64(y), canvas.ColorCyan)
	}
	// Yellow inner rectangle
	for x := 10; x < 30; x++ {
		canvasDemo.SetColor(float64(x), 6, canvas.ColorYellow)
		canvasDemo.SetColor(float64(x), 13, canvas.ColorYellow)
	}
	for y := 6; y < 14; y++ {
		canvasDemo.SetColor(10, float64(y), canvas.ColorYellow)
		canvasDemo.SetColor(29, float64(y), canvas.ColorYellow)
	}
	fmt.Println(canvasDemo.Frame())
	fmt.Println()
}

func demoColoredEyeball() {
	fmt.Println("24. Colored eyeball:")
	canvasDemo := canvas.New(30, 28, canvas.WithColor())
	centerX, centerY := 14.0, 13.0
	// White eye (filled circle)
	for y := 0; y < 28; y++ {
		for x := 0; x < 30; x++ {
			distanceX := float64(x) - centerX
			distanceY := float64(y) - centerY
			distance := math.Sqrt(distanceX*distanceX + distanceY*distanceY)
			if distance <= 10 {
				canvasDemo.SetColor(float64(x), float64(y), canvas.ColorWhite)
			}
		}
	}
	// Blue iris (circle outline)
	irisRadius := 5.0
	for angle := 0.0; angle < 2*math.Pi; angle += 0.05 {
		x := centerX + 2 + irisRadius*math.Cos(angle)
		y := centerY - 1 + irisRadius*math.Sin(angle)
		canvasDemo.SetColor(x, y, canvas.ColorBlue)
	}
	// Black pupil (filled circle)
	pupilCenterX, pupilCenterY := centerX+3, centerY-2
	for y := 0; y < 28; y++ {
		for x := 0; x < 30; x++ {
			distanceX := float64(x) - pupilCenterX
			distanceY := float64(y) - pupilCenterY
			distance := math.Sqrt(distanceX*distanceX + distanceY*distanceY)
			if distance <= 2 {
				canvasDemo.SetColor(float64(x), float64(y), canvas.ColorBlack)
			}
		}
	}
	fmt.Println(canvasDemo.Frame())
}
