package dwm

import (
	"encoding/binary"
	"fmt"
	"log"

	"golang.org/x/exp/io/spi"
)

type DWM1000 struct {
	dev   *spi.Device
	devID DevID
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

func (d *DevID) Fill(data []byte) {
	d.RIDTAG = int(binary.BigEndian.Uint16(data[0:2]))
	fmt.Printf("\n INT Value %d \t %0X", d.RIDTAG, d.RIDTAG)
	d.MODEL = int(uint8(data[2]))
	fmt.Printf("\n INT Value %d ", d.MODEL)

	d.VER = int(uint8((data[3] & 0xF0) >> 4))
	d.REV = int(uint8(data[3] & 0x0F))
	fmt.Printf("\n MODEL Value %d ", d.MODEL)
	fmt.Printf("\n VER Value %d ", d.VER)
	fmt.Printf("\n REV Value %d ", d.REV)
	// fmt.Printf("DEVICE %#v", d)
	// fmt.Printf("DEVICE %v", d)
}

func (d *DevID) String() string {

	return fmt.Sprintf("%0x,Model:%d,Ver:%v,Rev:%v", d.RIDTAG, d.MODEL, d.VER, d.REV)
}

const NOP byte = 0X00

func (d *DWM1000) DeviceID() string {

	cmd := []byte{0x00, NOP, NOP, NOP, NOP}
	output := make([]byte, len(cmd))
	error := d.dev.Tx(cmd, output)
	if error != nil {
		log.Println("Error reading REG ")
	}
	fmt.Printf("\n raw DEV_ID Response : %0x \n", output)

	id := flip(output[1:])
	fmt.Printf("\n DEV_ID Response : %0x \n", id)
	d.devID.Fill(id)
	return fmt.Sprintf("%0X", id)
}

func (d *DWM1000) DevID() DevID {
	return d.devID
}

func flip(data []byte) []byte {
	N := len(data)
	fmt.Printf("\nRecieved .. %0X", data)
	result := make([]byte, len(data))
	for i, d := range data {
		result[N-i-1] = d
	}
	fmt.Printf("\nReturning .. %0X", result)
	return result
}

func (d *DWM1000) EUI_ID() string {

	cmd := []byte{0x01, NOP, NOP, NOP, NOP, NOP, NOP, NOP, NOP}
	output := make([]byte, len(cmd))
	error := d.dev.Tx(cmd, output)
	if error != nil {
		log.Println("Error reading REG ")
	}
	fmt.Printf("\n raw EUI ID Response : %0x \n", output)

	id := flip(output[1:])
	fmt.Printf("\n EUI ID Response : %0x \n", id)

	return fmt.Sprintf("%0X", id)
}

func (d *DWM1000) TxPower() string {
	// 0x1E
	cmd := []byte{0x1E, NOP, NOP, NOP, NOP}
	output := make([]byte, len(cmd))
	error := d.dev.Tx(cmd, output)
	if error != nil {
		log.Println("Error reading REG ")
	}
	fmt.Printf("\n raw TxPower Response : %0x \n", output)

	id := flip(output[1:])
	fmt.Printf("\n TxPower Response : %0x \n", id)

	return fmt.Sprintf("%0x", id)
}
func (d *DWM1000) PANid() string {
	// 0x1E
	cmd := []byte{0x03, NOP, NOP, NOP, NOP}
	output := make([]byte, len(cmd))
	error := d.dev.Tx(cmd, output)
	if error != nil {
		log.Println("Error reading REG ")
	}
	fmt.Printf("\n raw PAN Response : %0x \n", output)

	id := flip(output[1:])
	fmt.Printf("\n PAN Response : %0x \n", id)

	return fmt.Sprintf("%0x", id)
}
