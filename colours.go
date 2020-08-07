package main

import "fmt"

const (
	// basic colours
	black = iota
	red
	green
	yellow
	blue
	magenta
	cyan
	white
)

func resetColourChanges() {
	fmt.Print("\x1b[0m")
}

func colourText(colour int, bright bool) {
	if bright {
		fmt.Printf("\x1b[9%vm", colour)
	} else {
		fmt.Printf("\x1b[3%vm", colour)
	}
}

func colourTextRGB(red, green, blue int) {
	fmt.Printf("\x1b[38;2;%v;%v;%vm", red, green, blue)
}

func whiteBackgroundBlackText() {
	fmt.Print("\x1b[47m\x1b[30m")
}
