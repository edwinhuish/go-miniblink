package main

import (
	gm "github.com/edwinhuish/go-miniblink"
	cs "github.com/edwinhuish/go-miniblink/forms/controls"
	gw "github.com/edwinhuish/go-miniblink/forms/windows"
)

func main() {
	cs.App = new(gw.Provider).Init()
	cs.App.SetIcon("app.ico")

	frm := new(gm.MiniblinkForm).Init()
	frm.SetTitle("miniblink窗口")
	frm.SetLocation(100, 100)
	frm.SetSize(800, 500)
	frm.View.SetProxy(gm.ProxyInfo{
		Type:     gm.ProxyType_HTTP,
		HostName: "127.0.0.1",
		Port:     58591,
	})
	frm.EvLoad["show"] = func(s cs.GUI) {
		frm.View.LoadUri("https://www.ip.cn")
	}
	cs.Run(&frm.Form)
}
