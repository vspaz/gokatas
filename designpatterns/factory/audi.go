package factory

type AUDI struct {
	car
}

func NewAUDI() ICar {
	return &AUDI{
		car: car{
			make:  "",
			speed: 0,
		},
	}
}
