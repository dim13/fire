package main

import (
	"image"
	"image/draw"
	"log"
	"time"

	"github.com/dim13/fire"
	"github.com/zserge/fenster"
)

func main() {
	win, err := fenster.New(320, 200, "Doom Fire")
	if err != nil {
		log.Fatal(err)
	}
	defer win.Close()
	f := fire.New(320, 200)
	for win.Loop(time.Second / 60) {
		if win.Key(27) || win.Key('Q') {
			break
		}
		f.Next()
		draw.Draw(win, win.Bounds(), f, image.Point{}, draw.Src)
	}
}
