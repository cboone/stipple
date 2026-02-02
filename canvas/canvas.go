package canvas

import (
	"math"
	"strings"
)

// Canvas represents a braille graphics canvas.
// Each terminal cell displays a 2x4 braille pattern, providing pixel-level control.
type Canvas struct {
	cells        [][]rune   // braille character grid [row][col]
	colorEnabled bool       // whether color support is enabled
	colors       [][]Color  // color grid [row][col], nil when colors disabled
	height       int        // pixel height
	invertY      bool       // Y-axis direction: false = down, true = up
	width        int        // pixel width
}

// New creates a new Canvas with the specified pixel dimensions.
// The canvas is initialized with all cells set to the empty braille pattern.
func New(width, height int, options ...Option) *Canvas {
	canvas := &Canvas{
		width:  width,
		height: height,
	}

	// Apply options
	for _, option := range options {
		option(canvas)
	}

	// Allocate cells grid
	rows := canvas.Rows()
	columns := canvas.Cols()
	canvas.cells = make([][]rune, rows)
	for row := range canvas.cells {
		canvas.cells[row] = make([]rune, columns)
		for column := range canvas.cells[row] {
			canvas.cells[row][column] = BrailleOffset
		}
	}

	// Allocate colors grid when color support is enabled
	if canvas.colorEnabled {
		canvas.colors = make([][]Color, rows)
		for row := range canvas.colors {
			canvas.colors[row] = make([]Color, columns)
		}
	}

	return canvas
}

// Width returns the pixel width of the canvas.
func (canvas *Canvas) Width() int {
	return canvas.width
}

// Height returns the pixel height of the canvas.
func (canvas *Canvas) Height() int {
	return canvas.height
}

// Rows returns the number of terminal rows (height / 4).
func (canvas *Canvas) Rows() int {
	return canvas.height / 4
}

// Cols returns the number of terminal columns (width / 2).
func (canvas *Canvas) Cols() int {
	return canvas.width / 2
}

// Set turns on the pixel at the specified coordinates.
func (canvas *Canvas) Set(x, y float64) {
	cellRow, cellColumn, dotRow, dotColumn, ok := canvas.pixelToCell(x, y)
	if !ok {
		return
	}
	canvas.cells[cellRow][cellColumn] |= pixelMap[dotRow][dotColumn]
}

// SetColor sets the pixel at the specified coordinates and assigns the given color
// to the containing cell. Without WithColor(), the pixel is set but color is ignored.
func (canvas *Canvas) SetColor(x, y float64, color Color) {
	cellRow, cellColumn, dotRow, dotColumn, ok := canvas.pixelToCell(x, y)
	if !ok {
		return
	}
	canvas.cells[cellRow][cellColumn] |= pixelMap[dotRow][dotColumn]
	if canvas.colors != nil {
		canvas.colors[cellRow][cellColumn] = color
	}
}

// Unset turns off the pixel at the specified coordinates.
func (canvas *Canvas) Unset(x, y float64) {
	cellRow, cellColumn, dotRow, dotColumn, ok := canvas.pixelToCell(x, y)
	if !ok {
		return
	}
	canvas.cells[cellRow][cellColumn] &^= pixelMap[dotRow][dotColumn]
}

// Toggle inverts the pixel at the specified coordinates.
func (canvas *Canvas) Toggle(x, y float64) {
	cellRow, cellColumn, dotRow, dotColumn, ok := canvas.pixelToCell(x, y)
	if !ok {
		return
	}
	canvas.cells[cellRow][cellColumn] ^= pixelMap[dotRow][dotColumn]
}

// Get returns true if the pixel at the specified coordinates is set.
func (canvas *Canvas) Get(x, y float64) bool {
	cellRow, cellColumn, dotRow, dotColumn, ok := canvas.pixelToCell(x, y)
	if !ok {
		return false
	}
	return canvas.cells[cellRow][cellColumn]&pixelMap[dotRow][dotColumn] != 0
}

// Clear resets all cells to the empty braille pattern.
func (canvas *Canvas) Clear() {
	for row := range canvas.cells {
		for column := range canvas.cells[row] {
			canvas.cells[row][column] = BrailleOffset
		}
	}
	if canvas.colors != nil {
		for row := range canvas.colors {
			for column := range canvas.colors[row] {
				canvas.colors[row][column] = ColorDefault
			}
		}
	}
}

// Frame renders the canvas to a string with rows joined by newlines.
func (canvas *Canvas) Frame() string {
	if !canvas.colorEnabled {
		// Fast path: no colors (existing implementation)
		rows := make([]string, len(canvas.cells))
		for index, row := range canvas.cells {
			rows[index] = string(row)
		}
		return strings.Join(rows, "\n")
	}

	// Color-enabled path
	var builder strings.Builder
	for rowIndex, row := range canvas.cells {
		if rowIndex > 0 {
			builder.WriteByte('\n')
		}
		for columnIndex, cell := range row {
			color := canvas.colors[rowIndex][columnIndex]
			if color != ColorDefault {
				builder.WriteString(color.ANSI())
				builder.WriteRune(cell)
				builder.WriteString(ANSIReset())
			} else {
				builder.WriteRune(cell)
			}
		}
	}
	return builder.String()
}

// pixelToCell converts pixel coordinates to cell and dot positions.
// Returns ok = false for out-of-bounds coordinates.
func (canvas *Canvas) pixelToCell(x, y float64) (cellRow, cellColumn, dotRow, dotColumn int, ok bool) {
	pixelX := int(math.Floor(x))
	pixelY := int(math.Floor(y))

	// Handle Y-axis inversion
	if canvas.invertY {
		pixelY = canvas.height - 1 - pixelY
	}

	// Check bounds against pixel dimensions
	if pixelX < 0 || pixelX >= canvas.width || pixelY < 0 || pixelY >= canvas.height {
		return 0, 0, 0, 0, false
	}

	// Calculate cell position
	cellColumn = pixelX / 2
	cellRow = pixelY / 4

	// Check bounds against cell dimensions (handles dimension truncation)
	if cellRow >= canvas.Rows() || cellColumn >= canvas.Cols() {
		return 0, 0, 0, 0, false
	}

	// Calculate dot position within cell
	dotColumn = pixelX % 2
	dotRow = pixelY % 4

	return cellRow, cellColumn, dotRow, dotColumn, true
}
