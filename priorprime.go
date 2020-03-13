func Priorprime(x int) int {

	if x < 2 {
		return 0
	}

	for i := x - 1; i > 2; i-- {
		if isPrime(i) {
			return i + Priorprime(i)
		}
	}

	return 2
}

func isPrime(x int) bool {
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}

	return true
}
