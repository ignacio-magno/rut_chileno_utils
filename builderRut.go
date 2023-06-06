package rut_chileno

type BuilderRut struct {
	Rut
	haveGuion             bool
	haveDots              bool
	HaveDigitoVerificador bool
}

func newBuilderRut(rut string) BuilderRut {
	newRut, err := NewRut(rut)
	if err != nil {
		panic(err)
	}

	return newRut.NewBuilder()
}

func NewBuilderRut(rut Rut) BuilderRut {
	return BuilderRut{
		Rut:                   rut,
		HaveDigitoVerificador: true,
	}
}

func (r BuilderRut) ConGuion() BuilderRut {
	// set guion to last position
	return BuilderRut{
		Rut:       r.Rut[:len(r.Rut)-1] + "-" + r.Rut[len(r.Rut)-1:],
		haveGuion: true,
	}
}

func (r BuilderRut) Build() string {
	return string(r.Rut)
}

func (r BuilderRut) ConPuntos() BuilderRut {
	rut := string(r.Rut)

	if r.haveGuion {
		if len(rut) == 9 {
			rut = rut[:1] + "." + rut[1:4] + "." + rut[4:7] + rut[7:]
		} else {
			rut = rut[:2] + "." + rut[2:5] + "." + rut[5:8] + rut[8:]
		}
	} else {
		if len(rut) == 8 {
			rut = rut[:1] + "." + rut[1:4] + "." + rut[4:7] + rut[7:]
		} else {
			rut = rut[:2] + "." + rut[2:5] + "." + rut[5:8] + rut[8:]
		}
	}

	return BuilderRut{
		Rut:       Rut(rut),
		haveGuion: r.haveGuion,
		haveDots:  true,
	}
}

func (r BuilderRut) SinDigitoVerificador() BuilderRut {
	if r.haveGuion {
		r.Rut = r.Rut[:len(r.Rut)-2]
	} else {
		r.Rut = r.Rut[:len(r.Rut)-1]
	}
	r.HaveDigitoVerificador = false
	return r
}
