package main

import (
	"fmt"

	gm "github.com/edwinhuish/go-miniblink"
	fm "github.com/edwinhuish/go-miniblink/forms"
	cs "github.com/edwinhuish/go-miniblink/forms/controls"
	gw "github.com/edwinhuish/go-miniblink/forms/windows"
)

func main() {
	cs.App = new(gw.Provider).Init()
	cs.App.SetIcon("app.ico")

	frm := new(cs.Form).Init()
	frm.SetTitle("JS互操作")
	frm.SetSize(800, 500)

	mb := new(gm.MiniblinkBrowser).Init()
	mb.SetAnchor(fm.AnchorStyle_Fill)
	mb.ResourceLoader = append(mb.ResourceLoader, gm.NewFileLoaderStatic("Res", "local"))
	mb.EvConsole["show"] = func(_ *gm.MiniblinkBrowser, e gm.ConsoleEvArgs) {
		fmt.Println(e.Message())
	}
	mb.JsFuncEx("Func1", func(n1, n2 float64) int {
		return int(n1 * n2)
	})
	mb.JsFuncEx("Func2", func(fn gm.JsFunc) {
		fn(5, 6)
	})
	mb.JsFuncEx("Func3", func(param map[string]interface{}) interface{} {
		rs := param["n1"].(float64) * param["n2"].(float64)
		return struct {
			Msg   string
			Value int
		}{
			Msg:   "n1*n2",
			Value: int(rs),
		}
	})
	mb.JsFuncEx("Func5", func() interface{} {
		return func(name string) string {
			return "姓名是：" + name
		}
	})
	frm.AddChild(mb)
	frm.EvLoad["show"] = func(s cs.GUI) {
		mb.LoadUri("https://local/js_call.html")
	}
	cs.Run(frm)
}
