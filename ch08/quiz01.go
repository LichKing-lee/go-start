package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// 파일명, 변경대상문자, 변경할문자 3개의 인자를 받아 출력하는 코드를 작성하라
// e.g) go run quiz01.go "sample.txt" hello world
func main() {
	flag.Parse()
	fileName := flag.Arg(0)
	word := flag.Arg(1)
	replaceWord := flag.Arg(2)

	fmt.Println(fileName, word, replaceWord)

	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(strings.Replace(string(content), word, replaceWord, -1))
}
