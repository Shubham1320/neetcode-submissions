
func numIslands(grid [][]byte) int {

    ans := 0 

    var dfs func(int, int)

    dfs = func(i int, j int) {
        if i >= len(grid) || i < 0 {
            return
        }

        if j >= len(grid[i]) || j < 0 {
            return
        }

        if grid[i][j] == '0'{
            return
        }
        grid[i][j] = '0'
        dfs(i+1, j)
        dfs(i-1, j)
        dfs(i, j+1)
        dfs(i, j-1)
    }

    for i:=0; i < len(grid); i++ {
        for j:=0; j < len(grid[i]); j++ {
            if grid[i][j] == '1' {
                fmt.Println(grid)
                dfs(i, j)
                fmt.Println(grid)
                ans++
            }
        }
    }

    return ans

}
