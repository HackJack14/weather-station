package db

import (
    "time"
    "log"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

const dbname string = "weatherstation.db"

type WeatherEntry struct {
    TempOutside float64
    TempInside float64
    Humidity float64
    TimeUnix int
}

type Database struct {
    db *sql.DB
}

func NewDatabase() *Database {
    db, err := sql.Open("sqlite3", dbname)
    if err != nil {
        log.Fatal(err)
    }
    return &Database{
        db,
    }
}

func (data *Database) SaveEntry(outTemp, inTemp, humidity float64) {
    const statement string = `INSERT INTO weatherdata
        VALUES (?, ?, ?, ?);`

    _, err := data.db.Exec(statement, outTemp, inTemp, humidity, time.Now().Unix())
    if err != nil {
        log.Fatal(err)
    }
}

func (data *Database) LoadLatestEntry() WeatherEntry {
    const statement string = `SELECT TempOutside, TempInside, Humidity, MAX(TimeUnix)
        FROM weatherdata`
        
    row := data.db.QueryRow(statement)
    var entry WeatherEntry
    err := row.Scan(&entry.TempOutside, &entry.TempInside, &entry.Humidity, &entry.TimeUnix)
    if err != nil {
        log.Fatal(err)
    }
    return entry
}
