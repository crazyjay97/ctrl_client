package views

import (
	"com.lierda.wsn.vc/core"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"log"
)

type FirmwareView struct {
	window fyne.Window
	app    fyne.App
}

func (v *FirmwareView) Launch(app fyne.App, window fyne.Window) {
	v.window = window
	v.app = app
	v.RenderViewContent()
}

func (v *FirmwareView) RenderViewContent() {
	tree := core.GetFirmwareTree()
	v.window.SetTitle("烧录")
	v.window.SetContent(container.NewMax(v.buildTree(tree.Data.Tree, 1)))
}

func (v *FirmwareView) buildTree(tree []core.FirmwareTree, level int) fyne.CanvasObject {
	box := container.NewVBox()
	accordion := widget.NewAccordion()
	box.Add(accordion)
	for _, item := range tree {
		vBox := container.NewVBox()
		space := make([]fyne.CanvasObject, 0)
		for i := 0; i < level; i++ {
			space = append(space, vBox, vBox, vBox, vBox, vBox)
		}
		if item.Type == 1 {
			items := widget.NewAccordionItem(func() string {
				version := ""
				if item.Type == 2 {
					version = "-" + item.Version
				}
				return item.Name + version
			}(), func() fyne.CanvasObject {
				if item.Children != nil && len(item.Children) > 0 {
					log.Println(item.Name, len(item.Children))
					c := container.NewHBox(space...)
					c.Add(v.buildTree(item.Children, level+1))
					return c
				}
				return container.NewVBox()
			}())
			accordion.Append(items)
		} else {
			box.Add(
				container.NewHBox(widget.NewLabel(item.Name+item.Version),
					widget.NewButton("烧录", func() {
						bar := widget.NewProgressBar()
						customDialog := dialog.NewCustomWithEvent("烧录中", "烧录中请勿退出", bar, v.window, func() {
							log.Println("---")
						})
						customDialog.Resize(fyne.NewSize(300, 30))
						customDialog.Show()
						log.Println("id", item.Id)
					})))
		}

	}
	return box
}
