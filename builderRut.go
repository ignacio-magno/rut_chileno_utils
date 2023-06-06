package rut_chileno_utils

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
		Rut: Rut{
			rut: r.Rut.rut[:len(r.Rut.rut)-1] + "-" + r.Rut.rut[len(r.Rut.rut)-1:],
		},
		haveGuion: true,
	}
}

func (r BuilderRut) Build() string {
	return r.Rut.rut
}

func (r BuilderRut) ConPuntos() BuilderRut {
	rut := r.Rut.rut

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
		Rut: Rut{
			rut: rut,
		},
		haveGuion: r.haveGuion,
		haveDots:  true,
	}
}

func (r BuilderRut) SinDigitoVerificador() BuilderRut {
	if r.haveGuion {
		r.Rut = Rut{
			rut: r.Rut.rut[:len(r.Rut.rut)-2],
		}
	} else {
		r.Rut = Rut{
			rut: r.Rut.rut[:len(r.Rut.rut)-1],
		}
	}
	r.HaveDigitoVerificador = false
	return r
}
