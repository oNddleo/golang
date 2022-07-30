package builder2

type carBuilder struct {
	car Car
}

func (cb *carBuilder) reset() {
	cb.car = Car{}
}

func (cb *carBuilder) setSeats() {
	cb.car.seats = 4
}
func (cb *carBuilder) setEngine() {
	cb.car.engine = "Sport Engine"
}

func (cb *carBuilder) setTripComputer() {
	cb.car.tripComputer = true
}

func (cb *carBuilder) setGPS() {
	cb.car.gps = false
}

func (cb *carBuilder) getProduct() Car {

	return Car{
		seats:        cb.car.seats,
		engine:       cb.car.engine,
		tripComputer: cb.car.tripComputer,
		gps:          cb.car.gps,
	}
}
