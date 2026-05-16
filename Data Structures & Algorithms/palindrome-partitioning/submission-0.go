var ans [][]string

func isPalindrome(s string) bool {

    i := 0
    j := len(s)-1

    for i <= j {

        if s[i] == s[j] {
            i++
            j--
        } else {
            return false
        }
    }
    return true
}


func traverse(i int, subString string, palindromes []string, s string) {

    if i >= len(s) {
        return
    }

    subString = subString + string(s[i])
    if isPalindrome(subString) {
        palindromes = append(palindromes, subString)
        if i == len(s) - 1 {
            fmt.Println(subString, palindromes)
            temp := make([]string, len(palindromes))
            copy(temp, palindromes)
            ans = append(ans, temp)
        }

        traverse(i+1, "", palindromes, s)
        palindromes = palindromes[:len(palindromes)-1]
        traverse(i+1, subString, palindromes, s)
    } else {
        traverse(i+1, subString, palindromes, s)
    }
}

func partition(s string) [][]string {
    ans = [][]string{}
    traverse(0, "", []string{}, s)

    return ans
}
