func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		z01.PrintRune('\n')
		return
	}

	printStr(alphamirror(args[0]))
	z01.PrintRune('\n')
}

func printStr(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}

func alphamirror(s string) string {
	res := ""
	for _, c := range s {
		newRune := c
		if c >= 'a' && c <= 'z' {
			newRune = 'a' + 'z' - c
		} else if c >= 'A' && c <= 'Z' {
			newRune = 'A' + 'Z' - c
		}
		res += string(newRune)
	}

	return res
}
