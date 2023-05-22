package conv

import "fmt"

type Kilogram float64
type Pound float64

const (
	KilogramsInPound Kilogram = 0.45359237
)

func (v Kilogram) String() string    { return fmt.Sprintf("%gkg", v) }
func (v Pound) String() string { return fmt.Sprintf("%glb", v) }
