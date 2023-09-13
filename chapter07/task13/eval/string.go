package eval

import (
	"fmt"
	"strings"
)

//!+String

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%.6g", l)
}

func (u unary) String() string {
	if u.op == '+' {
		return u.x.String()
	}
	if b, ok := u.x.(binary); ok && strings.ContainsRune("+-", b.op) {
		return fmt.Sprintf("-(%s)", u.x)
	}
	return fmt.Sprintf("-%s", u.x)
}

func (b binary) String() string {
	x := b.x.String()
	y := b.y.String()
	if strings.ContainsRune("*/", b.op) {
		if b, ok := b.x.(binary); ok && strings.ContainsRune("+-", b.op) {
			x = fmt.Sprintf("(%s)", x)
		}
	}
	if strings.ContainsRune("*/-", b.op) {
		if _, ok := b.y.(binary); ok {
			y = fmt.Sprintf("(%s)", y)
		}
	}
	return fmt.Sprintf("%s %c %s", x, b.op, y)
}

func (c call) String() string {
	s := fmt.Sprintf("%s(%s", c.fn, c.args[0])
	for _, arg := range c.args[1:] {
		s = s + ", " + arg.String()
	}
	return s + ")"
}

//!-String
