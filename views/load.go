package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Loader struct {
}

func (v *Loader) Launch(app fyne.App) {
	window := app.NewWindow("烧写中")
	bar := widget.NewProgressBar()
	window.Resize(fyne.NewSize(400, 30))
	window.SetContent(container.NewMax(bar))
	window.CenterOnScreen()
	window.SetPadded(true)
	window.Show()
}
