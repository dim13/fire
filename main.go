// http://fabiensanglard.net/doom_fire_psx/

package main

import (
	"errors"
	"image"
	"image/draw"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type fire struct {
	img  *image.Paletted
	off  Toggle
	gray Toggle
}

func seed(img *image.Paletted, c int) {
	p := img.Bounds().Max
	for x := 0; x < p.X; x++ {
		img.SetColorIndex(x, p.Y-1, uint8(c))
	}
}

func newFire(x, y int) *fire {
	rand.Seed(time.Now().UnixNano())
	img := image.NewPaletted(image.Rect(0, 0, x, y), palette)
	seed(img, len(palette)-1)
	return &fire{
		img: img,
		off: Toggle{
			On:  func() { seed(img, 0) },
			Off: func() { seed(img, len(palette)-1) },
		},
		gray: Toggle{
			On:  func() { img.Palette = toGray(palette) },
			Off: func() { img.Palette = palette },
		},
	}
}

func (f *fire) Draw(screen *ebiten.Image) {
	draw.Draw(screen, screen.Bounds(), f.img, image.Point{}, draw.Src)
}

func (f *fire) Layout(outsideWidth, outsideHeight int) (int, int) {
	p := f.img.Bounds().Max
	return p.X, p.Y
}

func (f *fire) Update() error {
	p := f.img.Bounds().Max
	for x := 0; x < p.X; x++ {
		for y := p.Y - 1; y > 0; y-- {
			z := rand.Intn(3) - 1 // -1, 0, 1
			n := f.img.ColorIndexAt(x, y)
			if n > 0 && z == 0 {
				n-- // next color
			}
			f.img.SetColorIndex(x+z, y-1, n)
		}
	}
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

func main() {
	ebiten.SetWindowTitle("Doom Fire")
	if err := ebiten.RunGame(newFire(320, 240)); err != nil {
		log.Println(err)
	}
}
