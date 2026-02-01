// Package canvas provides a braille graphics rendering canvas for terminal applications.
package canvas

// BrailleOffset is the Unicode code point for the empty braille pattern (U+2800).
const BrailleOffset rune = 0x2800

// pixelMap maps (row, column) positions within a braille cell to their bit values.
// A braille cell is 2 columns wide and 4 rows tall.
//
// Dot positions:     Bit values:
//
//	0  3               0x01  0x08
//	1  4               0x02  0x10
//	2  5               0x04  0x20
//	6  7               0x40  0x80
//
// Note: The order of rows in this array is meaningful and should not be changed.
var pixelMap = [4][2]rune{
	{0x01, 0x08}, // row 0: dots 0 and 3
	{0x02, 0x10}, // row 1: dots 1 and 4
	{0x04, 0x20}, // row 2: dots 2 and 5
	{0x40, 0x80}, // row 3: dots 6 and 7
}
