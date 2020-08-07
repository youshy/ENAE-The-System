package main

import "fmt"

var (
	pain string = "Can't feel your pain"
	face string = "But I can see your face"
)

func (s *Sized) lastChorus() {
	for i := 0; i < 4; i++ {
		noteRest(quaternote)
		s.printCenter(pain)
		noteRest(fullnote)
		s.printCenter(face)
		noteRest(fullnote + fullnote + halfnote + quaternote)
		cleanDisplay()
	}
}

func (s *Sized) printCenter(str string) {
	fmt.Printf("\x1b[%d;%df", s.height/2, 1)
	fmt.Println(centerText(str, s.width))
}
