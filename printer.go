package main

import (
	"fmt"
	"time"
)

// 4 bars takes 1678ms

func printer(done chan bool) {
	// intro
	intro()
	// verse 1
	verse1()
	// chorus 1
	chorusBig1()
	chorusBig1()
	overAndOver()
	chorusBig1()
	// verse 2
	verse2()
	// chorus 2 - ends faster and starts loading bar
	chorusBig1()
	chorusBig1()
	overAndOver()
	chorusSmall1()
	loading()
	// solo

	// chorus 3 - can't feel, can see...
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
