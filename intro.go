package main

import "time"

var (
	// welcome screen
	intro1 string = "Everybody Needs an Enemy\n"
	intro2 string = "System version: v1.0\n"
	intro3 string = "Build version: v1.0\n"
	intro4 string = "Truth: UNKNOWN\n"
)

func intro() {
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
}
