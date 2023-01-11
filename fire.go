package main

import (
	"image"
	"image/color"
	"math/rand"
)

type Fire struct {
	*image.Paletted
}

func NewFire(w, h int, p color.Palette) *Fire {
	r := image.Rect(0, 0, w, h)
	img := image.NewPaletted(r, p)
	b := r.Bounds().Max
	for x := 0; x < b.X; x++ {
		img.SetColorIndex(x, b.Y-1, uint8(len(p)-1))
	}
	return &Fire{Paletted: img}
}

func (f *Fire) Next() {
	b := f.Bounds().Max
	for x := 0; x < b.X; x++ {
		for y := b.Y - 1; y > 0; y-- {
			z := rand.Intn(3) - 1 // -1, 0, 1
			n := f.ColorIndexAt(x, y)
			if n > 0 && z == 0 {
				n-- // next color
			}
			f.SetColorIndex(x+z, y-1, n)
		}
	}
}
