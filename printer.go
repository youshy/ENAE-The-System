package main

import (
	"fmt"
	"time"
)

var (
	// welcome screen
	intro1 string = "Everybody Needs an Enemy\n"
	intro2 string = "System version: v1.0\n"
	intro3 string = "Build version: v1.0\n"
	intro4 string = "Truth: UNKNOWN\n"

	// verse 1
	verse1_1_1 string = "Take it all " // 12
	verse1_1_2 string = "away \n"      // 25
	verse1_1_3 string = "and watch your wave"
	verse1_1_4 string = " goodbye\n\n" // 9

	verse1_2_1 string = "They'll wipe away "
	verse1_2_2 string = "your smile\n"
	verse1_2_3 string = "before you even "
	verse1_2_4 string = "try\n\n"

	verse1_3_1 string = "Their creepy eyes "
	verse1_3_2 string = "always\n"
	verse1_3_3 string = "watch you "
	verse1_3_4 string = "walking by\n\n"

	verse1_4_1 string = "How can you "
	verse1_4_2 string = "win the race\n"
	verse1_4_3 string = "when you're still "
	verse1_4_4 string = "stuck in the vice\n\n"
)

// 4 bars takes 1678ms

func printer(done chan bool) {
	// intro
	cleanDisplay()
	printString(intro1, 33)
	time.Sleep(time.Millisecond*839 + 14)
	printString(intro2, 39)
	time.Sleep(time.Millisecond*839 + 20)
	printString(intro3, 41)
	time.Sleep(time.Millisecond*839 + 19)
	printString(intro4, 55)
	time.Sleep(time.Millisecond*839 + 14)
	cleanDisplay()
	time.Sleep(time.Millisecond * 3357)
	time.Sleep(time.Millisecond * 3357)

	// first verse
	printString(verse1_1_1, 104)                        // 3/1
	printString(verse1_1_2, 70)                         // 1/1
	printString(verse1_1_3, 88)                         // 4/1
	printString(verse1_1_4, 83)                         // 2/1
	time.Sleep(time.Millisecond*2518 + 11 + 0 + 6 + 9)  // 6/1 + leftovers
	printString(verse1_2_1, 69)                         // 3/1
	printString(verse1_2_2, 38)                         // 1/1
	printString(verse1_2_3, 78)                         // 3/3
	printString(verse1_2_4, 251)                        // 3/3
	time.Sleep(time.Millisecond*1678 + 17 + 2 + 11 + 4) // 4/1 + leftovers
	cleanDisplay()
	time.Sleep(time.Millisecond * 839) // 2/1

	printString(verse1_3_1, 69)                          // 3/1
	printString(verse1_3_2, 60)                          // 1/1
	printString(verse1_3_3, 83)                          // 2/1
	printString(verse1_3_4, 139)                         // 4/1
	time.Sleep(time.Millisecond*2518 + 17 + 0 + 9 + 10)  // 6/1 + leftovers
	printString(verse1_4_1, 69)                          // 2/1
	printString(verse1_4_2, 64)                          // 2/1
	printString(verse1_4_3, 69)                          // 3/1
	printString(verse1_4_4, 89)                          // 3/1
	time.Sleep(time.Millisecond*1678 + 11 + 7 + 17 + 13) // 4/1 + leftovers
	cleanDisplay()
	time.Sleep(time.Millisecond * 839) // 2/1

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
