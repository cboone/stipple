package canvas

import "testing"

func TestBrailleOffset(t *testing.T) {
	if BrailleOffset != 0x2800 {
		t.Errorf("BrailleOffset = %#x, want %#x", BrailleOffset, 0x2800)
	}
}

func TestPixelMapDotPositions(t *testing.T) {
	tests := []struct {
		name     string
		row      int
		column   int
		expected rune
	}{
		{"dot 0 (row 0, col 0)", 0, 0, 0x01},
		{"dot 3 (row 0, col 1)", 0, 1, 0x08},
		{"dot 1 (row 1, col 0)", 1, 0, 0x02},
		{"dot 4 (row 1, col 1)", 1, 1, 0x10},
		{"dot 2 (row 2, col 0)", 2, 0, 0x04},
		{"dot 5 (row 2, col 1)", 2, 1, 0x20},
		{"dot 6 (row 3, col 0)", 3, 0, 0x40},
		{"dot 7 (row 3, col 1)", 3, 1, 0x80},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			got := pixelMap[testCase.row][testCase.column]
			if got != testCase.expected {
				t.Errorf("pixelMap[%d][%d] = %#x, want %#x",
					testCase.row, testCase.column, got, testCase.expected)
			}
		})
	}
}

func TestPixelMapFullCell(t *testing.T) {
	var fullCell rune
	for row := 0; row < 4; row++ {
		for column := 0; column < 2; column++ {
			fullCell |= pixelMap[row][column]
		}
	}

	expected := rune(0xFF)
	if fullCell != expected {
		t.Errorf("full cell bits = %#x, want %#x", fullCell, expected)
	}

	fullBraille := BrailleOffset + fullCell
	expectedBraille := rune(0x28FF)
	if fullBraille != expectedBraille {
		t.Errorf("full braille = %#x, want %#x", fullBraille, expectedBraille)
	}
}

func TestPixelMapEmptyCell(t *testing.T) {
	emptyBraille := BrailleOffset
	expected := rune(0x2800)
	if emptyBraille != expected {
		t.Errorf("empty braille = %#x, want %#x", emptyBraille, expected)
	}
}
