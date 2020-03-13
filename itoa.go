
func Itoa(n int) string {
    nbr := n
    result := ""
    for n != 0 {
        index := n % 10
        if index < 0 {
            index *= -1
        }
        result = string(index+48) + result
        n /= 10
    }
    if nbr < 0 {
        result = "-" + result
    }
    return result
}
