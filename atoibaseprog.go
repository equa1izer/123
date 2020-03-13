func AtoiBase(s string, base string) int {
	if !isValid(base) {
		return 0
	}
	res := 0
	for _, x := range s {
		for i, y := range base {
			if x == y {
				res = res*len(base) + i
			}
		}
	}
	return res
}
func isValid(s string) bool {
	if s[0] == '-' || s[0] == '+' {
		return false
	}
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}
	return true
}
