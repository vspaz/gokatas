package factory

import "fmt"

type cars struct {
	AUDI string
	BMW  string
	FORD string
}


func getCar(car string) (iCar, error) {
	carMakes := &cars{AUDI: "AUDI", BMW: "BMW", FORD: "FORD"}
	switch car {
	case carMakes.AUDI:
		return NewAUID()
	case carMakes.BMW:
		return NewBMW()
	case carMakes.FORD:
		return NewFord()
	default:
		return nil, fmt.Errorf("'%s' is undefined", car)
	}
}
