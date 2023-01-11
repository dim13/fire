package main

import (
	"image"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/dim13/fire"
)

type gc struct {
	widget.BaseWidget
	*fire.Fire
	raster *canvas.Raster
}

func newWidget(f *fire.Fire) *gc {
	w := &gc{Fire: f}
	w.raster = canvas.NewRaster(w.draw)
	w.ExtendBaseWidget(w)
	return w
}

func (f *gc) draw(w, h int) image.Image {
	return f.Fire
}

func (f *gc) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(f.raster)
}

func main() {
	a := app.New()
	w := a.NewWindow("Doom Fire")
	f := newWidget(fire.New(320, 200))
	w.SetContent(f)
	w.Resize(fyne.NewSize(320, 200))
	go func() {
		t := time.NewTicker(time.Second / 50)
		defer t.Stop()
		for range t.C {
			f.Next()
			f.Refresh()
		}
	}()
	w.ShowAndRun()
}
