package conversions

import "testing"

func TestBodyTemperature(t *testing.T) {
	temp := Celsius(37)
	conv := temp.IntoFarenheit()
	
	if conv.String() != "98.60 °F" {
		t.Fatalf("Error converting body temp into farenheit %s", conv.String())
	}
}