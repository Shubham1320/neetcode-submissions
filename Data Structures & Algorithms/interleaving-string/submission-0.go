func isInterleave(s1 string, s2 string, s3 string) bool {

    if len(s1)+len(s2) != len(s3) {
        return false
    }

    m, n := len(s1), len(s2)
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
        for j := range dp[i] {
            dp[i][j] = -1
        }
    }

    var dfs func(i, j, k int) bool
    dfs = func(i, j, k int) bool {
        if k == len(s3) {
            return i == len(s1) && j == len(s2)
        }

        if dp[i][j] != -1 {
            return dp[i][j] == 1
        }

        res := false
        if i < len(s1) && s1[i] == s3[k] {
            res = dfs(i+1, j, k+1)
        }
        if !res && j < len(s2) && s2[j] == s3[k] {
            res = dfs(i, j+1, k+1)
        }

        if res {
            dp[i][j] = 1
        } else {
            dp[i][j] = 0
        }

        return res
    }

    return dfs(0, 0, 0)

}




