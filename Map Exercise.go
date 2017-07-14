package main

import (
  "golang.org/x/tour/wc"
  "strings"
)

func WordCount(s string) map[string]int {
  var x = make(map[string]int, len(s))
  var y = strings.Fields(s)
  for i,_ := range y {
    x[y[i]]++
  }
  return x
}

func main() {
  wc.Test(WordCount)
}
