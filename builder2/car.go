package builder2

type Car struct {
	seats        int
	engine       string
	tripComputer bool
	gps          bool
}

func (c *Car) GetSeats() int {
	return c.seats
}

func (c *Car) GetEngine() string {
	return c.engine
}

func (c *Car) GetTripComputer() bool {
	return c.tripComputer
}

func (c *Car) GetGPS() bool {
	return c.gps
}
