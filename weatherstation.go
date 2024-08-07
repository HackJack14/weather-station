package main

import (
    // "fmt"
    // "log"
    // "time"
    // "strconv"

    // "github.com/HackJack14/weather-station/dht"
    // "github.com/HackJack14/weather-station/temperature"
)

import "github.com/HackJack14/weather-station/database"

func main() {
    db.Execute("SELECT * FROM weatherdata)
    // dht := dht.NewDht20()
    // dsb := temperature.NewDs18b20()
    // if dsb.Begin() && dht.Begin() {
    // 	for {
    // 		dsb.Read()
    // 		dht.Read()
    // 		log.Println("ds18b20 Temperature:")
    // 		dsbTemp := dsb.GetTemperature()
    // 		log.Println(dsbTemp)
    // 		log.Println("dht20 Temperature:")
    // 		dht20Temp := dht.GetTemperature()
    // 		log.Println(dht20Temp)
    // 		log.Println("dht20 Humidity:")
    // 		dht20Humid := dht.GetHumidity()
    // 		log.Println(dht20Humid)
    // 		sql := fmt.Sprintf("INSERT INTO weatherdata VALUES (%d, %d, %d)", dht20Temp, dsbTemp, dht20Humid)
    // 		log.Println(sql)
    // 		time.Sleep(time.Minute)
    // 	}
    // } else {
    // 	log.Println("failed to initialize")
    // }
}

