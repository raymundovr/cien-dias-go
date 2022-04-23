package conversions

type Farenheit float32
type Celsius float32

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