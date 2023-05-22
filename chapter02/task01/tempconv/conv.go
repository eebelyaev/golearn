package tempconv

// CToK преобразует температуру по Цельсию в температуру по Кельвину
func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

// KToC преобразует температуру по Кельвину в температуру по Цельсию
func KToC(k Kelvin) Celsius { return Celsius(k - FreezingK) }
