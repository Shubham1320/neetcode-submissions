


func uniquePaths(m int, n int) int {

    var dfs func(int, int) int
    memo := make([][]int, m)
    for i := 0; i < m; i++ {
        memo[i] = make([]int, n)
        for j := 0; j < n; j++ {
            memo[i][j] = -1
        }
    }

    memo[m-1][n-1] = 1

    fmt.Println(memo)

    dfs = func(i int, j int) int {

        if i >= m || j >= n {
            return 0
        }

        if memo[i][j] != -1 {
            return memo[i][j]
        }

        memo[i][j] = dfs(i+1, j) + dfs(i , j+1)
        fmt.Println(memo)
        return memo[i][j]
    } 

    return dfs(0,0)
}
