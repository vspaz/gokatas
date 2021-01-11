package factory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCarFactory(t *testing.T) {
	carFactory := NewCarFactory()
	ford, _ := carFactory.getCar("FORD")
	assert.Equal(t, &FORD{car:car{make:"", speed:0}}, ford)
}