package bridge

import fm "github.com/edwinhuish/go-miniblink/forms"

type MsgBox interface {
	Show(param fm.MsgBoxParam) fm.MsgBoxResult
}
