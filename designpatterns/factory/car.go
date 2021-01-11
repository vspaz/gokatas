package factory

type car struct {
	make  string
	speed int
}

func (c *car) setMake(make string) {
	c.make = make
}

func (c *car) getMake() string {
	return c.make
}

func (c *car) setSpeed(speed int) {
	c.speed = speed
}

func (c *car) getSpeed() int {
	return c.speed
}
