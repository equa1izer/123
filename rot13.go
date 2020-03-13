func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		z01.PrintRune('\n')
		return
	}

	printStr(rot13(args[0]))

	z01.PrintRune('\n')
}

func printStr(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}

func rot13(s string) string {
	copy := []rune(s)
	for i, char := range copy {
		if (char >= 'a' && char <= 'n'-1) || (char >= 'A' && char <= 'N'-1) {
			copy[i] = char + 13
		} else if (char >= 'n' && char <= 'z') || (char >= 'N' && char <= 'Z') {
			copy[i] = char - 13
		}
	}

	return string(copy)
}
