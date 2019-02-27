// 커맨드라인 인수로 입력된 실수의 평균을 구하라
package main

import (
  "fmt"
  "os"
  "strconv"
)

func main() {
  arguments := os.Args
  size := len(arguments)

  var sum float64
  for i := 1; i < size; i++ {
    n, err := strconv.ParseFloat(arguments[i], 64)

    if err != nil {
      fmt.Println(arguments[i] + " is not a number")
      continue
    }

    sum += n
  }

  fmt.Println(sum / float64(size-1))
}
