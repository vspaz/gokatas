package factory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCarFactoryFordOK(t *testing.T) {
	carFactory := NewCarFactory()
	ford, _ := carFactory.getCar("FORD")
	_, err := ford.(*FORD)
	assert.Equal(t, nil, err)
	assert.Equal(t, &FORD{car: car{make: "", speed: 0}}, ford)
}

func TestNewCarFactoryAudiOK(t *testing.T) {
	carFactory := NewCarFactory()
	audi, _ := carFactory.getCar("AUDI")
	assert.Equal(t, &AUDI{car: car{make: "", speed: 0}}, audi)
}

func TestNewCarFactoryBMWOK(t *testing.T) {
	carFactory := NewCarFactory()
	bmw, _ := carFactory.getCar("BMW")
	assert.Equal(t, &BMW{car: car{make: "", speed: 0}}, bmw)
}

func TestNewCarFactoryUndefMake(t *testing.T) {
	carFactory := NewCarFactory()
	undefMake, err := carFactory.getCar("undef")
	assert.Equal(t, nil, undefMake)
	assert.NotEqual(t, nil, err)
}

func TestNewCarFactoryMethods(t *testing.T) {
	carFactory := NewCarFactory()
	ford, err := carFactory.getCar("FORD")

	assert.Equal(t, nil, err)

	carMake := "trailblazer"
	ford.setMake(carMake)
	assert.Equal(t, carMake, ford.getMake())

	carSpeed := 250
	ford.setSpeed(carSpeed)
	assert.Equal(t, carSpeed, ford.getSpeed())
}
