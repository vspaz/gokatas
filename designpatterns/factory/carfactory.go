package factory

import "fmt"

type cars struct {
	AUDI string
	BMW  string
	FORD string
}

type CarFactory struct {
	makes cars
}

func NewCarFactory() *CarFactory {
	return &CarFactory{
		cars{
			AUDI: "AUDI",
			BMW:  "BMW",
			FORD: "FORD",
		},
	}
}

func (factory *CarFactory) getCar(car string) (ICar, error) {
	switch car {
	case factory.makes.AUDI:
		return NewAUDI(), nil
	case factory.makes.BMW:
		return NewBMW(), nil
	case factory.makes.FORD:
		return NewFord(), nil
	default:
		return nil, fmt.Errorf("'%s' is undefined", car)
	}
}
