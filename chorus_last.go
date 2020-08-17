package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

// There is an extra quaternote from the solo section here!
var banners = map[int]string{
	0:  "::######::::::###:::::##::: ##::####::########:::::########::########::########::##:::::::",
	1:  ":##... ##::::## ##::: ###:: ##: ####:... ##..::::: ##.....:: ##.....:: ##.....:: ##:::::::",
	2:  ":##:::..::::##:. ##:: ####: ##:. ##::::: ##::::::: ##::::::: ##::::::: ##::::::: ##:::::::",
	3:  ":##::::::::##:::. ##: ## ## ##::##:::::: ##::::::: ######::: ######::: ######::: ##:::::::",
	4:  ":##::::::: #########: ##. ####:..::::::: ##::::::: ##...:::: ##...:::: ##...:::: ##:::::::",
	5:  ":##::: ##: ##.... ##: ##:. ###:::::::::: ##::::::: ##::::::: ##::::::: ##::::::: ##:::::::",
	6:  "::######:: ##:::: ##: ##::. ##:::::::::: ##::::::: ##::::::: ########: ########: ########:",
	7:  ":......:::..:::::..::..::::..:::::::::::..::::::::..::::::::........::........::........::",
	8:  "::##::::##:::#######:::##:::::##::########::::::########::::::###:::::####::##::: ##::",
	9:  ":. ##::##:::##.... ##: ##:::: ##: ##.... ##:::: ##.... ##::::## ##:::. ##:: ###:: ##::",
	10: "::. ####::: ##:::: ##: ##:::: ##: ##:::: ##:::: ##:::: ##:::##:. ##::: ##:: ####: ##::",
	11: ":::. ##:::: ##:::: ##: ##:::: ##: ########::::: ########:::##:::. ##:: ##:: ## ## ##::",
	12: ":::: ##:::: ##:::: ##: ##:::: ##: ##.. ##:::::: ##.....::: #########:: ##:: ##. ####::",
	13: ":::: ##:::: ##:::: ##: ##:::: ##: ##::. ##::::: ##:::::::: ##.... ##:: ##:: ##:. ###::",
	14: ":::: ##::::. #######::. #######:: ##:::. ##:::: ##:::::::: ##:::: ##::####: ##::. ##::",
	15: ":::::..::::::.......::::.......:::..:::::..:::::..:::::::::..:::::..::....::..::::..::",
	16: ":::########:::##:::::##::########:::::####::::::######::::::###:::::##::: ##::::::######:::########::########:::",
	17: ":::##.... ##: ##:::: ##:... ##..:::::. ##::::::##... ##::::## ##::: ###:: ##:::::##... ##: ##.....:: ##.....::::",
	18: ":::##:::: ##: ##:::: ##:::: ##:::::::: ##::::: ##:::..::::##:. ##:: ####: ##:::: ##:::..:: ##::::::: ##:::::::::",
	19: ":::########:: ##:::: ##:::: ##:::::::: ##::::: ##::::::::##:::. ##: ## ## ##::::. ######:: ######::: ######:::::",
	20: ":: ##.... ##: ##:::: ##:::: ##:::::::: ##::::: ##::::::: #########: ##. ####:::::..... ##: ##...:::: ##...::::::",
	21: ":::##:::: ##: ##:::: ##:::: ##:::::::: ##::::: ##::: ##: ##.... ##: ##:. ###:::::##::: ##: ##::::::: ##:::::::::",
	22: ":::########::. #######::::: ##::::::::####::::. ######:: ##:::: ##: ##::. ##::::. ######:: ########: ########:::",
	23: "::........::::.......::::::..::::::::....::::::......:::..:::::..::..::::..::::::......:::........::........::::",
	24: "::##::::##:::#######:::##:::::##::########::::::########:::::###::::::######:::########:",
	25: ":. ##::##:::##.... ##: ##:::: ##: ##.... ##:::: ##.....:::::## ##::::##... ##: ##.....::",
	26: "::. ####::: ##:::: ##: ##:::: ##: ##:::: ##:::: ##:::::::::##:. ##:: ##:::..:: ##:::::::",
	27: ":::. ##:::: ##:::: ##: ##:::: ##: ########::::: ######::::##:::. ##: ##::::::: ######:::",
	28: ":::: ##:::: ##:::: ##: ##:::: ##: ##.. ##:::::: ##...:::: #########: ##::::::: ##...::::",
	29: ":::: ##:::: ##:::: ##: ##:::: ##: ##::. ##::::: ##::::::: ##.... ##: ##::: ##: ##:::::::",
	30: ":::: ##::::. #######::. #######:: ##:::. ##:::: ##::::::: ##:::: ##:. ######:: ########:",
	31: "::::..::::::.......::::.......:::..:::::..:::::..::::::::..:::::..:::......:::........::",
}

func (s *Sized) formatLine(str string) string {
	add := int(math.RoundToEven((float64(s.width) - float64(len(str))) / 2))
	var pre, post string
	for i := 0; i < add; i++ {
		pre += ":"
		post += ":"
	}
	f := pre + str + post
	if len(f) != s.width {
		f = f[:len(f)-1]
	}

	return f
}

func (s *Sized) generateLine(char string) string {
	var str string
	for i := 0; i < s.width; i++ {
		str += char
	}
	return str
}

func (s *Sized) generateDisplay(char string) {
	var strArr []string
	for i := 0; i < s.height; i++ {
		str := s.generateLine(char)
		strArr = append(strArr, str)
	}
	temp := strings.Join(strArr[:], "\n")
	printInMicroseconds(temp, eightnote)
}

func (s *Sized) lastChorus() {
	cleanDisplay()
	s.generateDisplay(":")
	maxHeight := (s.height - 32) / 2
	var (
		iter int
	)
	noteRest(eightnote)
	moveCursor(maxHeight, 0)
	for i := 0; i < len(banners); i++ {
		moveCursor(maxHeight+iter, 0)
		fmt.Println(s.formatLine(banners[i]))
		if i == (len(banners)-1)/2 {
			noteRest(eightnote)
			iter++
			continue
		}
		noteRest(thirtysecondnote)
		iter++
	}
	noteRest(fullnote)

	rest := 5
	maxMilliseconds := int(math.RoundToEven(float64(16*fullnote) / float64(rest)))

	source := rand.NewSource(time.Now().UnixNano())
	for j := 0; j < maxMilliseconds; j++ {
		randban := rand.New(source).Intn(len(banners))
		banner := banners[randban]
		switch randban {
		case 8, 9, 10, 11, 12, 13, 14, 15:
			// want to get "your"
			mid := len(banner) / 2
			randpos := rand.New(source).Intn(mid) + mid + 1
			changedBanner := banner[:randpos-1] + ":" + banner[randpos:]
			moveCursor(maxHeight+randban, 0)
			fmt.Println(s.formatLine(changedBanner))
			banners[randban] = changedBanner
		case 24, 25, 26, 27, 28, 29, 30, 31:
			// we're leaving the "face" out, thus middle of the banner
			mid := len(banner) / 2
			randpos := rand.New(source).Intn(mid) + 1
			changedBanner := banner[:randpos-1] + ":" + banner[randpos:]
			moveCursor(maxHeight+randban, 0)
			fmt.Println(s.formatLine(changedBanner))
			banners[randban] = changedBanner
		default:
			// we don't want 0 here
			randpos := rand.New(source).Intn(len(banner)-1) + 1
			changedBanner := banner[:randpos-1] + ":" + banner[randpos:]
			moveCursor(maxHeight+randban, 0)
			fmt.Println(s.formatLine(changedBanner))
			banners[randban] = changedBanner
		}
		noteRest(rest)
	}
}

func (s *Sized) printCenter(str string) {
	fmt.Printf("\x1b[%d;%df", s.height/2, 1)
	fmt.Println(centerText(str, s.width))
}
