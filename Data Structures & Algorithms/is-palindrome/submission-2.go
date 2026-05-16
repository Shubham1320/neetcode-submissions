
func isAlphaNumeric(char byte) bool {
    if char >= 'a' && char <= 'z' {
        return true
    } else if char >= 'A' && char <= 'Z' {
        return true
    } else if char >= '0' && char <= '9' {
        return true
    }
    return false
}

func toLower(char byte) byte {
    if char >= 'A' && char <= 'Z' {
        return 'a' + (char - 'A')
    }
    return char
}

func isPalindrome(s string) bool {
    i := 0
    j := len(s) - 1

    for {
        if i >= j {
            return true
        }

        if isAlphaNumeric(s[i]) && isAlphaNumeric(s[j]) {
            if toLower(s[i]) == toLower(s[j]) {
                i++
                j--
            } else {
                return false
            }
        } else if !isAlphaNumeric(s[i]) {
            i++
        } else if !isAlphaNumeric(s[j]) {
            j--
        }
    }

}
