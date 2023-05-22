package conv

// CToK преобразует температуру по Цельсию в температуру по Кельвину
func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

// KToC преобразует температуру по Кельвину в температуру по Цельсию
func KToC(k Kelvin) Celsius { return Celsius(k - FreezingK) }

// MToFt преобразует метры в футы
func MToFt(v Meter) Foot { return Foot(v / MetersInFoot) }

// FtToM преобразует футы в метры
func FtToM(v Foot) Meter { return Meter(v) * MetersInFoot }

// KgToLbs преобразует килограммы в фунты
func KgToLbs(v Kilogram) Pound { return Pound(v / KilogramsInPound) }

// FtToM преобразует фунты в килограммы
func LbsToKg(v Pound) Kilogram { return Kilogram(v) * KilogramsInPound }
