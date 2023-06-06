# Doc

Rut representa un string formateado a 123456789

```go
rut, err := NewRut("12.345.678-9") // errr nil
rut, err := NewRut("12.345.678-0") // errr nil
rut, err := NewRut("12345678-0") // errr nil
rut, err := NewRut("123456780") // errr nil
rut, err := NewRut("12.345.678-21") // errr error
rut, err := NewRut("12.34a.678-2") // errr error
```

### Builder Rut Formater
```go
rut, _ := NewRut("123456780")

rut.NewBuilder().ConPuntos().Build() // 12.345.6780
rut.NewBuilder().ConPuntos().ConGuion().Build() // 12.345.678-0
rut.NewBuilder().ConPuntos().ConGuion().SinDigitoVerificador().Build() // 12.345.678

```