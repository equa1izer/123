func RotString(s string) string {
	s = leaveOneSpace(trim(s))
	copy := []rune(s)

	for i := 0; i < len(copy); i++ {
		if copy[i] == ' ' {
			return s[i+1:] + " " + s[:i]
		}

		if i == len(s)-1 {
			return s
		}
	}

	return ""
}

func leaveOneSpace(s string) string {
	res := ""
	copy := []rune(s)
	for i := 0; i < len(copy); i++ {
		if copy[i] == ' ' && copy[i+1] == ' ' {
			continue
		}

		res += string(copy[i])
	}

	return res
}

func trim(s string) string {
	copy := []rune(s)

	if len(copy) == 0 {
		return s
	}

	if copy[0] == ' ' {
		return trim(s[1:])
	}

	if copy[len(copy)-1] == ' ' {
		return trim(s[:len(s)-1])
	}

	return s
}
