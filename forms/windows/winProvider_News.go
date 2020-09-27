package windows

import (
	br "qq2564874169/goMiniblink/forms/bridge"
)

func (_this *Provider) NewForm(param br.FormParam) br.Form {
	return new(winForm).init(_this, param)
}

func (_this *Provider) NewControl() br.Control {
	return new(winControl).init(_this)
}
