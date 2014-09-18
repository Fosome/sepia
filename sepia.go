//sepia.go
package main

import (
  "os"
  "image"
  "image/color"
  "image/jpeg"

  _ "image/gif" // load for initialization effects
  _ "image/png" // so we can decode gifs and pngs
)

func main() {
  srcFilename := os.Args[1]
  srcFile, _  := os.Open(srcFilename)
  src, _, _   := image.Decode(srcFile)

  bounds := src.Bounds()
  sepia := image.NewRGBA(bounds)

  w, h := bounds.Max.X, bounds.Max.Y
  for x:= 0; x < w; x++ {
    for y:= 0; y < h; y++ {
      sepia.Set(x, y, colorToSepia(src.At(x, y)))
    }
  }

  out, _ := os.Create("sepia_image.jpg")
  defer out.Close()
  jpeg.Encode(out, sepia, &jpeg.Options{jpeg.DefaultQuality})
}

func colorToSepia(src color.Color) color.Color {
  r, g, b, _ := src.RGBA()

  fr := float64(r)
  fg := float64(g)
  fb := float64(b)

  sr := fr * .393 + fg * .769 + fb * .189
  sg := fr * .349 + fg * .686 + fb * .168
  sb := fr * .272 + fg * .534 + fb * .131

  if sr > 65535.0 { sr = 65535.0 }
  if sg > 65535.0 { sg = 65535.0 }
  if sb > 65535.0 { sb = 65535.0 }

  return color.RGBA64{uint16(sr), uint16(sg), uint16(sb), ^uint16(0)}
}
