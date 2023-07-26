package minmax

func Min(vals ...int) (res int, ok bool) {
	if len(vals) > 0 {
		ok = true
		res = vals[0]
		for _, v := range vals[1:] {
			if res > v {
				res = v
			}
		}
	}
	return
}

func Max(vals ...int) (res int, ok bool) {
	if len(vals) > 0 {
		ok = true
		res = vals[0]
		for _, v := range vals[1:] {
			if res < v {
				res = v
			}
		}
	}
	return
}
