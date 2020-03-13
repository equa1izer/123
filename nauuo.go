func Nauuo(plus, minus, rand int) string {
	if plus > minus+rand {
		return "+"
	}

	if minus > plus+rand {
		return "-"
	}

	if plus < minus+rand && plus > minus-rand {
		return "?"
	}

	return "0"
}
