package rut_chileno_utils

import (
	"fmt"
	"regexp"
	"strings"
)

// Rut representa un rut chileno sin número ni guión.
type Rut struct {
	rut string
}

func (r Rut) NewBuilder() BuilderRut {
	return BuilderRut{
		Rut: r,
	}
}

func NewRut(rut string) (Rut, error) {
	r := Rut{
		rut: removeDot(rut),
	}
	if ok, err := checkRegex(r.rut); !ok {
		return r, fmt.Errorf("rut %s no cumple con el formato esperado: %s", rut, err)
	}
	return r, nil
}

func (r Rut) DigitoVerificador() string {
	return r.rut[len(r.rut)-1:]
}

func removeDot(rut string) string {
	rut = strings.Replace(rut, ".", "", -1)
	rut = strings.ToLower(rut)
	rut = strings.Replace(rut, "-", "", -1)
	return rut
}

func checkRegex(rut string) (bool, error) {
	reg := "^[0-9]{7,8}[0-9kK]$"
	return regexp.MatchString(reg, rut)
}
