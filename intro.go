package main

var (
	// welcome screen
	intro1 string = "Everybody Needs an Enemy\n"
	intro2 string = "Low\n"
	intro3 string = "Build version: v1.0\n"
	intro4 string = "Written, Produced, Mixed, Mastered by Everybody Needs An Enemy\n"
	intro5 string = "Music Video programmed by Artur Kondas\n"
)

func intro() {
	cleanDisplay()
	printInMicroseconds(intro1, halfnote)
	noteRest(halfnote)
	printInMicroseconds(intro2, halfnote)
	noteRest(halfnote)
	printInMicroseconds(intro3, halfnote)
	noteRest(halfnote)
	printInMicroseconds(intro4, halfnote)
	noteRest(halfnote)
	cleanDisplay()
	printInMicroseconds(intro5, halfnote)
	noteRest(halfnote)
	cleanDisplay()
	noteRest(fullnote)
	noteRest(fullnote)
	noteRest(fullnote)
}
