package builder2

type Director struct {
	builder ICarBuilder
}

func NewDirector(builder ICarBuilder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) SetBuilder(builder ICarBuilder) {
	d.builder = builder
}

func (d *Director) BuildCar() Car {
	d.builder.setSeats()
	d.builder.setEngine()
	d.builder.setTripComputer()
	d.builder.setGPS()
	return d.builder.getProduct()
}
