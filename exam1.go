package main

import (
	"log"
	"os"
	"strconv"
)

// 시험 점수를 입력받아 90 ~ 100점은 A, 80 ~ 89점은 B, 70 ~ 79점은 C, 60 ~ 69점은 D, 나머지 점수는 F를 출력하는 프로그램을 작성하시오.

func main() {
	input := os.Args
	n, err := strconv.Atoi(input[1])
	if err != nil {
		log.Fatal(err)
	}

	log.Println(input[0], "::", examResult(n))
}

func examResult(num int) (str string) {
	switch {
	case num >= 90:
		str = "A"
	case num >= 80:
		str = "B"
	case num >= 70:
		str = "C"
	case num >= 60:
		str = "D"
	default:
		str = "F"
	}

	return
}