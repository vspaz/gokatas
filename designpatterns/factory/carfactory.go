package factory

import "fmt"

func getCar(car string) (iCar, error) {
	switch car {
	case "Ford":
		return NewFord()
	case "BMW":
		return NewBMW()
	case "AUDI":
		return NewAUID()
	default:
		return nil, fmt.Errorf("'%s' is undefined", car)
	}
}
