func longestCommonSubsequence(text1 string, text2 string) int {

    n := len(text1)
    m := len(text2)

    memo := make([][]int, n)
    for i := 0; i < n; i++ {
        memo[i] = make([]int, m)
    }

    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if text1[i] == text2[j] {
                if i - 1 >= 0 && j - 1 >= 0 {
                    memo[i][j] = 1 + memo[i-1][j-1]
                } else {
                    memo[i][j] = 1
                }
            } else {
                res := 0
                if i - 1 >= 0 {
                    res = max(res, memo[i-1][j])
                }

                if j - 1 >= 0 {
                    res = max(res, memo[i][j-1])
                }

                memo[i][j] = res
            } 
        }
    }

    return memo[n-1][m-1]

}
