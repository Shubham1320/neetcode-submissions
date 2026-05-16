func countSubstrings(s string) int {

    n := len(s)
    cache := make([][]bool, n)

    ans := 0

    for i := 0; i < n; i++ {
        cache[i] = make([]bool, n)
        cache[i][i] = true
        ans++
    }

    for i := n-1; i >= 0; i-- {

        for j := i+1; j < n; j++ {

            if s[i] == s[j] {

                if i+1 == j || (i+1 <= j-1 && cache[i+1][j-1]){
                    cache[i][j] = true
                    ans++
                }
            } 

        }
    }

    fmt.Println(cache)

    return ans


}
