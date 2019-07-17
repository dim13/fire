// http://fabiensanglard.net/doom_fire_psx/

package main

import (
	"errors"
	"flag"
	"image"
	"image/color"
	"image/draw"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

var palette = color.Palette{
	color.RGBA{0x07, 0x07, 0x07, 0xff},
	color.RGBA{0x1f, 0x07, 0x07, 0xff},
	color.RGBA{0x2f, 0x0f, 0x07, 0xff},
	color.RGBA{0x47, 0x0f, 0x07, 0xff},
	color.RGBA{0x57, 0x17, 0x07, 0xff},
	color.RGBA{0x67, 0x1f, 0x07, 0xff},
	color.RGBA{0x77, 0x1f, 0x07, 0xff},
	color.RGBA{0x8f, 0x27, 0x07, 0xff},
	color.RGBA{0x9f, 0x2f, 0x07, 0xff},
	color.RGBA{0xaf, 0x3f, 0x07, 0xff},
	color.RGBA{0xbf, 0x47, 0x07, 0xff},
	color.RGBA{0xc7, 0x47, 0x07, 0xff},
	color.RGBA{0xdf, 0x4f, 0x07, 0xff},
	color.RGBA{0xdf, 0x57, 0x07, 0xff},
	color.RGBA{0xdf, 0x57, 0x07, 0xff},
	color.RGBA{0xd7, 0x5f, 0x07, 0xff},
	color.RGBA{0xd7, 0x67, 0x0f, 0xff},
	color.RGBA{0xcf, 0x6f, 0x0f, 0xff},
	color.RGBA{0xcf, 0x77, 0x0f, 0xff},
	color.RGBA{0xcf, 0x7f, 0x0f, 0xff},
	color.RGBA{0xcf, 0X87, 0x17, 0xff},
	color.RGBA{0xc7, 0x87, 0x17, 0xff},
	color.RGBA{0xc7, 0x8f, 0x17, 0xff},
	color.RGBA{0xc7, 0x97, 0x1f, 0xff},
	color.RGBA{0xbf, 0X9f, 0X1f, 0xff},
	color.RGBA{0xbf, 0x9f, 0x1f, 0xff},
	color.RGBA{0xbf, 0xa7, 0x27, 0xff},
	color.RGBA{0xbf, 0xa7, 0x27, 0xff},
	color.RGBA{0xbf, 0xaf, 0x2f, 0xff},
	color.RGBA{0xb7, 0xaf, 0x2f, 0xff},
	color.RGBA{0xb7, 0xb7, 0x2f, 0xff},
	color.RGBA{0xb7, 0xb7, 0x37, 0xff},
	color.RGBA{0xcf, 0xcf, 0x6f, 0xff},
	color.RGBA{0xdf, 0xdf, 0x9f, 0xff},
	color.RGBA{0xef, 0xef, 0xc7, 0xff},
	color.RGBA{0xff, 0xff, 0xff, 0xff},
}

type drawContext struct {
	img *image.Paletted
	off bool
}

func (dc drawContext) pix() []byte {
	img := image.NewRGBA(dc.img.Rect)
	draw.Draw(img, dc.img.Rect, dc.img, image.ZP, draw.Src)
	return img.Pix
}

func newDrawContext(x, y int) *drawContext {
	rand.Seed(time.Now().UnixNano())
	img := image.NewPaletted(image.Rect(0, 0, x, y), palette)
	seed(img, len(palette)-1)
	return &drawContext{img: img}
}

func seed(img *image.Paletted, c int) {
	r := img.Bounds().Max
	for x := 0; x < r.X; x++ {
		img.SetColorIndex(x, r.Y-1, uint8(c))
	}
}

func (dc *drawContext) toggle() {
	if dc.off {
		seed(dc.img, len(palette)-1)
	} else {
		seed(dc.img, 0)
	}
	dc.off = !dc.off
}

func drawTo(img *image.Paletted) {
	r := img.Bounds().Max
	for x := 0; x < r.X; x++ {
		for y := r.Y - 1; y >= 0; y-- {
			z := rand.Intn(3)
			n := img.ColorIndexAt(x, y)
			if n > 0 && z == 1 {
				n--
			}
			img.SetColorIndex(x-z+1, y-1, n)
		}
	}
}

func (dc *drawContext) update(screen *ebiten.Image) error {
	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyQ):
		return errors.New("exit")
	case inpututil.IsKeyJustPressed(ebiten.KeySpace):
		dc.toggle()
	}
	drawTo(dc.img)
	if !ebiten.IsDrawingSkipped() {
		screen.ReplacePixels(dc.pix())
	}
	return nil
}

func main() {
	width := flag.Int("width", 320, "screen width")
	height := flag.Int("height", 200, "screen height")
	scale := flag.Float64("scale", 2.0, "scale")
	flag.Parse()

	dc := newDrawContext(*width, *height)
	ebiten.SetRunnableInBackground(true)
	if err := ebiten.Run(dc.update, *width, *height, *scale, "Fire"); err != nil {
		log.Fatal(err)
	}
}
