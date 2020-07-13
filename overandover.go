package main

import (
	"fmt"
	"time"
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
	fmt.Printf("%v", over)
	time.Sleep(time.Millisecond * 420)
	cleanDisplay()
	fmt.Printf("%v", overand)
	time.Sleep(time.Millisecond * 210)
	cleanDisplay()
	fmt.Printf("%v", overandover)
	time.Sleep(time.Millisecond * 629)
	cleanDisplay()
	fmt.Printf("%v", theyre)
	time.Sleep(time.Millisecond * 210)
	fmt.Printf("%v", sending)
	time.Sleep(time.Millisecond * 420)
	fmt.Printf("%v", you)
	time.Sleep(time.Millisecond * 210)
	fmt.Printf("%v", under)
	time.Sleep(time.Millisecond * 210)
	time.Sleep(time.Millisecond * 839)
	cleanDisplay()
	fmt.Printf("%v", pushing)
	time.Sleep(time.Millisecond * 420)
	fmt.Printf("%v", and)
	time.Sleep(time.Millisecond * 210)
	fmt.Printf("%v", pulling)
	time.Sleep(time.Millisecond * 420)
	cleanDisplay()
	fmt.Printf("%v", us)
	time.Sleep(time.Millisecond * 420)
	fmt.Printf("%v", away)
	time.Sleep(time.Millisecond * 420)
	fmt.Printf("%v", from)
	time.Sleep(time.Millisecond * 210)
	fmt.Printf("%v", each)
	time.Sleep(time.Millisecond * 210)
	fmt.Printf("%v", other)
	time.Sleep(time.Millisecond * 420)
	time.Sleep(time.Millisecond * 839)
	cleanDisplay()
}

func overAndOverCenter() {

}
