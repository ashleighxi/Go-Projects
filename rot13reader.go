package main

import (
  "io"
  "os"
  "strings"
  "log"
)

type rot13Reader struct {
  r io.Reader
}

func (c rot13Reader) Read(b []byte) (int, error) {
  n, err := c.r.Read(b)
  if err == io.EOF {
    return 0, io.EOF
  }
  for i:=0; i<n; i++ {
    if b[i] < 91 && b[i]+13 > 90 {
      b[i] = b[i] + 13 - 26
    } else if b[i]+13 > 122 {
      b[i] = b[i] + 13 - 26
    } else if b[i] > 64 && b[i] < 122 {
      b[i] = b[i] + 13
    }
  }
  return len(b), nil
}

func main() {
  s := strings.NewReader("Lbh penpxrq gur pbqr!\n")
  r := rot13Reader{s}
  if _, err := io.Copy(os.Stdout, &r); err != nil {
    log.Fatal(err)
  }
}