func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		z01.PrintRune('\n')
		return
	}

	printStr(inter(args[0], args[1]) + "\n")
}

func inter(s1, s2 string) string {
	res := ""

	for _, c := range s1 {
		if !contains(res, c) && contains(s2, c) {
			res += string(c)
		}
	}

	for _, c := range s2 {
		if !contains(res, c) && contains(s1, c) {
			res += string(c)
		}
	}

	return res
}

func printStr(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}

func contains(s string, r rune) bool {
	for _, c := range s {
		if c == r {
			return true
		}
	}

	return false
//============================================
	
	func main() {
	if len(os.Args[1:]) != 2 {
		z01.PrintRune('\n')
		return
	} else {
		a := []rune(os.Args[1])
		b := []rune(os.Args[2])
		c := ""
		for _, v := range a {
			for _, g := range b {
				if v == g {
					c += string(v)
				}
			}
		}
		res := ""
		for _, v := range c {
			if isUnique(v, []rune(res)) {
				res += string(v)
			}
		}
		for _, v := range res {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}

func isUnique(r rune, a []rune) bool {
	for _, v := range a {
		if r == v {
			return false
		}
	}
	return true
}
