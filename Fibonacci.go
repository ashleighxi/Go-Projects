package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func(int) int {
  fib := 0
  return func(x int) int {
    x1 := 1
    x2 := 0
    for i := 0; i < x; i++ {
      fib = x1 + x2
      x1 = x2
      x2 = fib
    }
    return fib
  }
}

func main() {
  f := fibonacci()
  for i := 0; i < 15; i++ {
    fmt.Println(f(i))
  }
}
