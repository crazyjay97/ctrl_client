package main

import (
	"com.lierda.wsn.vc/views"
	"fyne.io/fyne/v2/app"
	"github.com/flopp/go-findfont"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	defer os.Unsetenv("FYNE_FONT")
	myApp := app.New()
	view := views.LoginView{}
	view.Launch(myApp)
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
