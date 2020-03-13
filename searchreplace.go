func main() {
	arg1 := []rune(os.Args[1])
	arg2 := []rune(os.Args[2])
	arg3 := []rune(os.Args[3])

	if len(os.Args) == 4 {
		for i, e := range arg1 {
			if arg1[i] == arg2[0] {
				z01.PrintRune(arg3[0])
			} else {
				z01.PrintRune(e)
			}
		}
	}
	z01.PrintRune('\n')
}
