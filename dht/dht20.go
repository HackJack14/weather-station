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
	host.Init()

	_, err := driverreg.Init()
	if err != nil {
		log.Fatal(err)
	}

	bus, err := i2creg.Open("")
	if err != nil {
		log.Fatal(err)
	}
	defer bus.Close()

	dht.dev = &i2c.Dev{Addr: 38, Bus: bus}

	time.Sleep(100 * time.Millisecond)
	received := make([]byte, 1)
	err = dht.dev.Tx([]byte{0x71}, received)
	if err != nil {
		log.Fatal(err)
	}
	return received[0]&0x18 == 0x18
}

func (dht *dht20) Read() {
	status := make([]byte, 1)
	status[0] = 1
	for status[0]&0b10000000 == 1 {
		time.Sleep(10 * time.Millisecond)
		dht.dev.Write([]byte{0xAC, 0x33, 0x00})
		time.Sleep(80 * time.Millisecond)
		dht.dev.Tx(make([]byte, 0), status)
	}
	dataRaw := make([]byte, 6)
	dht.dev.Tx(make([]byte, 0), dataRaw)
	middleByte := dataRaw[2] //temp safe so it does not alter the data
	dht.humRaw = ((int(middleByte) & 0b11110000) >> 4) + (int(dataRaw[1]) << 4) + (int(dataRaw[0]) << 12)
	dht.tempRaw = ((int(dataRaw[2]) & 0b00001111) << 16) + (int(dataRaw[3]) << 8) + (int(dataRaw[4]))
}

func (dht *dht20) GetHumidity() float32 {
	return (float32(dht.humRaw) / 0b100000000000000000000) * 100
}

func (dht *dht20) GetTemperature() float32 {
	return (float32(dht.tempRaw)/0b100000000000000000000)*200 - 50
}
