package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
  var arr = make([][]uint8, dy, dy*dx)
  for i := 0; i < len(arr); i++ {
     arr[i] = make([]uint8, dx, dx)
    
  }
  return arr
}

func main() {
  pic.Show(Pic)
}
