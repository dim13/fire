// http://fabiensanglard.net/doom_fire_psx/

package main

import (
	"errors"
	"flag"
	"image"
	"image/draw"
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type drawContext struct {
	img   *image.Paletted
	off   bool
	gray  bool
	black int
	white int
}

func newDrawContext(x, y int) *drawContext {
	rand.Seed(time.Now().UnixNano())
	dc := drawContext{
		img:   image.NewPaletted(image.Rect(0, 0, x, y), palette),
		black: 0,
		white: len(palette) - 1,
	}
	seed(dc.img, dc.white)
	return &dc
}

func seed(img *image.Paletted, c int) {
	r := img.Bounds().Max
	for x := 0; x < r.X; x++ {
		img.SetColorIndex(x, r.Y-1, uint8(c))
	}
}

func (dc *drawContext) toggleGray() {
	p := palette
	if dc.gray = !dc.gray; dc.gray {
		p = toGray(p)
	}
	dc.img.Palette = p
}

func (dc *drawContext) toggleOff() {
	color := dc.white
	if dc.off = !dc.off; dc.off {
		color = dc.black
	}
	seed(dc.img, color)
}

func (dc *drawContext) drawTo(dst draw.Image) {
	r := dc.img.Bounds().Max
	for x := 0; x < r.X; x++ {
		for y := r.Y - 1; y > 0; y-- {
			z := rand.Intn(3) - 1 // -1, 0, 1
			n := dc.img.ColorIndexAt(x, y)
			if n > 0 && z == 0 {
				n-- // next color
			}
			dc.img.SetColorIndex(x+z, y-1, n)
		}
	}
	draw.Draw(dst, dst.Bounds(), dc.img, image.ZP, draw.Src)
}

func (dc *drawContext) update(screen *ebiten.Image) error {
	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyQ):
		return errors.New("exit")
	case inpututil.IsKeyJustPressed(ebiten.KeyG):
		dc.toggleGray()
	case inpututil.IsKeyJustPressed(ebiten.KeySpace):
		dc.toggleOff()
	}
	if !ebiten.IsDrawingSkipped() {
		dc.drawTo(screen)
	}
	return nil
}

func main() {
	width := flag.Int("width", 320, "screen width")
	height := flag.Int("height", 200, "screen height")
	scale := flag.Float64("scale", 2.0, "scale")
	profile := flag.String("profile", "", "CPU Profile")
	flag.Parse()

	if *profile != "" {
		f, err := os.Create(*profile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	dc := newDrawContext(*width, *height)
	ebiten.SetRunnableInBackground(true)
	if err := ebiten.Run(dc.update, *width, *height, *scale, "Fire"); err != nil {
		log.Println(err)
	}
}
