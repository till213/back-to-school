package main

import "fmt"

func drawHorizontalLine(screen []byte, width int, x1, x2, y int) {
	startOffset := x1 % 8
	firstFullByte := x1 / 8
	if startOffset != 0 {
		firstFullByte++
	}

	endOffset := x2 % 8
	lastFullByte := x2 / 8
	if endOffset != 7 {
		lastFullByte--
	}

	// set the full bytes
	for b := firstFullByte; b <= lastFullByte; b++ {
		screen[(width/8)*y+b] = 0xFF
	}

	startMask := 0xFF >> startOffset
	endMask := ^(0xFF >> (endOffset + 1))

	// Check if in the same byte
	if (x1 / 8) == (x2 / 8) {
		mask := startMask & endMask
		screen[(width/8)*y+(x1/8)] |= byte(mask)
	} else {
		if startOffset != 0 {
			byteNumber := (width/8)*y + firstFullByte - 1
			screen[byteNumber] |= byte(startMask)
		}
		if endOffset != 7 {
			byteNumber := (width/8)*y + lastFullByte + 1
			screen[byteNumber] |= byte(endMask)
		}
	}
}

// in pixel
const width = 32
const height = 16

func drawScreen(screen []byte) {
	height := len(screen) / (width / 8)
	for y := 0; y < height; y++ {
		for x := 0; x < (width / 8); x++ {
			fmt.Printf("%08b", screen[y*(width/8)+x])
		}
		fmt.Println()
	}
}

func main() {
	screen := make([]byte, width/8*height)
	drawHorizontalLine(screen, width, 0, 0, 0)
	drawHorizontalLine(screen, width, 31, 31, 0)
	drawHorizontalLine(screen, width, 0, 0, 15)
	drawHorizontalLine(screen, width, 31, 31, 15)
	for y := 1; y < height-1; y++ {
		drawHorizontalLine(screen, width, 1, 1+y, y)
	}
	drawScreen(screen)
}
