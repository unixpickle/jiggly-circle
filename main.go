package main

import (
	"image/gif"
	"os"

	"github.com/unixpickle/essentials"
)

func main() {
	basis := StandardBasis()
	scene := NewScene(500, 200)

	g := &gif.GIF{}
	for i := 0; i < 50; i++ {
		coeffs := BestFit(basis, scene.Particles)
		f := &CombinedFunc{Basis: basis, Coeffs: coeffs}
		g.Image = append(g.Image, RenderScene(scene, f))
		g.Delay = append(g.Delay, 10)
		scene.Step(0.01)
	}
	for i := len(g.Image) - 1; i >= 0; i-- {
		g.Image = append(g.Image, g.Image[i])
		g.Delay = append(g.Delay, 10)
	}

	f, err := os.Create("output.gif")
	essentials.Must(err)
	defer f.Close()
	essentials.Must(gif.EncodeAll(f, g))
}
