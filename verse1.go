package main

import "time"

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
	printString(verse1_1_1, 104)                        // 3/1
	printString(verse1_1_2, 70)                         // 1/1
	printString(verse1_1_3, 88)                         // 4/1
	printString(verse1_1_4, 83)                         // 2/1
	time.Sleep(time.Millisecond*2518 + 11 + 0 + 6 + 9)  // 6/1 + leftovers
	printString(verse1_2_1, 69)                         // 3/1
	printString(verse1_2_2, 38)                         // 1/1
	printString(verse1_2_3, 78)                         // 3/3
	printString(verse1_2_4, 251)                        // 3/3
	time.Sleep(time.Millisecond*1678 + 17 + 2 + 11 + 4) // 4/1 + leftovers
	cleanDisplay()
	time.Sleep(time.Millisecond * 839) // 2/1

	printString(verse1_3_1, 69)                          // 3/1
	printString(verse1_3_2, 60)                          // 1/1
	printString(verse1_3_3, 83)                          // 2/1
	printString(verse1_3_4, 139)                         // 4/1
	time.Sleep(time.Millisecond*2518 + 17 + 0 + 9 + 10)  // 6/1 + leftovers
	printString(verse1_4_1, 69)                          // 2/1
	printString(verse1_4_2, 64)                          // 2/1
	printString(verse1_4_3, 104)                         // 3/1
	printString(verse1_4_4, 89)                          // 3/1
	time.Sleep(time.Millisecond*1678 + 11 + 7 + 11 + 13) // 4/1 + leftovers
	cleanDisplay()
	time.Sleep(time.Millisecond * 839) // 2/1
}
