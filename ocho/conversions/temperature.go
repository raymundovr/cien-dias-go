package conversions

import "fmt"

type Farenheit float64
type Celsius float64

const (
	Freezing Celsius = 0
	Boiling Celsius = 100
	AbsoluteZero Celsius = -273.15
)

func IntoFarenheit(c Celsius) Farenheit {
	return Farenheit( c * 1.8000 + 32 )
}

func IntoCelsius(f Farenheit) Celsius {
	return Celsius( (f - 32) / 1.8000 )
}

func (c Celsius) String() string {
	return fmt.Sprintf("%.2f °C", c)
}

func (c Farenheit) String() string {
	return fmt.Sprintf("%.2f °F", c)
}