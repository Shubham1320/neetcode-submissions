func isValidSudoku(board [][]byte) bool {
    rowMaps := make([]map[byte]bool, 9)
    colMaps := make([]map[byte]bool, 9)
    squareMaps := make([]map[byte]bool, 9)

    for i := 0; i < 9; i++ {
        rowMaps[i] = make(map[byte]bool)
        colMaps[i] = make(map[byte]bool)
        squareMaps[i] = make(map[byte]bool)
    }

    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[i]); j++ {
            element := board[i][j]
            if element == '.' {
                continue
            }

            if _, ok := rowMaps[i][element]; ok {
                return false
            } else {
                rowMaps[i][element] = true
            }

            if _, ok := colMaps[j][element]; ok {
                return false
            } else {
                colMaps[j][element] = true
            }

            squareNum := 3*int(i/3) + int(j/3)
            if _, ok := squareMaps[squareNum][element]; ok {
                return false
            } else {
                squareMaps[squareNum][element] = true
            }
        }
    }
    return true
}
