package main

import (
	"fmt"
	"time"
)

var (
	transmission string = "Intercepting transmission\n"
	downloading  string = "Downloading:\n"
)

func loading() {
	printString(transmission, 52)
	printString(downloading, 52)
	totalTime := 2518 - 52 - 52 //
	progressBar(100, totalTime)
	cleanDisplay()
	solo()
}

func progressBar(length, totalTime int) {
	timePerChar := totalTime / (length + 2) // we need to wrap [ ]
	fmt.Printf("[")
	for i := 0; i < length; i++ {
		fmt.Printf("|")
		time.Sleep(time.Millisecond * time.Duration(timePerChar))
	}
	fmt.Printf("]")
}

func solo() {
	// 1st bar
	printNoteInBinary("D", 1678)
	printNoteInBinary("F", 1678)
	printNoteInBinary("G", 839)
	printNoteInBinary("A", 839)
	printNoteInBinary("C", 1678)
	cleanDisplay()
	// 2nd bar
	time.Sleep(time.Millisecond * 210) // sixteenth note rest
	printNoteInBinary("A", 105)
	printNoteInBinary("C", 105)
	printNoteInBinary("D", 210)
	printNoteInBinary("G", 420)
	printNoteInBinary("E", 210)
	printNoteInBinary("E", 210)
	printNoteInBinary("F", 420)
	printNoteInBinary("F#", 210)
	printNoteInBinary("F", 210)
	printNoteInBinary("E", 210)
	printNoteInBinary("F", 420)
	printNoteInBinary("E", 210)
	printNoteInBinary("D", 420)
	printNoteInBinary("F", 210)
	printNoteInBinary("E", 210)
	printNoteInBinary("C", 630)
	printNoteInBinary("E", 210)
	printNoteInBinary("A", 1890)
	cleanDisplay()
}

func printNoteInBinary(note string, speed int) {
	var toBinary string
	for _, c := range note {
		toBinary += fmt.Sprintf("%b ", c)
	}
	freq := speed / len(toBinary)
	for _, b := range toBinary {
		fmt.Printf("%c", b)
		time.Sleep(time.Millisecond * time.Duration(freq))
	}
	fmt.Println()
}
