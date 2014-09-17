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
  srcFile, _ := os.Open(os.Args[1])
  src, _, _ := image.Decode(srcFile)

  bounds := src.Bounds()
  sepia := image.NewRGBA(bounds)

  w, h := bounds.Max.X, bounds.Max.Y
  for x:= 0; x < w; x++ {
    for y:= 0; y < h; y++ {
      srcColor := src.At(x, y)

      r, g, b, _ := srcColor.RGBA()

      sr := float64(r) * .393 + float64(g) * .769 + float64(b) * .189
      sg := float64(r) * .349 + float64(g) * .686 + float64(b) * .168
      sb := float64(r) * .272 + float64(g) * .534 + float64(b) * .131

      if sr > 65535.0 {
        sr = 65535.0
      }

      if sg > 65535.0 {
        sg = 65535.0
      }

      if sb > 65535.0 {
        sb = 65535.0
      }

      sepiaColor := color.RGBA64{uint16(sr), uint16(sg), uint16(sb), ^uint16(0)}
      sepia.Set(x, y, sepiaColor)
    }
  }

  out, _ := os.Create("sepia_image.jpg")
  defer out.Close()
  jpeg.Encode(out, sepia, &jpeg.Options{jpeg.DefaultQuality})
}
