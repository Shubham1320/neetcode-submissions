var ans [][]string


func permutate(row int, n int, board []string, colMap map[int]bool, posDiag map[int]bool, negDiag map[int]bool) {

    if len(board) == n {
        temp := make([]string, len(board))
        copy(temp, board)
        ans = append(ans, temp)
        return
    }

    for i := 0; i < n; i++ {

        if _, ok := colMap[i]; ok {
            continue
        }

        if _, ok := negDiag[row-i]; ok {
            continue
        }

        if _, ok := posDiag[row+i]; ok {
            continue
        }

        str := ""

        for k := 0; k < n; k++ {
            if i == k {
                str = str + "Q"
            } else {
                str = str + "."
            }
        }

        board = append(board, str)
        colMap[i] = true
        negDiag[row-i] = true
        posDiag[row+i] = true

        fmt.Println(board, colMap, posDiag, negDiag)

        permutate(row+1, n, board, colMap, posDiag, negDiag)

        delete(colMap, i)
        delete(negDiag, row-i)
        delete(posDiag, row+i)
        board = board[:len(board)-1]
    } 

}

func solveNQueens(n int) [][]string {

    ans = [][]string{}

    permutate(0, n, []string{}, map[int]bool{}, map[int]bool{}, map[int]bool{})

    return ans
}
