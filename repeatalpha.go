func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		z01.PrintRune('\n')
		return
	}

	for _, c := range args[0] {
		repeat := 1

		if c >= 'a' && c <= 'z' {
			repeat += int(c - 'a')
		} else if c >= 'A' && c <= 'Z' {
			repeat += int(c - 'A')
		}

		for i := 0; i < repeat; i++ {
			z01.PrintRune(c)
		}
	}

	z01.PrintRune('\n')
}
