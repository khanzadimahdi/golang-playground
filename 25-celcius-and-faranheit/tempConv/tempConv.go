package tempConv

import "fmt"

type Celsius float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) ToFaranheit() Faranheit {
	return Faranheit((c * 9 / 5) + 32)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

type Faranheit float64

func (f Faranheit) ToCelsius() Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (f Faranheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
