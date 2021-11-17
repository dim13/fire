// http://fabiensanglard.net/doom_fire_psx/

package main

import (
	"errors"
	"image"
	"image/draw"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Fire struct {
	img  *image.Paletted
	off  Toggle
	gray Toggle
}

func (f *Fire) Draw(screen *ebiten.Image) {
	draw.Draw(screen, screen.Bounds(), f.img, image.Point{}, draw.Src)
}

func (f *Fire) Layout(outsideWidth, outsideHeight int) (int, int) {
	p := f.img.Bounds().Max
	return p.X, p.Y
}

func (f *Fire) Update() error {
	Draw(f.img)
	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyQ):
		return errors.New("exit")
	case inpututil.IsKeyJustPressed(ebiten.KeyG):
		f.gray.Toggle()
	case inpututil.IsKeyJustPressed(ebiten.KeySpace):
		f.off.Toggle()
	}
	return nil
}

func Seed(img *image.Paletted, c int) {
	p := img.Bounds().Max
	for x := 0; x < p.X; x++ {
		img.SetColorIndex(x, p.Y-1, uint8(c))
	}
}

func Draw(img *image.Paletted) {
	p := img.Bounds().Max
	for x := 0; x < p.X; x++ {
		for y := p.Y - 1; y > 0; y-- {
			z := rand.Intn(3) - 1 // -1, 0, 1
			n := img.ColorIndexAt(x, y)
			if n > 0 && z == 0 {
				n-- // next color
			}
			img.SetColorIndex(x+z, y-1, n)
		}
	}
}

func New(x, y int) *Fire {
	rand.Seed(cryptoSeed())
	img := image.NewPaletted(image.Rect(0, 0, x, y), palette)
	Seed(img, len(palette)-1)
	return &Fire{
		img: img,
		off: Toggle{
			On:  func() { Seed(img, 0) },
			Off: func() { Seed(img, len(palette)-1) },
		},
		gray: Toggle{
			On:  func() { img.Palette = gray(palette) },
			Off: func() { img.Palette = palette },
		},
	}
}

func main() {
	ebiten.SetWindowTitle("Doom Fire")
	ebiten.RunGame(New(320, 240))
}
