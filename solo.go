package main

import (
	"fmt"
	"time"
)

var (
	transmission string = "Intercepting transmission\n"
	downloading  string = "Downloading:\n"
	message      string = "This solo has been improvised from start to finish. It was take number 2 or so and we liked it so much, that we have banned Artur from redoing it. So, most likely, it will be different each time! For the nerds - that is Artur's heavily modded Fender Jazzmaster USA into Earthquaker Devices Tentacle into GodCityInstruments Baracus into Ampeg V4 into ultra-oversized Zilla 212 with V30 and Neodimium Creamback."
)

func loading() {
	printString(transmission, 52)
	printString(downloading, 52)
	totalTime := 2518 - 52 - 52
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
	noteRest(8)
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
	// 3rd bar
	noteRest(8)
	printNoteInBinary("C", 105)
	printNoteInBinary("D", 105)
	printNoteInBinary("A", 105)
	noteRest(16)
	printNoteInBinary("?", 210)
	printNoteInBinary("?", 210)
	printNoteInBinary("?", 210)
	noteRest(4)

	noteRest(8)
	printNoteInBinary("C", 105)
	printNoteInBinary("E", 105)
	noteRest(16)
	noteRest(16)
	printNoteInBinary("?", 210)
	printNoteInBinary("?", 210)
	printNoteInBinary("?", 210)
	noteRest(4)

	noteRest(8)
	printNoteInBinary("C", 105)
	printNoteInBinary("D", 105)
	printNoteInBinary("A", 105)
	noteRest(16)
	printNoteInBinary("?", 210)
	printNoteInBinary("?", 210)
	printNoteInBinary("?", 210)
	noteRest(4)

	noteRest(8)
	printNoteInBinary("A", 210)
	printNoteInBinary("G", 210)
	printNoteInBinary("E", 105)
	printNoteInBinary("G", 105)
	printNoteInBinary("E", 105)
	printNoteInBinary("D", 105)
	printNoteInBinary("E", 210)
	noteRest(4)
	cleanDisplay()
	// 4th bar
	printNoteInBinary("F", 105)
	printNoteInBinary("D", 105)
	printNoteInBinary("C", 105)
	printNoteInBinary("F", 105)
	printNoteInBinary("C", 105)
	printNoteInBinary("A", 105)
	printNoteInBinary("C", 105)
	printNoteInBinary("A", 105)
	printNoteInBinary("G#", 105)
	printNoteInBinary("A", 105)
	printNoteInBinary("G#", 105)
	printNoteInBinary("G", 105)
	printNoteInBinary("G#", 105)
	printNoteInBinary("G", 105)
	printNoteInBinary("F", 105)
	printNoteInBinary("G", 105)

	printNoteInBinary("F", 105)
	printNoteInBinary("D", 105)
	printNoteInBinary("F", 105)
	printNoteInBinary("D", 105)
	printNoteInBinary("C", 105)
	printNoteInBinary("D", 105)
	printNoteInBinary("C", 105)
	printNoteInBinary("A", 105)
	printNoteInBinary("G#", 105)
	printNoteInBinary("G", 105)
	printNoteInBinary("F", 105)
	printNoteInBinary("D", 105)
	printNoteInBinary("C", 105)
	printNoteInBinary("D", 105)
	printNoteInBinary("F", 105)
	printNoteInBinary("G", 105)
	cleanDisplay()
	// TOOD: figure out how to print uneven message
	printMessage(message, 1678) // represent as milliseconds?
	noteRest(1)
	cleanDisplay()
	fmt.Println("last chorus")
}

// As it goes by divisions
// full note - 1 - 1678
// half note - 2 - 839
// quater note - 4 - 420
// eight note - 8 - 210
// sixteenth note - 16 - 105
// If extended, needs to validate. Good for now.
func noteRest(note int) {
	switch note {
	case 1:
		time.Sleep(time.Millisecond * 1678)
	case 2:
		time.Sleep(time.Millisecond * 839)
	case 4:
		time.Sleep(time.Millisecond * 420)
	case 8:
		time.Sleep(time.Millisecond * 210)
	case 16:
		time.Sleep(time.Millisecond * 105)
	default:
		panic("Unknown rest note value")
	}

}

func printNoteInBinary(note string, speed int) {
	var toBinary string
	for _, c := range note {
		toBinary += fmt.Sprintf("%b ", c)
	}
	freq := (float64(speed) / float64(len(toBinary))) // to microseconds
	for _, b := range toBinary {
		fmt.Printf("%c", b)
		time.Sleep(time.Millisecond * time.Duration(freq))
	}
	fmt.Println()
}
