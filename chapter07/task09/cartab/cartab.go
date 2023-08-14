package cartab

import (
	"fmt"
)

type Car struct {
	Brand string
	Model string
	Year  uint16
}

func (c *Car) String() string {
	return fmt.Sprintf("{%s %s, %d}", c.Brand, c.Model, c.Year)
}

type StateSort struct {
	Act bool
	Col int
	Asc bool
}

func Compare(c1, c2 *Car, ss StateSort) int {
	if !ss.Act {
		return 0
	}
	if !ss.Asc {
		c1, c2 = c2, c1
	}
	switch ss.Col {
	case 0:
		if c1.Brand == c2.Brand {
			return 0
		} else if c1.Brand < c2.Brand {
			return 1
		}

	case 1:
		if c1.Model == c2.Model {
			return 0
		} else if c1.Model < c2.Model {
			return 1
		}
	case 2:
		if c1.Year == c2.Year {
			return 0
		} else if c1.Year < c2.Year {
			return 1
		}
	}
	return -1
}

type SortCars struct {
	Cars   []*Car
	Orders [3]StateSort
}

func (sc SortCars) Len() int { return len(sc.Cars) }
func (sc SortCars) Less(i, j int) bool {
	for _, order := range sc.Orders {
		if cmp := Compare(sc.Cars[i], sc.Cars[j], order); cmp != 0 {
			return cmp > 0
		}
	}
	return false
}
func (sc SortCars) Swap(i, j int) {
	sc.Cars[i], sc.Cars[j] = sc.Cars[j], sc.Cars[i]
}

var Cars = []*Car{
	{"Audi", "A5", 2019},
	{"Audi", "A6", 2021},
	{"BMW", "M5", 2022},
	{"Lexus", "GS", 2018},
	{"Mersedes", "GL", 2020},
	{"Mersedes", "S", 2021},
	{"Mersedes", "S", 2016},
	{"Lexus", "IS", 2019},
}
