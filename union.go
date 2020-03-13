func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		z01.PrintRune('\n')
		return
	}

	printStr(union(args[0], args[1]) + "\n")
}

func printStr(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}

func union(s1, s2 string) string {
	res := ""

	theUltimateString := s1 + s2

	for _, c1 := range theUltimateString {
		if !contains(res, c1) {
			res += string(c1)
		}
	}

	return res
}

func contains(s string, r rune) bool {
	for _, c := range s {
		if c == r {
			return true
		}
	}

	return false
}
//=============================
func main() {
	s1 := "rien"
	s2 := "cette phrase ne cache rien"
	temp := ""
	for _, v := range s1 {
		if uniq(temp, v) {
			temp = temp + string(v)
		}
	}
	for _, v := range s2 {
		if uniq(temp, v) {
			temp = temp + string(v)
		}
	}
	fmt.Println(temp)
}

func uniq(s string, r rune) bool {
	for _, v := range s {
		if v == r {
			return false
		}
	}
	return true
}
