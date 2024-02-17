package controls

import (
	"github.com/edwinhuish/go-miniblink/forms/bridge"
)

type MainForm interface {
	getFormImpl() bridge.Form
}

var App bridge.Provider

func Run(form MainForm) {
	App.RunMain(form.getFormImpl())
}
