package builder

/**
	implement iBuilder
**/
type iglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}

func (nb *iglooBuilder) setWindowType() {
	nb.windowType = "snow window"
}

func (nb *iglooBuilder) setDoorType() {
	nb.doorType = "snow door"
}

func (nb *iglooBuilder) setNumberFloor() {
	nb.floor = 1
}

func (nb *iglooBuilder) getHouse() House {
	return House{
		windowType: nb.windowType,
		doorType:   nb.doorType,
		floor:      nb.floor,
	}
}
