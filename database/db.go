package db

import (
    "os"
    "log"
    "encoding/csv"
)

type Database struct {
    reader *csv.Reader
    writer *csv.Writer
}

func NewDatabase(file *os.File) Database {
    return Database{
        csv.NewReader(file),
        csv.NewWriter(file),
    }
}

func (data *Database) SaveEntry(outTemp, inTemp, humidity string) {
    entry := []string{outTemp, inTemp, humidity}
    err := data.writer.Write(entry)
    if err != nil {
        log.Fatal(err)
    }
    log.Println(entry)
}

func (data *Database) Close() {
    data.writer.Flush()
}
