

func generate(i int, j int, board [][]byte, subString []byte, target string, processed map[string]bool) bool {

    if i >= len(board) || j >= len(board[0]) || i < 0 || j < 0 {
        return false
    }

    if _, ok := processed[string(i) + string(j)]; ok {
        return false
    }

    subString = append(subString, board[i][j])
    processed[string(i)+string(j)] = true
    if len(subString) == len(target){
        fmt.Println(string(subString))
        if string(subString) == target {
            fmt.Println(string(subString))
            return true
        }
        delete(processed, string(i)+string(j))
        return false
    } else if len(subString) > len(target) {
        delete(processed, string(i)+string(j))
        return false
    }

    //subString = append(subString, board[i][j])
    //processed[string(i)+string(j)] = true

    ans1 := generate(i+1, j, board, subString, target, processed)
    ans2 := generate(i, j+1, board, subString, target, processed)
    ans3 := generate(i-1, j, board, subString, target, processed)
    ans4 := generate(i, j-1, board, subString, target, processed)
    //subString = subString[0:len(subString)-1]
    delete(processed, string(i)+string(j))
    return ans1 || ans2 || ans3 || ans4
}

func exist(board [][]byte, word string) bool {

    ans := false

    for i, _ := range board {
        for j, _ := range board[i] {
            if board[i][j] == byte(word[0]) {
                fmt.Println(i,j)
                ans = generate(i, j, board, []byte{}, word, map[string]bool{})
                if ans == true {
                    return ans
                }
            }
        }
    }
    return false
}
