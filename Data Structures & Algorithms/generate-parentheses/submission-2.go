var ans []string

func generate(str string, open int, close int, n int) {

    if close > open || open > n || close > n {
        return
    }
    
    if open == close && open == n {
        temp := make([]rune, len(str))
        copy(temp, []rune(str))
        ans = append(ans, string(temp))
        return
    }

    generate(str+"(", open+1, close, n)
    generate(str+")", open, close+1, n)

}


func generateParenthesis(n int) []string {

    ans = []string{}

    generate("(", 1, 0, n)

    return ans
}
