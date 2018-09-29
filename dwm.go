package dwm

import (
	"fmt"
	"log"

	"golang.org/x/exp/io/spi"
)

type DWM1000 struct {
	dev *spi.Device
}

func (d *DWM1000) Init() error {
	var err error
	d.dev, err = spi.Open(&spi.Devfs{
		Dev:      "/dev/spidev0.0",
		Mode:     spi.Mode0,
		MaxSpeed: 20000000,
	})

	d.dev.SetBitOrder(spi.MSBFirst)
	if err == nil {
		log.Printf("Device Initated : *Success*")
	}
	return err

}

// DevID as described in 7.2.2 Register file: 0x00 – Device Identifier
type DevID struct {
	RIDTAG, MODEL, VER, REV int
}

const NOP byte = 0X00

func (d *DWM1000) DeviceID() string {

	cmd := []byte{0x00, NOP, NOP, NOP, NOP}
	output := make([]byte, len(cmd))
	error := d.dev.Tx(cmd, output)
	if error != nil {
		log.Println("Error reading REG ")
	}
	id := flip(output[1:])

	fmt.Printf("\n DEV_ID Response : %x \n", id)
	return fmt.Sprintf("%x", id)
}

func flip(data []byte) []byte {
	N := len(data)
	result := make([]byte, len(data))
	for i, d := range data {
		result[N-i-1] = d
	}
	return result
}
