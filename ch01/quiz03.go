// STOP이란 단어를 입력할때까지 입력된 정수값을 읽어라
package main

import (
  "fmt"
  "os"
  "bufio"
  "strconv"
)

func main() {
  var sum int
  for true {
    f := os.Stdin
    defer f.Close()

    scanner := bufio.NewScanner(f)
    
    if scanner.Scan() {
      input := scanner.Text()
      if input == "STOP" {
        fmt.Println(sum)
        os.Exit(1)
      }
      n, _ := strconv.Atoi(input)
      sum += n
    }
  }
}
