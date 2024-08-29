package server

import (
    "net/http"
    "log"
    "io"
    "bytes"
    "encoding/json"
    
    "github.com/HackJack14/weather-station/database"
)

type Server struct {
    
}

func getWeatherData(writer http.ResponseWriter, request *http.Request) {
    data := db.NewDatabase()
    entry := data.LoadLatestEntry()
    body, err := json.Marshal(entry)
    if err != nil {
        log.Fatal(err)
    }
    reader := bytes.NewReader(body)
    io.Copy(writer, reader)
}

func Listen() {
    http.HandleFunc("/weatherdata", getWeatherData)
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}
