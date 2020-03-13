func main() {
	args := os.Args[1:]

	s, err := parseArgs(args)

	if err != "" {
		fmt.Println(err)
		return
	}

	bits := []byte("00000000000000000000000000000000")

	for _, c := range s {
		index := int(c - 'a')
		// fmt.Println(index)
		bits[len(bits)-index-1] = '1'
	}

	printBits(bits)
}

func printBits(bits []byte) {
	for i, c := range bits {

		if i != 0 && i%8 == 0 {
			fmt.Print(" ")
		}

		fmt.Print(string(c))

	}

	fmt.Println()
}

func parseArgs(args []string) (string, string) {
	helperMsg := "options: abcdefghijklmnopqrstuvwxyz"

	if len(args) == 0 {
		return "", helperMsg
	}

	s := ""
	for _, arg := range args {
		if len(arg) < 0 || arg[0] != '-' {
			continue
		}

		arg = arg[1:]
		for _, c := range arg {
			if c < 'a' || c > 'z' {
				return "", "Invalid Option"
			}
		}

		s += arg
	}

	if len(s) == 0 || contains(s, 'h') {
		return "", helperMsg
	}

	return s, ""
}

func contains(s string, r rune) bool {
	for _, c := range s {
		if c == r {
			return true
		}
	}

	return false
}
