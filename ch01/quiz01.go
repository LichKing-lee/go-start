// 커맨드라인 인수로 입력된 숫자 값들을 모두 더하라
package main

import (
  "os"
  "fmt"
  "strconv"
)

func main() {
  arguments := os.Args

  var sum int

  for i := 1; i < len(arguments); i++ {
    n, err := strconv.Atoi(arguments[i])
    
    if err != nil {
      fmt.Println(arguments[i] + " is not a number")
      continue
    }
    sum += n
  }
  fmt.Println(sum)
}
