package rut_chileno

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRutCorrectFormatRegex(t *testing.T) {
	ruts := []string{
		"12345678-9",
		"12.345.678-9",
		"12345678-k",
		"12.345.678-k",
		"12345678-K",
		"12.345.678-K",
		"123456789",
		"12345678k",
	}

	for _, rut := range ruts {
		_, err := NewRut(rut)
		assert.NoError(t, err)
	}
}

func TestIncorrectFormat(t *testing.T) {
	ruts := []string{
		"12345678-92",
		"122345678-9",
		"1a2345678-9",
		"1a2345678-x",
	}

	for _, rut := range ruts {
		_, err := NewRut(rut)
		assert.Error(t, err)
	}

}
