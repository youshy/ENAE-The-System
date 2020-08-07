package main

import (
	"fmt"
	"math"
	"time"
)

var (
	transmission string = "Intercepting transmission\n"
	downloading  string = "Downloading:\n"
	message      string = "This solo has been improvised from start to finish. It was take number 2 or so and we liked it so much, that we have banned Artur from redoing it. So, most likely, it will be different each time! For the nerds - that is Artur's heavily modded Fender Jazzmaster USA into Earthquaker Devices Tentacle into GodCityInstruments Baracus into Ampeg V4 into ultra-oversized Zilla 212 with Celestion V30 and Celestion Neo-Creamback."
)

func (s *Sized) solo() {
	printInMicroseconds(transmission, halfnote)
	printInMicroseconds(downloading, quaternote)
	totalTime := halfnote + fullnote
	progressBar(100, totalTime)
	cleanDisplay()
	s.guitarSolo()
}

var (
	R_White float64 = 255
	G_White float64 = 255
	B_White float64 = 255
	R_Green float64 = 0
	G_Green float64 = 204
	B_Green float64 = 0
)

func progressBar(length, totalTime int) {
	timePerChar := totalTime / (length + 2) // we need to wrap [ ]
	fmt.Printf("[")
	// it goes from R:255, G:255, B:255 (White) to R:102, G:255, B:51 (Vivid green)
	var (
		stepR float64 = (R_White - R_Green) / float64(length)
		stepG float64 = (G_White - G_Green) / float64(length)
		stepB float64 = (B_White - B_Green) / float64(length)
		tempR float64 = R_White
		tempG float64 = G_White
		tempB float64 = B_White
	)

	for i := 0; i < length; i++ {
		roundR := int(math.RoundToEven(tempR))
		roundG := int(math.RoundToEven(tempG))
		roundB := int(math.RoundToEven(tempB))
		colourTextRGB(roundR, roundG, roundB)
		fmt.Printf("|")
		time.Sleep(time.Millisecond * time.Duration(timePerChar))
		tempR -= stepR
		tempG -= stepG
		tempB -= stepB
	}
	fmt.Printf("]")
}

func (s *Sized) guitarSolo() {
	colourTextRGB(int(R_Green), int(G_Green), int(B_Green))
	// 1st bar
	s.printNoteInBinary("D", fullnote)
	s.printNoteInBinary("F", fullnote)
	s.printNoteInBinary("G", halfnote)
	s.printNoteInBinary("A", halfnote)
	s.printNoteInBinary("C", fullnote)
	cleanDisplay()
	// 2nd bar
	noteRest(eightnote)
	s.printNoteInBinary("A", sixteenthnote)
	s.printNoteInBinary("C", sixteenthnote)
	s.printNoteInBinary("D", eightnote)
	s.printNoteInBinary("G", quaternote)
	s.printNoteInBinary("E", eightnote)
	s.printNoteInBinary("E", eightnote)
	s.printNoteInBinary("F", quaternote)
	s.printNoteInBinary("F#", eightnote)
	s.printNoteInBinary("F", eightnote)
	s.printNoteInBinary("E", eightnote)
	s.printNoteInBinary("F", quaternote)
	s.printNoteInBinary("E", eightnote)
	s.printNoteInBinary("D", quaternote)
	s.printNoteInBinary("F", eightnote)
	s.printNoteInBinary("E", eightnote)
	s.printNoteInBinary("C", quaternote+eightnote)
	s.printNoteInBinary("E", eightnote)
	s.printNoteInBinary("A", fullnote+eightnote)
	cleanDisplay()
	// 3rd bar
	noteRest(eightnote)
	s.printNoteInBinary("C", sixteenthnote)
	s.printNoteInBinary("D", sixteenthnote)
	s.printNoteInBinary("A", sixteenthnote)
	noteRest(sixteenthnote)
	s.printNoteInBinary("?", eightnote)
	s.printNoteInBinary("?", eightnote)
	s.printNoteInBinary("?", eightnote)
	noteRest(quaternote)

	noteRest(eightnote)
	s.printNoteInBinary("C", sixteenthnote)
	s.printNoteInBinary("E", sixteenthnote)
	noteRest(sixteenthnote)
	noteRest(sixteenthnote)
	s.printNoteInBinary("?", eightnote)
	s.printNoteInBinary("?", eightnote)
	s.printNoteInBinary("?", eightnote)
	noteRest(quaternote)

	noteRest(eightnote)
	s.printNoteInBinary("C", sixteenthnote)
	s.printNoteInBinary("D", sixteenthnote)
	s.printNoteInBinary("A", sixteenthnote)
	noteRest(sixteenthnote)
	s.printNoteInBinary("?", eightnote)
	s.printNoteInBinary("?", eightnote)
	s.printNoteInBinary("?", eightnote)
	noteRest(quaternote)

	noteRest(eightnote)
	s.printNoteInBinary("A", eightnote)
	s.printNoteInBinary("G", eightnote)
	s.printNoteInBinary("E", sixteenthnote)
	s.printNoteInBinary("G", sixteenthnote)
	s.printNoteInBinary("E", sixteenthnote)
	s.printNoteInBinary("D", sixteenthnote)
	s.printNoteInBinary("E", eightnote)
	noteRest(quaternote)
	cleanDisplay()
	// 4th bar
	s.printNoteInBinary("F", sixteenthnote)
	s.printNoteInBinary("D", sixteenthnote)
	s.printNoteInBinary("C", sixteenthnote)
	s.printNoteInBinary("F", sixteenthnote)
	s.printNoteInBinary("C", sixteenthnote)
	s.printNoteInBinary("A", sixteenthnote)
	s.printNoteInBinary("C", sixteenthnote)
	s.printNoteInBinary("A", sixteenthnote)
	s.printNoteInBinary("G#", sixteenthnote)
	s.printNoteInBinary("A", sixteenthnote)
	s.printNoteInBinary("G#", sixteenthnote)
	s.printNoteInBinary("G", sixteenthnote)
	s.printNoteInBinary("G#", sixteenthnote)
	s.printNoteInBinary("G", sixteenthnote)
	s.printNoteInBinary("F", sixteenthnote)
	s.printNoteInBinary("G", sixteenthnote)

	s.printNoteInBinary("F", sixteenthnote)
	s.printNoteInBinary("D", sixteenthnote)
	s.printNoteInBinary("F", sixteenthnote)
	s.printNoteInBinary("D", sixteenthnote)
	s.printNoteInBinary("C", sixteenthnote)
	s.printNoteInBinary("D", sixteenthnote)
	s.printNoteInBinary("C", sixteenthnote)
	s.printNoteInBinary("A", sixteenthnote)
	s.printNoteInBinary("G#", sixteenthnote)
	s.printNoteInBinary("G", sixteenthnote)
	s.printNoteInBinary("F", sixteenthnote)
	s.printNoteInBinary("D", sixteenthnote)
	s.printNoteInBinary("C", sixteenthnote)
	s.printNoteInBinary("D", sixteenthnote)
	s.printNoteInBinary("F", sixteenthnote)
	s.printNoteInBinary("G", sixteenthnote)
	cleanDisplay()
	colourText(blue, false)
	printBinaryInMicroseconds(message, fullnote)
	resetColourChanges()
	noteRest(halfnote)
	noteRest(quaternote)
	cleanDisplay()
}
