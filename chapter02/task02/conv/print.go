package conv

import "fmt"

func PrintTKC(t float64) {
	k := Kelvin(t)
	c := Celsius(t)
	fmt.Printf("%s = %s, %s = %s\n",
		k, KToC(k), c, CToK(c))
}

func PrintLMFt(t float64) {
	k := Meter(t)
	c := Foot(t)
	fmt.Printf("%s = %s, %s = %s\n",
		k, MToFt(k), c, FtToM(c))
}

func PrintWKgLb(t float64) {
	k := Kilogram(t)
	c := Pound(t)
	fmt.Printf("%s = %s, %s = %s\n",
		k, KgToLbs(k), c, LbsToKg(c))
}
