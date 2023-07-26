package sjoin

import "strings"

func Join(sep string, vals ...string) string {
	if len(vals) == 0 {
		return ""
	}
	var b strings.Builder
	b.WriteString(vals[0])
	for _, v := range vals[1:] {
		b.WriteString(sep)
		b.WriteString(v)
	}
	return b.String()
}
