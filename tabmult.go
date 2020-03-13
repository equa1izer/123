func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		z01.PrintRune('\n')
		return
	}

	// The parameter will always be a strictly positive number that fits in an int.
	n, _ := strconv.Atoi(args[0])

	for i := 1; i <= 9; i++ {
		// print "i x n = i*n"
		printInt(i)
		printStr(" x ")
		printInt(n)
		printStr(" = ")
		printInt(n * i)
		z01.PrintRune('\n')
	}

}

func printInt(n int) {
	if n >= 10 {
		printInt(n / 10)
	}

	z01.PrintRune(rune(n%10) + '0')
}

func printStr(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}
