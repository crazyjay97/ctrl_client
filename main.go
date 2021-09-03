package main

import (
	"com.lierda.wsn.vc/core"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/axgle/mahonia"
	"github.com/flopp/go-findfont"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	defer os.Unsetenv("FYNE_FONT")
	myApp := app.New()
	mainWindow := myApp.NewWindow("菜单")
	layout := container.NewVBox(SerialOpView()...)
	mainWindow.SetContent(layout)
	mainWindow.Resize(fyne.NewSize(800, 500))
	mainWindow.FixedSize()
	mainWindow.ShowAndRun()
}

func SerialOpView() []fyne.CanvasObject {
	boundString := binding.NewString()
	label := widget.NewLabelWithData(boundString)
	label.Resize(fyne.Size{Height: 30})
	contents := ""
	go func() {
		for {
			select {
			case msg := <-core.Msg:
				decoder := mahonia.NewDecoder("gbk")
				log.Println(decoder.ConvertString(msg))
				contents = contents + msg + "/n"
				//label
			}
		}
	}()
	return []fyne.CanvasObject{
		widget.NewSelect(core.ReadSerialList(), func(s string) {
			log.Println("Com Choose:", s)
			err := core.OpenPort(s)
			if err != nil {
				fyne.NewNotification("error", err.Error())
			}
		}),
		widget.NewButton("打开串口", func() {
			core.ReadSerialList()
		}),
		label,
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
	fmt.Println("=============")
	log.Println("Set Font:", fontName)
	fmt.Println("=============")
	os.Setenv("FYNE_FONT", fontName)
}
