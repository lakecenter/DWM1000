package dwm1000

import (
	"golang.org/x/exp/io/spi"
	"log"
	"fmt"
)

type DWM1000 struct{
	dev *spi.Device

}

func (d *DWM1000) Init() error {
	var err error
	d.dev, err = spi.Open(&spi.Devfs{
		Dev:      "/dev/spidev0.0",
		Mode:     spi.Mode0,
		MaxSpeed: 20000000,
	})

	dev.SetBitOrder(spi.MSBFirst)
	if err==nil{
		log.Printf("Device Initated : *Success*")
	}
	return err

}

func DeviceID() {

	cmd := []byte{0x00, 0x00, 0x00, 0x00}
	output := make([]byte, len(cmd))
	error := dev.Tx(cmd, output)
	if error != nil {
		log.Println("Error reading REG ")
	}
	fmt.Printf("\n SYS0 Response : %0x \n", output[1:])
}

func flip(data []byte) []byte{
	N:=len(data)
	result:=make([]byte,len(data))
	for d,i:=range data {
	result[N-i-1]=d
	}
return result
}