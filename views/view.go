package views

import "fyne.io/fyne/v2"

type View interface {
	Launch(app fyne.App)
}
