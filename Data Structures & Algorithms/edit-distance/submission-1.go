func minDistance(word1 string, word2 string) int {


    var dfs func(string, int, int) int
    len1 := len(word1)
    len2 := len(word2)

    dp := make([][]int, len1)
    for i := range dp {
        dp[i] = make([]int, len2)
        for j := range dp[i] {
            dp[i][j] = -1
        }
    }


    dfs = func(currString string, i int, j int) int {

        fmt.Println(currString, i, j)

        if i == len1 {
            return len2 - j
        }
        if j == len2 {
            return len1 - i
        }

        if dp[i][j] != -1 {
            return dp[i][j]
        }
        ans := math.MaxInt

        if word1[i] == word2[j] {
            ans1 := dfs(currString + string(word1[i]), i+1, j+1)
            if ans1 != -1 {
                ans = min(ans, ans1)
            }
        } else {
            //insert
            ans1 := dfs(currString + string(word2[j]), i, j+1)
            if ans1 != -1 {
                ans = min(ans, 1 + ans1)
            }
            
            ans2 := dfs(currString + string(word2[j]), i+1, j+1)
            if ans2 != -1 {
                ans = min(ans, 1 + ans2)
            }

            ans3 := dfs(currString, i+1, j)
            if ans3 != -1 {
                ans = min(ans, 1 + ans3)
            }
        }

        if ans != math.MaxInt {
            dp[i][j] = ans
            return ans
        }
        return -1
    }

    return dfs("", 0 , 0)

}
