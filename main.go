package main

import (
	"com.lierda.wsn.vc/core"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/axgle/mahonia"
	"github.com/flopp/go-findfont"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var mainWindow fyne.Window

func main() {
	defer os.Unsetenv("FYNE_FONT")
	myApp := app.New()
	mainWindow = myApp.NewWindow("菜单")
	mainWindow.SetOnClosed(func() {
		log.Println("===")
	})
	box := container.NewMax(SerialOpView()...)
	mainWindow.SetContent(box)
	mainWindow.Resize(fyne.NewSize(800, 600))
	mainWindow.ShowAndRun()
}

func SerialOpView() []fyne.CanvasObject {
	progress := widget.NewProgressBar()
	currentPort := ""
	contents := ""
	entry := widget.NewMultiLineEntry()
	entry.MultiLine = true
	entry.OnChanged = func(s string) {
		contents = s
	}
	entry.Wrapping = fyne.TextWrapWord
	var button *widget.Button
	var (
		openPortText  = "打开串口"
		closePortText = "关闭串口"
	)
	button = widget.NewButton(openPortText, func() {
		if button.Text == openPortText {
			/*if currentPort == "" {
				dialog.NewInformation("Warning", "请选择串口!!!", mainWindow).Show()
				return
			}*/
			process := make(chan int)
			//go core.WiSunLoader(".test3", false, false, process)
			go func() {
				for {
					select {
					case val := <-process:
						progress.SetValue(float64(val) / float64(100))
						progress.Refresh()
						log.Println("get value", val)
						if val == 100 {
							return
						}
					}
				}
			}()
			button.SetText(closePortText)
		} else {
			core.ClosePort()
			button.SetText(openPortText)
		}
	})
	// read serial data
	go func() {
		for {
			select {
			case msg := <-core.Msg:
				decoder := mahonia.NewDecoder("gbk")
				contents += decoder.ConvertString(msg) + "\n"
				entry.SetText(contents)
				entry.Refresh()
				entry.CursorRow = len(strings.Split(contents, "\n"))
			}
		}
	}()

	selectWidget := widget.NewSelect(core.ReadSerialList(), func(s string) {
		log.Println("Com Choose:", s)
		currentPort = s
	})
	withoutLayout := container.NewMax(entry)
	sendContent := ""
	bottom := container.NewVBox(
		func() fyne.CanvasObject {
			sendEntry := widget.NewEntry()
			sendEntry.PlaceHolder = "发送内容"
			sendEntry.OnChanged = func(s string) {
				sendContent = s
			}
			return sendEntry
		}(),
		func() fyne.CanvasObject {
			newButton := widget.NewButton("发送", func() {
				if core.Write == nil {
					dialog.NewInformation("Warning", "请打开串口!!!", mainWindow).Show()
				} else {
					core.Write(sendContent)
				}
			})
			return newButton
		}(),
	)
	return []fyne.CanvasObject{
		container.NewBorder(container.NewVBox(selectWidget, button, progress), bottom, nil, nil, withoutLayout),
	}
}

func init() {
	setFont()
}

func setFont() {
	fileList, _ := ioutil.ReadDir("." + string(os.PathSeparator))
	fontName := ""
	for _, info := range fileList {
		if !info.IsDir() && strings.LastIndex(info.Name(), ".ttf") == (len(info.Name())-4) {
			fontName = info.Name()
			break
		}
	}
	if fontName == "" {
		fontPaths := findfont.List()
		for _, path := range fontPaths {
			//楷体:simkai.ttf 黑体:simhei.ttf
			if strings.Contains(path, "simhei.ttf") {
				fontName = path
				break
			}
		}
	}
	log.Println("=============")
	log.Println("Set Font:", fontName)
	log.Println("=============")
	os.Setenv("FYNE_FONT", fontName)
}
