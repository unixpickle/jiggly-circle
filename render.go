package main

import (
	"image"
	"image/color"
	"math"
)

func RenderScene(s *Scene, f PolarFunc) *image.Paletted {
	size := int(math.Round(s.Radius * 2))
	palette := color.Palette{color.RGBA{0xff, 0xff, 0xff, 0xff}, color.RGBA{0x65, 0xbc, 0xd4, 0xff}}
	img := image.NewPaletted(image.Rect(0, 0, size, size), palette)

	for i := 0; i < size; i++ {
		y := float64(i) - s.Radius
		for j := 0; j < size; j++ {
			x := float64(j) - s.Radius
			distance := (Vector2{X: x, Y: y}).Norm()
			theta := math.Atan2(y, x)
			radius := f.Eval(theta)
			if distance < radius {
				img.SetColorIndex(j, i, 1)
			} else {
				img.SetColorIndex(j, i, 0)
			}
		}
	}

	// for _, p := range s.Particles {
	// 	i := int(math.Round(p.Position.Y + s.Radius))
	// 	j := int(math.Round(p.Position.X + s.Radius))
	// 	img.SetColorIndex(j, i, 1)
	// }

	return img
}
