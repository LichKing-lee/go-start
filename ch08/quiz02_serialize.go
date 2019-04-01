package main

import (
	"encoding/gob"
	"log"
	"os"
)

// encoding/gob 패키지로 구조체에 대한 슬라이스와 Go 맵을 직렬화하고 역직렬화하라
func main() {
	var buses []bus
	buses = append(buses, bus{1, "750A"})
	buses = append(buses, bus{2, "506"})
	buses = append(buses, bus{3, "501"})

	fileName := "serialize"
	_ = os.Remove(fileName)

	f, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	encoder := gob.NewEncoder(f)
	err = encoder.Encode(buses)
	if err != nil {
		log.Fatalln(err)
	}
}

type bus struct {
	BusId int
	BusNo string
}