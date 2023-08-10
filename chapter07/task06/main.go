package main

import (
	"flag"
	"fmt"
	"golearn/chapter02/task02/conv"
)

type kelvinFlag struct{ conv.Kelvin }

func (f *kelvinFlag) Set(s string) error {
	var v float64
	var unit string
	fmt.Sscanf(s, "%f%s", &v, &unit)
	switch unit {
	case "K", "°K":
		f.Kelvin = conv.Kelvin(v)
		return nil
	case "C", "°C":
		f.Kelvin = conv.CToK(conv.Celsius(v))
		return nil
	case "F", "°F":
		f.Kelvin = conv.CToK(conv.FToC(conv.Fahrenheit(v)))
		return nil
	}
	return fmt.Errorf("неизвестный тип температуры")
}

func KelvinFlag(name string, temp conv.Kelvin, usage string) *conv.Kelvin {
	var tk kelvinFlag
	tk.Kelvin = temp
	flag.CommandLine.Var(&tk, name, usage)
	return &tk.Kelvin
}

var tempk = KelvinFlag("tempk", 5, "температура по Кельвину")

func main() {
	flag.Parse()
	fmt.Printf("tempk = %v\n", *tempk)
}
