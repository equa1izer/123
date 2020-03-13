func UniqueOccurences(s string) bool {
	m := map[rune]int{}
	for _, c := range s {
		m[c]++
	}

	// fmt.Println(m)

	prev := 0
	for _, n := range m {
		if prev != 0 && prev == n {
			return false
		}

		prev = n
	}

	return true
}
