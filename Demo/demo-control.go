package main

import (
	"runtime"

	gm "github.com/edwinhuish/go-miniblink"
	fm "github.com/edwinhuish/go-miniblink/forms"
	cs "github.com/edwinhuish/go-miniblink/forms/controls"
	gw "github.com/edwinhuish/go-miniblink/forms/windows"
)

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	cs.App = new(gw.Provider).Init()
	cs.App.SetIcon("app.ico")

	mb := new(gm.MiniblinkBrowser).Init()
	mb.SetSize(700, 400)
	mb.SetLocation(50, 50)
	//固定4边与父级的距离
	mb.SetAnchor(fm.AnchorStyle_Left | fm.AnchorStyle_Top | fm.AnchorStyle_Right | fm.AnchorStyle_Bottom)

	frm := new(cs.Form).Init()
	frm.SetTitle("普通窗口")
	frm.SetSize(800, 500)
	frm.SetLocation(100, 100)
	frm.SetBgColor(0x2FAEE3)
	frm.AddChild(mb)
	frm.EvLoad["show"] = func(s cs.GUI) {
		mb.LoadUri("https://www.baidu.com")
	}
	cs.Run(frm)
}
