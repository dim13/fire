package main

import (
	"image"
	"image/draw"
	"log"

	"github.com/dim13/fire"
	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/mobile/event/lifecycle"
)

func main() {
	driver.Main(func(s screen.Screen) {
		w, err := s.NewWindow(&screen.NewWindowOptions{Width: 320, Height: 200, Title: "Doom fire"})
		if err != nil {
			log.Fatal(err)
		}
		defer w.Release()
		b, err := s.NewBuffer(image.Pt(320, 200))
		if err != nil {
			log.Fatal(err)
		}
		defer b.Release()
		go func() {
			f := fire.New(320, 200)
			for {
				f.Next()
				draw.Draw(b.RGBA(), b.RGBA().Bounds(), f, image.Point{}, draw.Src)
				w.Upload(image.Point{}, b, b.Bounds())
				w.Publish()
			}
		}()
		for {
			switch e := w.NextEvent().(type) {
			case lifecycle.Event:
				if e.To == lifecycle.StageDead {
					return
				}
			}
		}
	})
}
