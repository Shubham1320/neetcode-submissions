type Solution struct{}

func (sol *Solution) Encode(strs []string) string {
    var ans string

    if len(strs) == 1 && strs[0] == "" {
        ans = ans + string('#')
        return ans
    }

    if len(strs) == 0 {
        return ans
    }

    for _, str := range strs {
        length := len(str)
        lenStr := strconv.Itoa(length)
        ans = ans + lenStr + string('#') + str
    }
    return ans
}

func (sol *Solution) isNumber(char byte) bool {
    if char >= '0' && char <= '9' {
        return true
    }
    return false
}

func (sol *Solution) Decode(s string) []string {

    fmt.Println(s)

    var ans []string
    var length string

    if s == "" {
        return ans
    }

    if len(s) == 1 && s[0] == '#' {
        ans = append(ans, "")
        return ans
    }

    for i := 0; i < len(s); {
        
        for {
            if s[i] == '#' {
                fmt.Println("break")
                i++
                break
            } else {
                length += string(s[i])
                i++
            }
        }
        fmt.Println(length, i)
        actualLen, _ := strconv.Atoi(length)
        str := s[i:i+actualLen]
        fmt.Println(str)
        length = ""
        ans = append(ans, str)
        i = i + actualLen
        fmt.Println(i)
    }
    return ans

}
