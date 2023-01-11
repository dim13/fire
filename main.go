// http://fabiensanglard.net/doom_fire_psx/

package main

import (
	"image"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type fire struct {
	widget.BaseWidget
	*Fire
	raster *canvas.Raster
}

func newWidget(f *Fire) *fire {
	w := &fire{Fire: f}
	w.raster = canvas.NewRaster(w.draw)
	w.ExtendBaseWidget(w)
	return w
}

func (f *fire) draw(w, h int) image.Image {
	return f.Fire
}

func (f *fire) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(f.raster)
}

func main() {
	a := app.New()
	w := a.NewWindow("Doom Fire")
	f := newWidget(NewFire(320, 200, palette))
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
