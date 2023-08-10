package tempconv

import (
	"flag"
	"fmt"
	"golearn/chapter02/task02/conv"
)

// *celsiusFlag соответствует интерфейсу flag.Value,
type celsiusFlag struct{ conv.Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // Проверки ошибок не нужны
	switch unit {
	case "C", "°C":
		f.Celsius = conv.Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = conv.FToC(conv.Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("неверная температура %q", s)
}

// CelsiusFlag определяет флаг Celsius с указанным именем, значением
// по умолчанию и строкой-инструкцией по применению и возвращает адрес
// переменной-флага. Аргумент флага должен содержать числовое значение
// и единицу измерения, например "100С".
func CelsiusFlag(name string, value conv.Celsius, usage string) *conv.Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
