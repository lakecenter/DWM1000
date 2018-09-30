package main

import (
	"fmt"

	dwm "github.com/hi-manshu/DWM1000"
)

func main() {

	var tag dwm.DWM1000
	tag.Init()
	fmt.Println(tag.DeviceID())

	fmt.Println(tag.EUI_ID())
<<<<<<< HEAD
	fmt.Println(tag.TxPower())

	fmt.Println(tag.PANid())
=======
>>>>>>> ac03ff24d7e5a2bb2c6b563e0eb5a60d8273d46f
}
