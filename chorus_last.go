package main

import (
	"fmt"
	"math"
)

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

func (s *Sized) lastChorus() {
	cleanDisplay()
	maxHeight := (s.height - 32) / 2
	var iter int
	for j := 0; j < s.height; j++ {
		fmt.Println(s.generateLine(":"))
	}
	moveCursor(maxHeight, 0)
	for i := 0; i < len(banners); i++ {
		moveCursor(maxHeight+iter, 0)
		fmt.Println(s.formatLine(banners[i]))
		noteRest(sixteenthnote)
		iter++
	}
	noteRest(eightnote)
}

func (s *Sized) printCenter(str string) {
	fmt.Printf("\x1b[%d;%df", s.height/2, 1)
	fmt.Println(centerText(str, s.width))
}
