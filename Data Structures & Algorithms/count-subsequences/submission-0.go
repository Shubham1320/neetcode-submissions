func numDistinct(s string, t string) int {

    var dfs func(string, int, int) int
    n := len(s)
    lent := len(t)

    memo := make([][]int, n)
    for idx, _ := range memo {
        memo[idx] = make([]int, lent)
    }

    for i := 0; i < n; i++ {
        for j := 0; j < lent; j++ {
            memo[i][j] = -1
        }
    }

    dfs = func(currString string, i int, j int) int {

        fmt.Println(currString, i, memo)

        if len(currString) == lent {
            if currString == t {
                return 1
            } else {
                return 0
            }
        }

        if i >= n {
            return 0
        }

        if memo[i][j] != -1 {
            return memo[i][j]
        }

        ans := 0
        ans += dfs(currString, i+1, j)
        if s[i] == t[j] {
            ans += dfs(currString + string(s[i]), i+1, j+1)
        }
        memo[i][j] = ans
        return ans
    }

    return dfs("", 0, 0)

}





            


            