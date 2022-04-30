package conversions

import "fmt"

type Farenheit float64
type Celsius float64
type Kelvin float64

type Equivalents struct {
	f Farenheit
	c Celsius
	k Kelvin
}

const (
	Freezing Celsius = 0
	Boiling Celsius = 100
	AbsoluteZero Celsius = -273.15
)

func (c Celsius) IntoFarenheit() Farenheit {
	return Farenheit( c * 1.8 + 32 )
}

func (k Kelvin) IntoFarenheit() Farenheit {
	return Farenheit((k - 273.15) * 1.8 + 32)
}

func (f Farenheit) IntoCelsius() Celsius {
	return Celsius( (f - 32) / 1.8 )
}

func (k Kelvin) IntoCelsius() Celsius {
	return Celsius(k - 273.15)
}

func (c Celsius) IntoKelvin() Kelvin {
	return Kelvin(c + 273.15)
}

func (f Farenheit) IntoKelvin() Kelvin {
	return Kelvin( (f - 32) / 1.8 + 273.15 )
}

func (c Celsius) CalculateEquivalents() Equivalents {
	return Equivalents{ f: c.IntoFarenheit(), c: c, k: c.IntoKelvin() }
}

func (k Kelvin) CalculateEquivalents() Equivalents {
	return Equivalents{ f: k.IntoFarenheit(), c: k.IntoCelsius(), k: k }
}

func (f Farenheit) CalculateEquivalents() Equivalents {
	return Equivalents{ f: f, c: f.IntoCelsius(), k: f.IntoKelvin() }
}

func (c Celsius) String() string {
	return fmt.Sprintf("%.2f °C", c)
}

func (c Farenheit) String() string {
	return fmt.Sprintf("%.2f °F", c)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%.2f °K", k)
}

func (e Equivalents) String() string {
	return fmt.Sprintf("Estas unidades son equivalentes => %s - %s - %s", e.c, e.f, e.k)
}