package builder2

type ICarBuilder interface {
	reset()
	setSeats()
	setEngine()
	setTripComputer()
	setGPS()
	getProduct() Car
}

func GetCarBuilder() ICarBuilder {
	return &carBuilder{}
}

func GetCarManualBuilder() ICarBuilder {
	return &carManualBuilder{}
}
