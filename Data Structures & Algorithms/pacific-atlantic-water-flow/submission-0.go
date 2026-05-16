



func pacificAtlantic(heights [][]int) [][]int {

    var ans [][]int

    rows, cols := len(heights), len(heights[0])
    fmt.Println(rows, cols)

    flow := make([][]int, rows)

    for index, _ := range flow {
        flow[index] = make([]int, cols)
    }

    visited := make([][]int, rows)

    for index, _ := range visited {
        visited[index] = make([]int, cols)
    }

    var dfs func(int, int, int)

    dfs = func(i int, j int, prevHeight int) {
        if i < 0 || j < 0 || i >= rows || j >= cols {
            return
        }

        if heights[i][j] < prevHeight || visited[i][j] == 1 {
            return
        }

        fmt.Println(i, j, heights[i][j], prevHeight)

        flow[i][j]++

        visited[i][j] = 1
        dfs(i+1, j, heights[i][j])
        dfs(i-1, j, heights[i][j])
        dfs(i, j+1, heights[i][j])
        dfs(i, j-1, heights[i][j])
    }

    fmt.Println(visited)

    for j := 0; j < cols; j++ {
        dfs(0, j, 0)
    }

    for i := 1; i < rows; i++ {
        dfs(i, 0, 0)
    }

    fmt.Println(flow, visited)

    for i := 0; i < rows; i++ {
        for j:=0; j < cols; j++ {
            visited[i][j] = 0
        }
    }

    for j := 0; j < cols; j++ {
        dfs(rows-1, j, 0)
    }

    for i := 0; i < rows-1; i++ {
        dfs(i, cols-1, 0)
    }

    fmt.Println(flow)

    for i := 0; i < rows; i++ {
        for j:=0; j < cols; j++ {
            if flow[i][j] == 2 {
                ans = append(ans, []int{i,j})
            }
        }
    }


    return ans

}
