package factory


type iCar interface {
	setName(name string)
	getName() string
	setSpeed(speed int)
	getSpeed() int
}