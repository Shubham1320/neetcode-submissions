func solve(board [][]byte) {
    rows, cols := len(board), len(board[0])

    visited := make([][]int, rows)
    for i := 0; i < rows; i++ {
        visited[i] = make([]int, cols)
    }

    fmt.Println(visited)

    var dfs func(int, int)

    dfs = func(i int, j int) {

        if i < 0 || j < 0 || i >= rows || j >= cols {
            return
        }

        if visited[i][j] == 1 {
            return
        }

        if board[i][j] == 'X' {
            return
        }

        visited[i][j] = 1

        if board[i][j] == 'O' {
            dfs(i+1, j)
            dfs(i-1, j)
            dfs(i, j+1)
            dfs(i, j-1)
        }
    }

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {

            if i > 0 && i < rows-1 && j > 0 && j < cols-1 {
                continue
            }

            if board[i][j] == 'O' && visited[i][j] == 0 {
                fmt.Println(i,j)
                dfs(i, j)
            }
        }
    }

    fmt.Println(visited)

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            if board[i][j] == 'O' && visited[i][j] == 0 {
                board[i][j] = 'X'
            }
        }
    }
}


//["O","X","X","O","X"],
//["X","O","O","X","O"],
//["X","O","X","O","X"],
//["O","X","O","O","O"],
//["X","X","O","X","O"]
