package rut_chileno_utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithGuion(t *testing.T) {
	builder := newBuilderRut("12345678-9")

	response := builder.ConGuion().Build()

	assert.Equal(t, "12345678-9", response)
}

func TestWithDots(t *testing.T) {
	values := make(map[string]string)
	values["12345678-9"] = "12.345.6789"
	values["12345678-k"] = "12.345.678k"
	values["12345678k"] = "12.345.678k"

	for key, value := range values {
		builder := newBuilderRut(key)
		response := builder.ConPuntos().Build()
		assert.Equal(t, value, response)
	}
}

func TestWithDotAndGuion(t *testing.T) {
	builder := newBuilderRut("12345678-9")
	response := builder.ConPuntos().ConGuion().Build()
	assert.Equal(t, "12.345.678-9", response)
}

func TestEmpty(t *testing.T) {
	builder := newBuilderRut("123456789")
	response := builder.ConPuntos().ConGuion().Build()
	assert.Equal(t, "12.345.678-9", response)
}

func TestWithoutDigitoVerificador(t *testing.T) {
	builder := newBuilderRut("123456789")
	response := builder.ConPuntos().ConGuion().SinDigitoVerificador().Build()
	assert.Equal(t, "12.345.678", response)
}

func TestWithoutDigitoVerificador2(t *testing.T) {
	builder := newBuilderRut("12345678-9")
	response := builder.SinDigitoVerificador().Build()
	assert.Equal(t, "12345678", response)
}
