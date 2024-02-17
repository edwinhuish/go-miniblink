package main

import (
	gm "github.com/edwinhuish/go-miniblink"
	fm "github.com/edwinhuish/go-miniblink/forms"
	cs "github.com/edwinhuish/go-miniblink/forms/controls"
	gw "github.com/edwinhuish/go-miniblink/forms/windows"
)

func main() {
	cs.App = new(gw.Provider).Init()
	cs.App.SetIcon("app.ico")

	frm := new(gm.MiniblinkForm).Init()
	frm.SetTitle("miniblink窗口")
	frm.SetSize(800, 500)
	frm.SetStartPosition(fm.FormStart_Screen_Center)
	frm.EvLoad["加载网址"] = func(s cs.GUI) {
		frm.View.LoadUri("https://www.baidu.com")
	}
	cs.Run(&frm.Form)
}
