package main

import "fmt"

func (s *Sized) lastChorus() {
	noteRest(quaternote)
	s.printCenter("Can't feel your pain")
	noteRest(fullnote)
	s.printCenter("But I can see your face")
	noteRest(fullnote + fullnote + halfnote + quaternote)
	cleanDisplay()
	noteRest(quaternote)
	s.printCenter("Can't feel your pain")
	noteRest(fullnote)
	s.printCenter("But I can see your face")
	noteRest(fullnote + fullnote + halfnote + quaternote)
	cleanDisplay()
	noteRest(quaternote)
	s.printCenter("Can't feel your pain")
	noteRest(fullnote)
	s.printCenter("But I can see your face")
	noteRest(fullnote + fullnote + halfnote + quaternote)
	cleanDisplay()
	noteRest(quaternote)
	s.printCenter("Can't feel your pain")
	noteRest(fullnote)
	s.printCenter("But I can see your face")
	noteRest(fullnote + fullnote + halfnote + quaternote)
	cleanDisplay()
}

func (s *Sized) printCenter(str string) {
	fmt.Printf("\x1b[%d;%df", s.height/2, 1)
	fmt.Println(centerText(str, s.width))
}
