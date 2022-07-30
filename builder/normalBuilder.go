package builder

/**
	implement iBuilder
**/
type normalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}

func (nb *normalBuilder) setWindowType() {
	nb.windowType = "normal window"
}

func (nb *normalBuilder) setDoorType() {
	nb.doorType = "normal door"
}

func (nb *normalBuilder) setNumberFloor() {
	nb.floor = 3
}

func (nb *normalBuilder) getHouse() House {
	return House{
		windowType: nb.windowType,
		doorType:   nb.doorType,
		floor:      nb.floor,
	}
}
