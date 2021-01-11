package factory

type ICar interface {
	setMake(make string)
	getMake() string
	setSpeed(speed int)
	getSpeed() int
}
