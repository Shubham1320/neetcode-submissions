func isMatch(s string, p string) bool {
    m, n := len(s), len(p)
    dp := make([][]int, m+1)
    for i := range dp {
        dp[i] = make([]int, n+1)
        for j := range dp[i] {
            dp[i][j] = -1
        }
    }

    var dfs func(i, j int) bool
    dfs = func(i, j int) bool {
        if j == n {
            return i == m
        }
        if dp[i][j] != -1 {
            return dp[i][j] == 1
        }

        match := i < m && (s[i] == p[j] || p[j] == '.')

        if (j+1) < n && p[j+1] == '*' {
            dp[i][j] = boolToInt(dfs(i, j+2) || (match && dfs(i+1, j)))
            return dp[i][j] == 1
        }

        if match {
            dp[i][j] = boolToInt(dfs(i+1, j+1))
            return dp[i][j] == 1
        }

        dp[i][j] = 0
        return false
    }

    return dfs(0, 0)
}

func boolToInt(b bool) int {
    if b {
        return 1
    }
    return 0
}