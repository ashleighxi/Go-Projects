package main

import (
  "fmt"
  "strings"
)

// This function interprets the SmallF*ck Programming/Eso Language
func Interpreter(code, tape string) string {
  c := strings.Split(code, "")
  x := strings.Split(tape, "")
  var j int
  fmt.Println(c)
  fmt.Println(x)
  bit := 0
  for i := range c {
    if c[i] == ">" {
      bit++
      if bit == len(tape) { return strings.Join(x,"") }
    } else if c[i] == "<" {
      bit--
      if bit < 0 { return strings.Join(x,"") }
    } else if c[i] == "*" {
      if x[bit] == "0" { 
        x[bit] = "1"
      } else {
        x[bit] = "0"
      }
    } else if c[i] == "[" {
      if x[bit] == "0" {
        j = i
        for {
          j++
          if c[j] == "]" { break }
        }
        j++     
        for j < len(c) {
          if c[j] == ">" {
            bit++
            if bit == len(tape) { return strings.Join(x,"") }
          } else if c[j] == "<" {
            bit--
            if bit < 0 { return strings.Join(x,"") }
          } else if c[j] == "*" {
            if x[bit] == "0" { 
              x[bit] = "1"
            } else {
              x[bit] = "0"
            }
          } else if c[j] == "]" {
            if x[bit] == "1" { break }
            break
          }
          j++
          fmt.Println(x)
        }
        break
      }
    }
    
    fmt.Println(x)
      
  }
  return strings.Join(x, "")
  
}

func main() {
  v := Interpreter("*>*>[>*>*>*>]*>*>*", "00001100")
  if v == "11110100" {
    fmt.Println("success")
  } else {
    fmt.Println("Something went wrong")
  }
  fmt.Println(v)
}
