
func Clean(arg string) []rune {
    var brackets []rune
    for _, v := range arg {
        switch v {
        case '(', ')', '{', '}', '[', ']':
            brackets = append(brackets, v)
        }
    }
    return brackets
}

func Brackets(brackets []rune) bool {
    var opener, closer rune
    var left, right []rune
    if len(brackets)%2 != 0 {
        return false
    }
    if len(brackets) == 0 {
        return true
    }
    opener = brackets[0]
    switch opener {
    case '(':
        closer = ')'
    case '{':
        closer = '}'
    case '[':
        closer = ']'
    }
    counter := 0
    for i := 1; i < len(brackets); i++ {
        if counter == 0 && brackets[i] == closer {
            left = brackets[1:i]
            right = brackets[i+1:]
            break
        }
        switch brackets[i] {
        case opener:
            counter++
        case closer:
            counter--
        }
        if i == len(brackets)-1 {
            return false
        }
    }
    if !Brackets(left) || !Brackets(right) {
        return false
    }
    return true
}
