package main

import (
	"com.lierda.wsn.vc/assets"
	"com.lierda.wsn.vc/util"
	"com.lierda.wsn.vc/views"
	"fyne.io/fyne/v2/app"
	"github.com/flopp/go-findfont"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"syscall"
)

func main() {
	defer os.Unsetenv("FYNE_FONT")
	myApp := app.New()
	view := views.LoginView{}
	view.Launch(myApp)
}

func init() {
	setFont()
	mkHidFolder()
	openFile()
}

func mkHidFolder() {
	_, err := os.Stat(util.HDF)
	if err != nil && !os.IsExist(err) {
		os.Mkdir(util.HDF, os.ModeDir)
	}
	namePtr, _ := syscall.UTF16PtrFromString(util.HDF)
	syscall.SetFileAttributes(namePtr, syscall.FILE_ATTRIBUTE_HIDDEN)

}

func openFile() {
	asset, _ := assets.Asset("static/wisun-loader.exe")
	_, err := os.Stat(util.EXE_PATH)
	if err != nil && !os.IsExist(err) {
		file, _ := os.Create(util.EXE_PATH)
		file.Write(asset)
	}
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
