package main

var (
	// welcome screen
	intro1 string = "Everybody Needs an Enemy\n"
	intro2 string = "Low\n"
	intro3 string = "Build version: v1.0\n"
	intro4 string = "Written, Produced, Mixed, Mastered by Everybody Needs An Enemy\n"
)

func intro() {
	cleanDisplay()
	printInMicroseconds(intro1, halfnote)
	noteRest(2)
	printInMicroseconds(intro2, halfnote)
	noteRest(2)
	printInMicroseconds(intro3, halfnote)
	noteRest(2)
	printInMicroseconds(intro4, halfnote)
	noteRest(2)
	cleanDisplay()
	noteRest(1)
	noteRest(1)
	noteRest(1)
	noteRest(1)
}
