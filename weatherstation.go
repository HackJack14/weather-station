package main

import (
	"log"

	"github.com/HackJack14/weather-station/dht"
)

func main() {
	dht := dht.NewDht20()
	dht.Begin()
	dht.Read()
	log.Println(dht.GetHumidity())
	log.Println(dht.GetTemperature())
}
