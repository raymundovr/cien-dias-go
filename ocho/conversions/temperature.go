package conversions

import "fmt"

type Farenheit float64
type Celsius float64
type Kelvin float64

const (
	Freezing Celsius = 0
	Boiling Celsius = 100
	AbsoluteZero Celsius = -273.15
)

func (c Celsius) IntoFarenheit() Farenheit {
	return Farenheit( c * 1.8000 + 32 )
}

func (k Kelvin) IntoFarenheit() Farenheit {
	return k.IntoCelsius().IntoFarenheit()
}

func (f Farenheit) IntoCelsius() Celsius {
	return Celsius( (f - 32) / 1.8000 )
}

func (k Kelvin) IntoCelsius() Celsius {
	return Celsius(k - 273.15)
}

func (c Celsius) IntoKelvin() Kelvin {
	return Kelvin(c + 273.15)
}

func (f Farenheit) IntoKelvin() Kelvin {
	return f.IntoCelsius().IntoKelvin()
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