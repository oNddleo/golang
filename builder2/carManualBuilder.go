package builder2

type carManualBuilder struct {
	carManual Car
}

func (cb *carManualBuilder) reset() {
	cb.carManual = Car{}
}

func (cb *carManualBuilder) setSeats() {

}

func (cb *carManualBuilder) setEngine() {

}

func (cb *carManualBuilder) setTripComputer() {

}

func (cb *carManualBuilder) setGPS() {

}

func (cb *carManualBuilder) getProduct() Car {
	return Car{
		seats:        cb.carManual.seats,
		engine:       cb.carManual.engine,
		tripComputer: cb.carManual.tripComputer,
		gps:          cb.carManual.gps,
	}

}
