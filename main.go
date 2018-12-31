// http://fabiensanglard.net/doom_fire_psx/

package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/disintegration/imaging"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
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

const (
	screenWidth  = 320
	screenHeight = 200
	scale        = 2
)

type drawContext struct {
	img   *image.Paletted
	isOn  bool
	debug bool
}

func newDrawContext(x, y int) *drawContext {
	rand.Seed(time.Now().UnixNano())
	img := image.NewPaletted(image.Rect(0, 0, x, y), palette)
	return &drawContext{img: img}
}

func (dc *drawContext) toggle() *drawContext {
	var c uint8
	if !dc.isOn {
		c = uint8(len(palette) - 1)
	}
	for x := 0; x < screenWidth; x++ {
		dc.img.SetColorIndex(x, 0, c)
	}
	dc.isOn = !dc.isOn
	return dc
}

func (dc *drawContext) update(screen *ebiten.Image) error {
	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyQ):
		return errors.New("exit")
	case inpututil.IsKeyJustPressed(ebiten.KeyD):
		dc.debug = !dc.debug
	case inpututil.IsKeyJustPressed(ebiten.KeySpace):
		dc.toggle()
	}
	for x := 0; x < dc.img.Bounds().Max.X; x++ {
		for y := 1; y < dc.img.Bounds().Max.Y; y++ {
			z := rand.Intn(3)
			n := dc.img.ColorIndexAt(x, y-1)
			if n > 0 && z == 0 {
				n--
			}
			dc.img.SetColorIndex(x-z+1, y, n)
		}
	}
	if !ebiten.IsDrawingSkipped() {
		screen.ReplacePixels(imaging.FlipV(dc.img).Pix)
		if dc.debug {
			ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
		}
	}
	return nil
}

func main() {
	dc := newDrawContext(screenWidth, screenHeight).toggle()
	ebiten.SetRunnableInBackground(true)
	if err := ebiten.Run(dc.update, screenWidth, screenHeight, scale, "Fire"); err != nil {
		log.Fatal(err)
	}
}
