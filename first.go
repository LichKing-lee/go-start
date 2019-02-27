package main

import(
  "fmt"
  "strings"
)

func main() {
  s := "hello world"
  fmt.Println(strings.Split(s, " ")[0])
}
