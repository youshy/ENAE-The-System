package main

import (
	"fmt"
	"time"
)

// 4 bars takes 1678ms

func printer(done chan bool) {
	intro()
	verse1()
	chorusBig()
	chorusBig()
	overAndOver()
	chorusBig()
	<-done
}

func printString(str string, speed int) {
	// fmt.Printf("\n\nstring: %v length: %v\n", str, len(str))
	for i := 0; i < len(str); i++ {
		fmt.Printf("%v", string(str[i]))
		time.Sleep(time.Millisecond * time.Duration(speed))
	}
}

func cleanDisplay() {
	fmt.Printf("\x1b[2J")
	moveCursor(0, 0)
}

func moveCursor(row, col int) {
	fmt.Printf("\x1b[%d;%df", row+1, col+1)
}

func centerText(s string, w int) string {
	return fmt.Sprintf("%[1]*s", -w, fmt.Sprintf("%[1]*s", (w+len(s))/2, s))
}
