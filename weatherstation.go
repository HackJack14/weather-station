package main

import (
	"log"
	"time"

	"github.com/HackJack14/weather-station/dht"
	"github.com/HackJack14/weather-station/temperature"
)

func main() {
	dht := dht.NewDht20()
	dsb := temperature.NewDs18b20()
	if dsb.Begin() && dht.Begin() {
		for {
			dsb.Read()
			dht.Read()
			log.Println("ds18b20 Temperature:")
			log.Println(dsb.GetTemperature())
			log.Println("dht20 Temperature:")
			log.Println(dht.GetTemperature())
			log.Println("dht20 Humidity:")
			log.Println(dht.GetHumidity())
			time.Sleep(10 * time.Second)
		}
	} else {
		log.Println("failed to initialize")
	}
}

