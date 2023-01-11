package main

import (
	"github.com/dim13/fire"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	w, h := 320, 200
	cfg := pixelgl.WindowConfig{
		Title:  "Doom Fire",
		Bounds: pixel.R(0, 0, float64(w), float64(h)),
		VSync:  true,
	}
	f := fire.New(w, h, fire.Palette)
	pixelgl.Run(func() {
		win, err := pixelgl.NewWindow(cfg)
		if err != nil {
			panic(err)
		}
		for !win.Closed() {
			f.Next()
			switch {
			case win.JustPressed(pixelgl.KeyQ):
				return
			}
			p := pixel.PictureDataFromImage(f)
			s := pixel.NewSprite(p, p.Bounds())
			m := pixel.IM.Moved(win.Bounds().Center())
			s.Draw(win, m)
			win.Update()
		}
	})
}
