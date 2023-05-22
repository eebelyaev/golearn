package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

type Kelvin float64

const (
	AbsoluteZeroK Kelvin = 0
	FreezingK     Kelvin = 273.15
	BoilingK      Kelvin = 373.15
)

func (k Kelvin) String() string { return fmt.Sprintf("%g°K", k) }
