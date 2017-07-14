package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// Implement a Reader type that emits an infinite stream of ASCII character 'A'
func (r MyReader) Read(b []byte) (int, error) {
  for i := range b {
    b[i] = 65
  }
  return len(b), nil
}

func main() {
  reader.Validate(MyReader{})
}