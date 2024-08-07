package main

import (
	"log"
	"time"
	"strconv"

	"github.com/HackJack14/weather-station/dht"
	"github.com/HackJack14/weather-station/temperature"
)

import "github.com/HackJack14/weather-station/database"

func main() {
    dht := dht.NewDht20()
    dsb := temperature.NewDs18b20()
    if dsb.Begin() && dht.Begin() {
    	for {
    		dsb.Read()
    		dht.Read()
    		log.Println("ds18b20 Temperature:")
    		dsbTemp := dsb.GetTemperature()
    		log.Println(dsbTemp)
    		log.Println("dht20 Temperature:")
    		dht20Temp := dht.GetTemperature()
    		log.Println(dht20Temp)
    		log.Println("dht20 Humidity:")
    		dht20Humid := dht.GetHumidity()
    		log.Println(dht20Humid)
            db.Execute("INSERT INTO weatherdata VALUES (" + strconv.FormatFloat(dht20Temp, 'f', -1, 64) + ", " + strconv.FormatFloat(dsbTemp, 'f', -1, 64) + ", " + strconv.FormatFloat(dht20Humid, 'f', -1, 64) + ")")
    		time.Sleep(time.Minute)
    	}
    } else {
    	log.Println("failed to initialize")
    }
}

