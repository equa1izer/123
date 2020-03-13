func Gcd(first, second int) int {

	least := first
	if second < first {
		least = second
	}

	for i := 2; i <= least; i++ {
		if second%i == 0 && first%i == 0 {
			return i * Gcd(first/i, second/i)
		}
	}

	return 1
}
