package factory

import "fmt"

type cars struct {
	AUDI string
	BMW  string
	FORD string
}


func getCar(car string) (ICar, error) {
	carMakes := &cars{AUDI: "AUDI", BMW: "BMW", FORD: "FORD"}
	switch car {
	case carMakes.AUDI:
		return NewAUDI(), nil
	case carMakes.BMW:
		return NewBMW(), nil
	case carMakes.FORD:
		return NewFord(), nil
	default:
		return nil, fmt.Errorf("'%s' is undefined", car)
	}
}
