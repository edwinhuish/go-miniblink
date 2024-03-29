package controls

import fm "github.com/edwinhuish/go-miniblink/forms"

type GUI interface {
	GetHandle() uintptr
	GetBound() fm.Bound
	SetSize(width, height int)
	SetLocation(x, y int)
	SetBgColor(color int32)
	IsInvoke() bool
	Invoke(fn func())
	InvokeEx(fn func(state interface{}), state interface{})
	Enable(b bool)
	IsEnable() bool
	GetParent() GUI
	GetOwner() GUI
}
