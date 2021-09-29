package views

import (
	"com.lierda.wsn.vc/core"
	"com.lierda.wsn.vc/util"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strconv"
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
						confirm(v, item)
					}),
				),
			)
		}

	}
	return box
}

var currentPort string
var isBLMode = false

func confirm(v *FirmwareView, item core.FirmwareTree) {
	var confirmBtn *widget.Button
	selectWidget := widget.NewSelect(core.ReadSerialList(), func(s string) {
		log.Println("Com Choose:", s)
		currentPort = s
		if s != "" {
			confirmBtn.Enable()
		}
	})
	group := widget.NewRadioGroup([]string{
		"AP MODE",
		"BL MODE",
	}, func(s string) {
		isBLMode = "BL MODE" == s
	})
	group.SetSelected("AP MODE")
	var form = []*widget.FormItem{
		widget.NewFormItem("串口", selectWidget),
		widget.NewFormItem("模式", group),
	}
	confirmDialog, confirmBtn := dialog.NewFormWithBtn("烧录确认", "开始烧录", "取消", form, func(b bool) {
		if b {
			process(v, item)
		}
	}, v.window)
	confirmBtn.Disable()
	confirmDialog.Show()
}

func process(v *FirmwareView, item core.FirmwareTree) {
	bar := widget.NewProgressBar()
	file, err := ioutil.TempFile(".", "vsc-*.d")
	if err != nil {
		log.Println(err)
	}
	defer os.Remove(file.Name())
	query := url.Values{}
	query.Add("id", strconv.Itoa(item.Id))
	util.GetDownload("firmware/download", query, func(reader io.Reader) {
		util.DecryptFile(reader, file)
		process := make(chan int)
		go core.WiSunLoader("vsc-402840195.d", false, isBLMode, process, currentPort[3:])
		customDialog := dialog.NewCustomWithEvent("烧录中", "开始烧录", bar, v.window, func() {

		})
		customDialog.Resize(fyne.NewSize(300, 30))
		customDialog.Show()
		for {
			select {
			case val := <-process:
				bar.SetValue(float64(val) / float64(100))
				bar.Refresh()
				log.Println("getValue", val)
				if val == 100 {
					customDialog.Hide()
					dialog.NewInformation("消息", "烧录成功", v.window).Show()
					return
				} else if val == -1 {
					customDialog.Hide()
					dialog.NewInformation("错误", "烧录失败", v.window).Show()
					return
				}
			}
		}
	})
}
