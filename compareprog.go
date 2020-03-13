func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		z01.PrintRune('\n')
		return
	}

	if args[0] > args[1] {
		z01.PrintRune('1')
	} else if args[0] < args[1] {
		z01.PrintRune('-')
		z01.PrintRune('1')
	} else {
		z01.PrintRune('0')
	}
	z01.PrintRune('\n')
}
