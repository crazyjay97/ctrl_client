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
)

type LoginView struct {
	Window fyne.Window
	app    fyne.App
}

func (v *LoginView) Launch(app fyne.App) {
	v.app = app
	asset, _ := assets.Asset("static/lierda_black.png")
	v.app.SetIcon(fyne.NewStaticResource("static/lierda.png", asset))
	v.Window = app.NewWindow("登录")
	v.Window.SetCloseIntercept(func() {
		if util.InProcess {

		} else {
			close(util.Done)
			v.Window.Close()
		}
	})
	v.Window.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("菜单",
			fyne.NewMenuItem("注销", func() {
				v.RenderViewContent()
			}))))

	v.Window.Resize(fyne.NewSize(400, 300))
	v.RenderViewContent()
	v.Window.CenterOnScreen()
	v.Window.ShowAndRun()
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
		defer func() {
			if err := recover(); err != nil {
				dialog.NewInformation("", "登录失败!!!", v.Window).Show()
			}
		}()
		res := core.Login(core.LoginRequest{Username: usernameEntry.Text, Password: passwordEntry.Text})
		if res != nil {
			util.Token = res.Data.Token
			dialog := dialog.NewInformation("", "登录成功!!!", v.Window)
			dialog.SetOnClosed(func() {
				firmwareView := &FirmwareView{}
				firmwareView.Launch(v.app, v.Window)
			})
			dialog.Show()
		} else {
			dialog.NewInformation("", "登录失败!!!", v.Window).Show()
		}
	})
	asset, _ := assets.Asset("static/lierda_black.png")
	logo := canvas.NewImageFromResource(fyne.NewStaticResource("static/lierda.png", asset))
	logo.FillMode = canvas.ImageFillOriginal
	loginFormBox.Add(container.NewCenter(container.NewMax(logo)))
	loginFormBox.Add(container.NewMax(usernameEntry))
	loginFormBox.Add(container.NewMax(passwordEntry))
	loginFormBox.Add(container.NewCenter(loginBtn))
	mainLayout.Add(container.NewBorder(nil, nil, nil, nil, loginFormBox))
	v.Window.SetContent(mainLayout)
}
