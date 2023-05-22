package conv

import "fmt"

type Meter float64
type Foot float64

const (
	MetersInFoot Meter = 0.3048
)

func (v Meter) String() string { return fmt.Sprintf("%gm", v) }
func (v Foot) String() string  { return fmt.Sprintf("%gft", v) }
