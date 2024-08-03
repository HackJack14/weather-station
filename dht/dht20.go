package dht

import (
	"log"
	"time"

	"periph.io/x/conn/v3/driver/driverreg"
	"periph.io/x/conn/v3/i2c"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/host/v3"
)

type dht20 struct {
	dev     *i2c.Dev
	humRaw  int
	tempRaw int
}

func NewDht20() *dht20 {
	return &dht20{}
}

func (dht *dht20) Begin() bool {
	_, err := host.Init()
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(1 * time.Second)
	_, err = driverreg.Init()
	if err != nil {
		log.Fatal(err)
	}

	bus, err := i2creg.Open("")
	if err != nil {
		log.Fatal(err)
	}

	dht.dev = &i2c.Dev{Addr: 0x38, Bus: bus}

	time.Sleep(100 * time.Millisecond)
	received := make([]byte, 1)
	err = dht.dev.Tx([]byte{0x71}, received)
	if err != nil {
		log.Fatal(err)
	}
	return received[0]&0x18 == 0x18
}

func (dht *dht20) Read() {
	time.Sleep(10 * time.Millisecond)
	err := dht.dev.Tx([]byte{0xAC, 0x33, 0x00}, nil)
	if err != nil {
		log.Fatal(err)
	}
	dataRaw := make([]byte, 7)
	var status byte
	status = 0b10000000
	for status&0b10000000 == 0b10000000 {
		time.Sleep(100 * time.Millisecond)
		err = dht.dev.Tx([]byte{0x71}, dataRaw)
		if err != nil {
			log.Fatal(err)
		}
		status = dataRaw[0]
	}
	middleByte := dataRaw[3] //temp safe so it does not alter the data
	dht.humRaw = ((int(middleByte) & 0b11110000) >> 4) + (int(dataRaw[2]) << 4) + (int(dataRaw[1]) << 12)
	dht.tempRaw = ((int(dataRaw[3]) & 0b00001111) << 16) + (int(dataRaw[4]) << 8) + (int(dataRaw[5]))
}

func (dht *dht20) GetHumidity() float64 {
	return (float64(dht.humRaw) / 0b100000000000000000000) * 100
}

func (dht *dht20) GetTemperature() float64 {
	return (float64(dht.tempRaw)/0b100000000000000000000)*200 - 50
}
