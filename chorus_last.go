package main

import "fmt"

func (s *Sized) lastChorus() {
	noteRest(4)
	s.printCenter("Can't feel your pain")
	noteRest(1)
	s.printCenter("But I can see your face")
	noteRest(1)
	noteRest(1)
	noteRest(2)
	noteRest(4)
	cleanDisplay()
	noteRest(4)
	s.printCenter("Can't feel your pain")
	noteRest(1)
	s.printCenter("But I can see your face")
	noteRest(1)
	noteRest(1)
	noteRest(2)
	noteRest(4)
	cleanDisplay()
	noteRest(4)
	s.printCenter("Can't feel your pain")
	noteRest(1)
	s.printCenter("But I can see your face")
	noteRest(1)
	noteRest(1)
	noteRest(2)
	noteRest(4)
	cleanDisplay()
	noteRest(4)
	s.printCenter("Can't feel your pain")
	noteRest(1)
	s.printCenter("But I can see your face")
	noteRest(1)
	noteRest(1)
	noteRest(2)
	noteRest(4)
	cleanDisplay()
}

func whiteBackgroundBlackText() {
	fmt.Print("\x1b[47m\x1b[30m")
}

func (s *Sized) printCenter(str string) {
	fmt.Printf("\x1b[%d;%df", s.height/2, 1)
	fmt.Println(centerText(str, s.width))
}
