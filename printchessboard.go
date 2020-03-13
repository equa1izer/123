func PrintChessBoard(x, y int) {
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i%2 == 0 && j%2 == 0 || i%2 == 1 && j%2 == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}
}
