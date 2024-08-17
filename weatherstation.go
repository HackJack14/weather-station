package main

import (
    "fmt"
    "log"
    "time"
    "os"

    "github.com/HackJack14/weather-station/dht"
    "github.com/HackJack14/weather-station/temperature"
    "github.com/HackJack14/weather-station/database"
)

func main() {
    file, err := os.OpenFile("data.csv", os.O_APPEND, 0777)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    data := db.NewDatabase(file)
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
    		data.SaveEntry(fmt.Sprintf("%d", dsbTemp), fmt.Sprintf("%d", dht20Temp), fmt.Sprintf("%d", dht20Humid))
    		time.Sleep(time.Minute)
    	}
    } else {
    	log.Println("failed to initialize")
    }
}

