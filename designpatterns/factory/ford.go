package factory

type FORD struct {
	car
}

func NewFORD() ICar {
	return &FORD{
		car: car{
			make:  "",
			speed: 0,
		},
	}
}
