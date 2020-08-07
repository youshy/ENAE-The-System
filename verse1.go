package main

var (
	verse1_1_1 string = "Take it all " // 12
	verse1_1_2 string = "away \n"      // 25
	verse1_1_3 string = "and watch your wave"
	verse1_1_4 string = " goodbye\n\n" // 9

	verse1_2_1 string = "They'll wipe away "
	verse1_2_2 string = "your smile\n"
	verse1_2_3 string = "before you even "
	verse1_2_4 string = "try\n\n"

	verse1_3_1 string = "Their creepy eyes "
	verse1_3_2 string = "always\n"
	verse1_3_3 string = "watch you "
	verse1_3_4 string = "walking by\n\n"

	verse1_4_1 string = "How can you "
	verse1_4_2 string = "win the race\n"
	verse1_4_3 string = "when you're "
	verse1_4_4 string = "stuck in the vice\n\n"
)

func verse1() {
	printInMicroseconds(verse1_1_1, halfnote+quaternote) // 3/1
	printInMicroseconds(verse1_1_2, quaternote)          // 1/1
	printInMicroseconds(verse1_1_3, fullnote)            // 4/1
	printInMicroseconds(verse1_1_4, halfnote)            // 2/1
	noteRest(fullnote)                                   // 4/1
	noteRest(halfnote)                                   // 2/1
	printInMicroseconds(verse1_2_1, halfnote+quaternote) // 3/1
	printInMicroseconds(verse1_2_2, quaternote)          // 1/1
	printInMicroseconds(verse1_2_3, halfnote+quaternote) // 3/3
	printInMicroseconds(verse1_2_4, halfnote+quaternote) // 3/3
	noteRest(fullnote)                                   // 4/1
	cleanDisplay()
	noteRest(halfnote)                                   // 2/1
	printInMicroseconds(verse1_3_1, halfnote+quaternote) // 3/1
	printInMicroseconds(verse1_3_2, quaternote)          // 1/1
	printInMicroseconds(verse1_3_3, halfnote)            // 2/1
	printInMicroseconds(verse1_3_4, fullnote)            // 4/1
	noteRest(fullnote)                                   // 4/1
	noteRest(halfnote)                                   // 2/1
	printInMicroseconds(verse1_4_1, halfnote)            // 2/1
	printInMicroseconds(verse1_4_2, halfnote)            // 2/1
	printInMicroseconds(verse1_4_3, halfnote+quaternote) // 3/1
	printInMicroseconds(verse1_4_4, halfnote+quaternote) // 3/1
	noteRest(fullnote)                                   // 4/1
	cleanDisplay()
	noteRest(halfnote) // 2/1
}
