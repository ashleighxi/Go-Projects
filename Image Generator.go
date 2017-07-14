package main

import "golang.org/x/tour/pic"
import "image"
import "image/color"

type Image struct{
  w int
  h int
}

func (v Image) Bounds() image.Rectangle {
  return image.Rect(0,0,v.w,v.h)
}

func (v Image) ColorModel() color.Model {
  return color.RGBAModel
}

func (v Image) At(x, y int) color.Color {
  return color.RGBA{uint8(x),uint8(y),uint8((x+y)/2),255}
}

func main() {
  m := Image{200,200}
  pic.ShowImage(m)
}
