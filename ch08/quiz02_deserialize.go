package main

import (
	"encoding/gob"
	"log"
	"os"
)

// encoding/gob 패키지로 구조체에 대한 슬라이스와 Go 맵을 직렬화하고 역직렬화하라
func main() {
	fileName := "serialize"

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	decoder := gob.NewDecoder(f)

	var buses []bus1
	err = decoder.Decode(&buses)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(buses)
}

type bus1 struct {
	BusId int
	BusNo string
}