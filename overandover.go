package main

import (
	"fmt"
)

var (
	over string = `
   ____ _    ____________ 
  / __ \ |  / / ____/ __ \
 / / / / | / / __/ / /_/ /
/ /_/ /| |/ / /___/ _, _/ 
\____/ |___/_____/_/ |_|  
`
	overand string = `
   ____ _    ____________     ___    _   ______ 
  / __ \ |  / / ____/ __ \   /   |  / | / / __ \
 / / / / | / / __/ / /_/ /  / /| | /  |/ / / / /
/ /_/ /| |/ / /___/ _, _/  / ___ |/ /|  / /_/ / 
\____/ |___/_____/_/ |_|  /_/  |_/_/ |_/_____/  
`

	overandover string = `
   ____ _    ____________     ___    _   ______     ____ _    ____________ 
  / __ \ |  / / ____/ __ \   /   |  / | / / __ \   / __ \ |  / / ____/ __ \
 / / / / | / / __/ / /_/ /  / /| | /  |/ / / / /  / / / / | / / __/ / /_/ /
/ /_/ /| |/ / /___/ _, _/  / ___ |/ /|  / /_/ /  / /_/ /| |/ / /___/ _, _/ 
\____/ |___/_____/_/ |_|  /_/  |_/_/ |_/_____/   \____/ |___/_____/_/ |_|  
`

	theyre string = `
  _______ _    _ ________     ___ _____  ______ 
 |__   __| |  | |  ____\ \   / ( )  __ \|  ____|
    | |  | |__| | |__   \ \_/ /|/| |__) | |__   
    | |  |  __  |  __|   \   /   |  _  /|  __|  
    | |  | |  | | |____   | |    | | \ \| |____ 
    |_|  |_|  |_|______|  |_|    |_|  \_\______|
`

	sending string = `
   _____ ______ _   _ _____ _____ _   _  _____ 
  / ____|  ____| \ | |  __ \_   _| \ | |/ ____|
 | (___ | |__  |  \| | |  | || | |  \| | |  __ 
  \___ \|  __| | . ' | |  | || | | . ' | | |_ |
  ____) | |____| |\  | |__| || |_| |\  | |__| |
 |_____/|______|_| \_|_____/_____|_| \_|\_____|
`

	you string = `
 __     ______  _    _ 
 \ \   / / __ \| |  | |
  \ \_/ / |  | | |  | |
   \   /| |  | | |  | |
    | | | |__| | |__| |
    |_|  \____/ \____/ 
`

	under string = `
  _    _ _   _ _____  ______ _____  
 | |  | | \ | |  __ \|  ____|  __ \ 
 | |  | |  \| | |  | | |__  | |__) |
 | |  | | . ' | |  | |  __| |  _  / 
 | |__| | |\  | |__| | |____| | \ \ 
  \____/|_| \_|_____/|______|_|  \_\
`

	pushing string = `
  ___ _   _ ___ _  _ ___ _  _  ___ 
 | _ \ | | / __| || |_ _| \| |/ __|
 |  _/ |_| \__ \ __ || || .' | (_ |
 |_|  \___/|___/_||_|___|_|\_|\___|
`

	and string = `
    _   _  _ ___  
   /_\ | \| |   \ 
  / _ \| .' | |) |
 /_/ \_\_|\_|___/ 
`

	pulling string = `
    ___    __  __  __     __     ____  _  __  _____
   / _ \  / / / / / /    / /    /  _/ / |/ / / ___/
  / ___/ / /_/ / / /__  / /__  _/ /  /    / / (_ / 
 /_/     \____/ /____/ /____/ /___/ /_/|_/  \___/  
`

	us string = `
          _____                    _____          
         /\    \                  /\    \         
        /::\____\                /::\    \        
       /:::/    /               /::::\    \       
      /:::/    /               /::::::\    \      
     /:::/    /               /:::/\:::\    \     
    /:::/    /               /:::/__\:::\    \    
   /:::/    /                \:::\   \:::\    \   
  /:::/    /      _____    ___\:::\   \:::\    \  
 /:::/____/      /\    \  /\   \:::\   \:::\    \ 
|:::|    /      /::\____\/::\   \:::\   \:::\____\
|:::|____\     /:::/    /\:::\   \:::\   \::/    /
 \:::\    \   /:::/    /  \:::\   \:::\   \/____/ 
  \:::\    \ /:::/    /    \:::\   \:::\    \     
   \:::\    /:::/    /      \:::\   \:::\____\    
    \:::\__/:::/    /        \:::\  /:::/    /    
     \::::::::/    /          \:::\/:::/    /     
      \::::::/    /            \::::::/    /      
       \::::/    /              \::::/    /       
        \::/____/                \::/    /        
         ~~                       \/____/         
`

	away  string = "\n\tAWAY "
	from  string = "\tFROM "
	each  string = "\tEACH "
	other string = "\tOTHER"
)

func overAndOver() {
	fmt.Print(over)
	noteRest(quaternote)
	cleanDisplay()
	fmt.Print(overand)
	noteRest(eightnote)
	cleanDisplay()
	fmt.Print(overandover)
	noteRest(quaternote + eightnote)
	cleanDisplay()
	fmt.Print(theyre)
	noteRest(eightnote)
	fmt.Print(sending)
	noteRest(quaternote)
	fmt.Print(you)
	noteRest(eightnote)
	fmt.Print(under)
	noteRest(halfnote + eightnote)
	cleanDisplay()
	fmt.Print(pushing)
	noteRest(quaternote)
	fmt.Print(and)
	noteRest(eightnote)
	fmt.Print(pulling)
	noteRest(quaternote)
	cleanDisplay()
	fmt.Print(us)
	noteRest(quaternote)
	fmt.Print(away)
	noteRest(quaternote)
	fmt.Print(from)
	noteRest(eightnote)
	fmt.Print(each)
	noteRest(eightnote)
	fmt.Print(other)
	noteRest(halfnote + quaternote)
	cleanDisplay()
}
