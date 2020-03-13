func Wdmatch(a, b string) string {

    str := ""
    str1 := ""
    for _, v := range a {
        for i, h := range b {
            if v == h {
                b = b[i:]
                str += string(v)
                //    fmt.Println(string(h))
                break
            }
        }
    }
    if a == str {
        return str
    }
    return str1
}
