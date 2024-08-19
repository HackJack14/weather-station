package main

import (
    "log"
    "time"

    "github.com/HackJack14/weather-station/dht"
    "github.com/HackJack14/weather-station/temperature"
    "github.com/HackJack14/weather-station/database"
)

func main() {
    data := db.NewDatabase()
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
    		data.SaveEntry(dht20Temp, dsbTemp, dht20Humid);
    		time.Sleep(time.Minute)
    	}
    } else {
    	log.Println("failed to initialize")
    }
}

