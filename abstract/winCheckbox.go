package abstract

type WinCheckBox struct{}

func (wc *WinCheckBox) paint() {
	fmt.Printf("Win checkbox painting")
}
