package abstract

type winFactory struct {
}

func (wf *winFactory) createButton() IButton {
	return &WinButton{}
}

func (wf *winFactory) createCheckbox() ICheckBox {

}
