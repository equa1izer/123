func Lcm(first, second int) int {

	return first * second / Gcd(first, second)
}
