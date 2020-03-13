func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		z01.PrintRune('\n')
		return
	}

	trimmed := trim(args[0])
	lastW := lastWord(trimmed)
	printStr(lastW)
	z01.PrintRune('\n')

	// printStr(lastWord(trim(args[0])) + "\n")
}

func printStr(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
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

func lastWord(s string) string {
	copy := []rune(s)
	for i := len(s) - 1; i >= 0; i-- {
		if copy[i] == ' ' {
			return string(copy[i+1:])
		}

		if i == 0 {
			return s
		}
	}

	return ""
}
