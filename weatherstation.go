package main

import (
	"log"
	"time"

	"github.com/HackJack14/weather-station/dht"
)

func main() {
	dht := dht.NewDht20()
	if dht.Begin() {
		for {
			dht.Read()
			log.Print("Humidity: ")
			log.Println(dht.GetHumidity())
			log.Print("Temperature: ")
			log.Println(dht.GetTemperature())
			time.Sleep(time.Second)
		}
	} else {
		log.Println("failed to initialize")
	}
}

