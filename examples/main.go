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
	fmt.Println(tag.TxPower())

	fmt.Println(tag.PANid())
}
