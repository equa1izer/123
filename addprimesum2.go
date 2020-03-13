func main() {
    if len(os.Args[1:]) != 1 {
        z01.PrintRune('\n')
        return
    } else {
        n, e := Atoi(os.Args[1])
        if e {
            z01.PrintRune('\n')
            return
        }
        res := 0
        for i := 2; i <= n; i++ {
            if isPrime(i) {
                res += i
            }
        }
        s := Itoa(res)
        for _, v := range s {
            z01.PrintRune(v)
        }
        z01.PrintRune('\n')

    }
}

func isPrime(n int) bool {
    a := 0
    for i := 2; i <= n; i++ {
        if n%i == 0 {
            a++
        }
    }
    if a > 1 {
        return false
    }
    return true
}

func Atoi(s string) (int, bool) {

    neg := false
    if s[0] == '+' {
        s = s[1:]
    } else if s[0] == '-' {
        s = s[1:]
        neg = true
    }
    res := 0
    for _, v := range s {

        if v >= '0' && v <= '9' {
            res = res*10 + int(v-48)
        } else {
            return 0, true
        }
    }
    if neg {
        res *= -1
    }
    return res, false
}

func Itoa(n int) string {
    neg := false
    if n < 0 {
        neg = true
    }
    res := ""
    for n > 0 {
        res = string(n%10+48) + res
        n /= 10
    }
    if neg {
        res = "-" + res
    }
    return res

}
