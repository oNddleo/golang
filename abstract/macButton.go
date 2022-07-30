package abstract

type MacButton struct{}

func (mb *MacButton) paint() {
	fmt.Printf("Mac button painting")
}
