package views

import (
	"com.lierda.wsn.vc/assets"
	"com.lierda.wsn.vc/core"
	"com.lierda.wsn.vc/util"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"log"
)

type LoginView struct {
	window fyne.Window
	app    fyne.App
}

func (v *LoginView) Launch(app fyne.App) {
	v.app = app
	v.window = app.NewWindow("登录")
	v.window.Resize(fyne.NewSize(400, 300))
	v.RenderViewContent()
	v.window.CenterOnScreen()
	v.window.ShowAndRun()
}

func (v *LoginView) RenderViewContent() {
	mainLayout := container.NewMax()
	loginFormBox := container.NewVBox()
	usernameEntry := widget.NewEntry()
	usernameEntry.Move(fyne.NewPos(0, 200))
	usernameEntry.PlaceHolder = "请输入用户名"
	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.PlaceHolder = "请输入密码"
	loginBtn := widget.NewButton("login", func() {
		log.Println(usernameEntry.Text, passwordEntry.Text)
		res := core.Login(core.LoginRequest{Username: usernameEntry.Text, Password: passwordEntry.Text})
		if res != nil {
			util.Token = res.Data.Token
			dialog := dialog.NewInformation("", "登录成功!!!", v.window)
			dialog.SetOnClosed(func() {
				firmwareView := &FirmwareView{}
				firmwareView.Launch(v.app, v.window)
			})
			dialog.Show()
		}
	})
	asset, _ := assets.Asset("static/lierda.png")
	logo := canvas.NewImageFromResource(fyne.NewStaticResource("static/lierda.png", asset))
	logo.FillMode = canvas.ImageFillOriginal
	loginFormBox.Add(container.NewCenter(container.NewMax(logo)))
	loginFormBox.Add(container.NewMax(usernameEntry))
	loginFormBox.Add(container.NewMax(passwordEntry))
	loginFormBox.Add(container.NewCenter(loginBtn))
	mainLayout.Add(container.NewBorder(nil, nil, nil, nil, loginFormBox))
	v.window.SetContent(mainLayout)
}
