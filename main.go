// http://fabiensanglard.net/doom_fire_psx/

package main

import (
	"errors"
	"image"
	"image/draw"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type fire struct {
	img   *image.Paletted
	off   bool
	gray  bool
	black int
	white int
	x, y  int
}

func newFire(x, y int) *fire {
	rand.Seed(time.Now().UnixNano())
	f := fire{
		img:   image.NewPaletted(image.Rect(0, 0, x, y), palette),
		black: 0,
		white: len(palette) - 1,
		x:     x,
		y:     y,
	}
	seed(f.img, f.white)
	return &f
}

func seed(img *image.Paletted, c int) {
	r := img.Bounds().Max
	for x := 0; x < r.X; x++ {
		img.SetColorIndex(x, r.Y-1, uint8(c))
	}
}

func (f *fire) toggleGray() {
	p := palette
	if f.gray = !f.gray; f.gray {
		p = toGray(p)
	}
	f.img.Palette = p
}

func (f *fire) toggleOff() {
	color := f.white
	if f.off = !f.off; f.off {
		color = f.black
	}
	seed(f.img, color)
}

func (f *fire) Update(screen *ebiten.Image) error {
	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyQ):
		return errors.New("exit")
	case inpututil.IsKeyJustPressed(ebiten.KeyG):
		f.toggleGray()
	case inpututil.IsKeyJustPressed(ebiten.KeySpace):
		f.toggleOff()
	}
	r := f.img.Bounds().Max
	for x := 0; x < r.X; x++ {
		for y := r.Y - 1; y > 0; y-- {
			z := rand.Intn(3) - 1 // -1, 0, 1
			n := f.img.ColorIndexAt(x, y)
			if n > 0 && z == 0 {
				n-- // next color
			}
			f.img.SetColorIndex(x+z, y-1, n)
		}
	}
	return nil
}

func (f *fire) Draw(screen *ebiten.Image) {
	draw.Draw(screen, screen.Bounds(), f.img, image.Point{}, draw.Src)
}

func (f *fire) Layout(outsideWidth, outsideHeight int) (int, int) {
	return f.x, f.y
}

func main() {
	f := newFire(320, 200)
	ebiten.SetWindowSize(640, 400)
	ebiten.SetWindowTitle("Fire")
	ebiten.SetWindowResizable(true)
	ebiten.SetRunnableInBackground(true)
	if err := ebiten.RunGame(f); err != nil {
		log.Println(err)
	}
}
