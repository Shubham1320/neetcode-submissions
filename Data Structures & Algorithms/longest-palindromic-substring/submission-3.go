func longestPalindrome(s string) string {

    n := len(s)
    cache := make([][]bool, n)

    for i := 0; i < n; i++ {
        cache[i] = make([]bool, n)
        cache[i][i] = true
    }

    strLen := 1
    strIndex := 0

    for i := n-1; i >= 0; i-- {

        for j := i+1; j < n; j++ {

            if s[i] == s[j] {

                if i+1 == j || (i+1 <= j-1 && cache[i+1][j-1]){
                    cache[i][j] = true
                    if j - i + 1 >= strLen {
                        strIndex = i
                        strLen = j - i + 1
                    }
                }
            } 

        }
    }

    fmt.Println(cache)
    fmt.Println(strLen, strIndex)

    return s[strIndex: strIndex+strLen]

}

