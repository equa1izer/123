func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		z01.PrintRune('\n')
		return
	}

	n, err := strconv.Atoi(args[0])
	if err != nil {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	ans := toHex(n)

	for _, c := range ans {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func toHex(n int) string {

	if n == 0 {
		return "0"
	}

	base := "0123456789abcdef"

	res := ""

	for n != 0 {
		res = string(base[n%len(base)]) + res
		n /= len(base)
	}

	return res
}
