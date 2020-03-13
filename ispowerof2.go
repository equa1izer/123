func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		z01.PrintRune('\n')
		return
	}

	// When there is only 1 argument, it will always be a positive valid int.
	n, _ := strconv.Atoi(args[0])
	for n > 1 {
		if n%2 == 1 {
			printStr("false\n")
			return
		}

		n /= 2
	}

	printStr("true\n")
}

func printStr(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}
