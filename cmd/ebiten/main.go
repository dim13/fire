// http://fabiensanglard.net/doom_fire_psx/

package main

import (
	"errors"
	"image"
	"image/draw"

	"github.com/dim13/fire"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	*fire.Fire
}

func (g *Game) Draw(screen *ebiten.Image) {
	draw.Draw(screen, screen.Bounds(), g, image.Point{}, draw.Src)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	p := g.Bounds().Max
	return p.X, p.Y
}

func (g *Game) Update() error {
	g.Next()
	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyQ):
		return errors.New("exit")
	}
	return nil
}

func main() {
	ebiten.SetWindowTitle("Doom Fire")
	ebiten.RunGame(&Game{fire.New(320, 240, fire.Palette)})
}
