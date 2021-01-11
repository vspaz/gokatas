package factory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCarFactoryFordOK(t *testing.T) {
	carFactory := NewCarFactory()
	ford, _ := carFactory.getCar("FORD")
	assert.Equal(t, &FORD{car:car{make:"", speed:0}}, ford)
}

func TestNewCarFactoryAudiOK(t *testing.T) {
	carFactory := NewCarFactory()
	audi, _ := carFactory.getCar("AUDI")
	assert.Equal(t, &AUDI{car:car{make:"", speed:0}}, audi)
}

