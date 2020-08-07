package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var (
	// notes
	fullnote      int = 1678
	halfnote      int = 839
	quaternote    int = 420
	eightnote     int = 210
	sixteenthnote int = 105
)

type Sized struct {
	width  int
	height int
}

func printer(done chan bool) {
	s := &Sized{}
	s.width, s.height = checkTerminalSize()
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
	// solo
	s.solo()
	// chorus 3 - can't feel, can see...
	s.lastChorus()
	<-done
}

func (s *Sized) printNoteInBinary(note string, speed int) {
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

func printInMicroseconds(s string, spd int) {
	speed := spd * 1000
	freq := (float64(speed) / float64(len(s)))
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v", string(s[i]))
		time.Sleep(time.Microsecond * time.Duration(freq))
	}
}

func printBinaryInMicroseconds(s string, spd int) {
	speed := spd * 1000
	var toBinary string
	for _, c := range s {
		toBinary += fmt.Sprintf("%b ", c)
	}
	freq := (float64(speed) / float64(len(toBinary))) // to microseconds
	for _, b := range toBinary {
		fmt.Printf("%c", b)
		time.Sleep(time.Microsecond * time.Duration(freq))
	}
	fmt.Println()
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

// As it goes by divisions
// full note - 1 - 1678
// half note - 2 - 839
// quater note - 4 - 420
// eight note - 8 - 210
// sixteenth note - 16 - 105
// If extended, needs to validate. Good for now.
func noteRest(note int) {
	time.Sleep(time.Millisecond * time.Duration(note))
}

func checkTerminalSize() (int, int) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
		return 0, 0
	}
	valuePairs := strings.Fields(string(out))
	height, err := strconv.Atoi(valuePairs[0])
	if err != nil {
		log.Fatal(err)
		return 0, 0
	}
	width, err := strconv.Atoi(valuePairs[1])
	if err != nil {
		log.Fatal(err)
		return 0, 0
	}
	return width, height
}
