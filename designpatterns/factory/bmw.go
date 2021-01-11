package factory

type BMW struct {
	car
}

func NewBMW() ICar {
	return &BMW{
		car: car{
			make:  "",
			speed: 0,
		},
	}
}
