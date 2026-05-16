var ans []string
func generate(currString string, open int, close int, n int) {


    if close > open || open > n{
        return        
    }

    if open == close && open == n {
        fmt.Println(currString)
        ans = append(ans, currString)
        return
    }

    generate(currString + "(", open+1, close, n)
    generate(currString + ")", open, close+1, n)

}


func generateParenthesis(n int) []string {

    ans = []string{}

    generate("(", 1, 0, n)

    return ans
}
