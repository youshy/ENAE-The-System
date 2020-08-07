package main

import "fmt"

func whiteBackgroundBlackText() {
	fmt.Print("\x1b[47m\x1b[30m")
}
