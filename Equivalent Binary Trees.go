package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
  defer close(ch)  // closes channel when function returns
    var walk func(t *tree.Tree)
    walk = func(t *tree.Tree) {
      if t == nil { return }
      walk(t.Left)
      ch <- t.Value
      walk(t.Right)
    }
  walk(t)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
  ch1 := make(chan int)
  ch2 := make(chan int)
  go Walk(t1, ch1)
  go Walk(t2, ch2)
  
  for {
    i,ok1 := <-ch1
    j,ok2 := <-ch2
    
    if i != j || ok1 != ok2 {
      return false
    }
    
    if !ok1 {
      break
    }
  }
  return true
}

func main() {
  fmt.Println("1 and 2 are same: ", Same(tree.New(1),tree.New(1)))
  fmt.Println("1 and 2 are same: ", Same(tree.New(1),tree.New(2)))
  
}