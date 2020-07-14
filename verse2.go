package main

import (
	"fmt"
	"time"
)

var (
	frame1 string = `
















` // 16 lines

	frame2 string = `















          Never gave
` // 16 lines

	frame3 string = `














          Never gave
          A chance
` // 16 lines

	frame4 string = `













          Never gave
          A chance
					'Cause they had
` // 16 lines

	frame5 string = `












          Never gave
          A chance
					'Cause they had
					Other plans
` // 16 lines

	frame6 string = `











          Never gave
          A chance
					'Cause they had
					Other plans

` // 16 lines

	frame7 string = `










          Never gave
          A chance
					'Cause they had
					Other plans


` // 16 lines

	frame8 string = `









          Never gave
          A chance
					'Cause they had
					Other plans



` // 16 lines

	frame9 string = `








          Never gave
          A chance
					'Cause they had
					Other plans




` // 16 lines

	frame10 string = `







          Never gave
          A chance
					'Cause they had
					Other plans




          When you're
` // 16 lines

	frame11 string = `






          Never gave
          A chance
					'Cause they had
					Other plans




          When you're
					Kept low
` // 16 lines

	frame12 string = `




          Never gave
          A chance
					'Cause they had
					Other plans




          When you're
					Kept low
					You can't show up
` // 16 lines

	frame13 string = `




          Never gave
          A chance
					'Cause they had
					Other plans




          When you're
					Kept low
					You can't show up
					And make a stand
` // 16 lines

	frame14 string = `



          Never gave
          A chance
					'Cause they had
					Other plans




          When you're
					Kept low
					You can't show up
					And make a stand

` // 16 lines

	frame15 string = `


          Never gave
          A chance
					'Cause they had
					Other plans




          When you're
					Kept low
					You can't show up
					And make a stand


` // 16 lines

	frame16 string = `

          Never gave
          A chance
					'Cause they had
					Other plans




          When you're
					Kept low
					You can't show up
					And make a stand



` // 16 lines

	frame17 string = `
          Never gave
          A chance
					'Cause they had
					Other plans




          When you're
					Kept low
					You can't show up
					And make a stand




` // 16 lines

	frame18 string = `
          A chance
					'Cause they had
					Other plans




          When you're
					Kept low
					You can't show up
					And make a stand




          Who's gonna help
` // 16 lines

	frame19 string = `
					'Cause they had
					Other plans




          When you're
					Kept low
					You can't show up
					And make a stand




          Who's gonna help
					The ones
` // 16 lines

	frame20 string = `
					Other plans




          When you're
					Kept low
					You can't show up
					And make a stand




          Who's gonna help
					The ones
					That's hurting
` // 16 lines

	frame21 string = `




          When you're
					Kept low
					You can't show up
					And make a stand




          Who's gonna help
					The ones
					That's hurting
					By your side
` // 16 lines

	frame22 string = `



          When you're
					Kept low
					You can't show up
					And make a stand




          Who's gonna help
					The ones
					That's hurting
					By your side

` // 16 lines

	frame23 string = `


          When you're
					Kept low
					You can't show up
					And make a stand




          Who's gonna help
					The ones
					That's hurting
					By your side


` // 16 lines

	frame24 string = `

          When you're
					Kept low
					You can't show up
					And make a stand




          Who's gonna help
					The ones
					That's hurting
					By your side



` // 16 lines

	frame25 string = `
          When you're
					Kept low
					You can't show up
					And make a stand




          Who's gonna help
					The ones
					That's hurting
					By your side




` // 16 lines

	frame26 string = `
					Kept low
					You can't show up
					And make a stand




          Who's gonna help
					The ones
					That's hurting
					By your side




          Will never
` // 16 lines

	frame27 string = `
					You can't show up
					And make a stand




          Who's gonna help
					The ones
					That's hurting
					By your side




          Will never
					Go away
` // 16 lines

	frame28 string = `
					And make a stand




          Who's gonna help
					The ones
					That's hurting
					By your side




          Will never
					Go away
					If no one
` // 16 lines

	frame29 string = `




          Who's gonna help
					The ones
					That's hurting
					By your side




          Will never
					Go away
					If no one
					Tries to fight
` // 16 lines

	frame30 string = `



          Who's gonna help
					The ones
					That's hurting
					By your side




          Will never
					Go away
					If no one
					Tries to fight

` // 16 lines

	frame31 string = `


          Who's gonna help
					The ones
					That's hurting
					By your side




          Will never
					Go away
					If no one
					Tries to fight


` // 16 lines

	frame32 string = `

          Who's gonna help
					The ones
					That's hurting
					By your side




          Will never
					Go away
					If no one
					Tries to fight



` // 16 lines

	frame33 string = `
          Who's gonna help
					The ones
					That's hurting
					By your side




          Will never
					Go away
					If no one
					Tries to fight




` // 16 lines

	frame34 string = `
					The ones
					That's hurting
					By your side




          Will never
					Go away
					If no one
					Tries to fight





` // 16 lines

	frame35 string = `
					That's hurting
					By your side




          Will never
					Go away
					If no one
					Tries to fight






` // 16 lines

	frame36 string = `
					By your side




          Will never
					Go away
					If no one
					Tries to fight







` // 16 lines

	frame37 string = `




          Will never
					Go away
					If no one

					Tries to fight







` // 16 lines

	frame38 string = `



          Will never
					Go away
					If no one


					Tries to fight







` // 16 lines

	frame39 string = `


          Will never
					Go away
					If no one



					Tries to fight







` // 16 lines

	frame40 string = `

          Will never
					Go away
					If no one




					Tries to fight







` // 16 lines

	frame41 string = `
          Will never
					Go away
					If no one





					Tries to fight







` // 16 lines

	frame42 string = `
					Go away
					If no one






					Tries to fight







` // 16 lines

	frame43 string = `
					If no one







					Tries to fight







` // 16 lines

	frame44 string = `








					Tries to fight







` // 16 lines

	frame45 string = `








					         fight







` // 16 lines

	frame46 string = `








					         FIGHT







` // 16 lines

	frame47 string = `






           ___ ___ ___ _  _ _____ 
          | __|_ _/ __| || |_   _|
          | _| | | (_ | __ | | |  
          |_| |___\___|_||_| |_|  






` // 16 lines

	frame48 string = `





        ______ _____ _____ _    _ _______ 
       |  ____|_   _/ ____| |  | |__   __|
       | |__    | || |  __| |__| |  | |   
       |  __|   | || | |_ |  __  |  | |   
       | |     _| || |__| | |  | |  | |   
       |_|    |_____\_____|_|  |_|  |_|   





` // 16 lines

	frame49 string = `





        ###### ##### ##### #    # ####### 
       #  ######   ## ##### #  # ###  ####
       # ###    # ## #  ### #### #  # #   
       #  ###   # ## # ## #  ##  #  # #   
       # #     ## ## #### # #  # #  # #   
       ###    ###############  ###  ###   





` // 16 lines

	frame50 string = `





        ****** ***** ***** *    * ******* 
       *  ******   ** ***** *  * ***  ****
       * ***    * ** *  *** **** *  * *   
       *  ***   * ** * ** *  **  *  * *   
       * *     ** ** **** * *  * *  * *   
       ***    ***************  ***  ***   





` // 16 lines

	frame51 string = `





        ...... ..... ..... .    . ....... 
       .  ......   .. ..... .  . ...  ....
       . ...    . .. .  ... .... .  . .   
       .  ...   . .. . .. .  ..  .  . .   
       . .     .. .. .... . .  . .  . .   
       ...    ...............  ...  ...   





` // 16 lines
)

func verse2() {
	cleanDisplay()
	printFrame(frame1, 410)
	printFrame(frame2, 839)
	printFrame(frame3, 839)
	printFrame(frame4, 839)
	printFrame(frame5, 839)
	printFrame(frame6, 839)
	printFrame(frame7, 839)
	printFrame(frame8, 839)
	printFrame(frame9, 839)
	printFrame(frame10, 839)
	printFrame(frame11, 839)
	printFrame(frame12, 839)
	printFrame(frame13, 839)
	printFrame(frame14, 839)
	printFrame(frame15, 839)
	printFrame(frame16, 839)
	printFrame(frame17, 839)
	printFrame(frame18, 839)
	printFrame(frame19, 839)
	printFrame(frame20, 839)
	printFrame(frame21, 839)
	printFrame(frame22, 839)
	printFrame(frame23, 839)
	printFrame(frame24, 839)
	printFrame(frame25, 839)
	printFrame(frame26, 839)
	printFrame(frame27, 839)
	printFrame(frame28, 839)
	printFrame(frame29, 210)
	printFrame(frame30, 210)
	printFrame(frame31, 210)
	printFrame(frame32, 210)
	printFrame(frame33, 210)
	printFrame(frame34, 210)
	printFrame(frame35, 210)
	printFrame(frame36, 210)
	printFrame(frame37, 210)
	printFrame(frame38, 210)
	printFrame(frame39, 105)
	printFrame(frame40, 105)
	printFrame(frame41, 105)
	printFrame(frame42, 105)
	printFrame(frame43, 105)
	printFrame(frame44, 105)
	printFrame(frame45, 105)
	printFrame(frame46, 105)
	printFrame(frame47, 105)
	printFrame(frame48, 210)
	printFrame(frame49, 210)
	printFrame(frame50, 210)
	printFrame(frame51, 210)
}

func printFrame(frame string, wait int) {
	fmt.Printf("%v", frame)
	time.Sleep(time.Millisecond * time.Duration(wait))
	cleanDisplay()
}
