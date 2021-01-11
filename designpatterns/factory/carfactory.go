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

func (cf *CarFactory) getCar(car string) (ICar, error) {
	switch car {
	case cf.makes.AUDI:
		return NewAUDI(), nil
	case cf.makes.BMW:
		return NewBMW(), nil
	case cf.makes.FORD:
		return NewFORD(), nil
	default:
		return nil, fmt.Errorf("'%s' is undefined", car)
	}
}
