package abstract

type WinButton struct{}

func (wb *WinButton) paint() {
	fmt.Printf("Win button painting")
}
