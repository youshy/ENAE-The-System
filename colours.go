package main

import "fmt"

const (
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

func whiteBackgroundBlackText() {
	fmt.Print("\x1b[47m\x1b[30m")
}
