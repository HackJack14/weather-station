package temperature

import (
	"os"
	"log"
	"strings"
	"path/filepath"
	"bufio"
	"strconv"
)

type ds18b20 struct {
	path string
	rawTemp int
}

func NewDs18b20() *ds18b20 {
	return &ds18b20{}
}

func (dsb *ds18b20) Begin() bool {
	dsb.path = "/sys/bus/w1/devices/"

	entries, err := os.ReadDir(dsb.path)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		if strings.Contains(entry.Name(), "28") {
			dsb.path = filepath.Join(dsb.path, entry.Name(), "w1_slave")
			return true;
		}
	}
	return false;
}

func (dsb *ds18b20) Read() {
	file, err := os.Open(dsb.path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	tempFound := false
	temp := ""
	lines := 2
	for lines != 0 {
		scanner.Scan()
		_, temp, tempFound = strings.Cut(scanner.Text(), "t=")
		lines -= 1
	}
	if tempFound {
		dsb.rawTemp, err = strconv.Atoi(temp)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (dsb *ds18b20) GetTemperature() float64 {
	return float64(dsb.rawTemp) / 1000
}
