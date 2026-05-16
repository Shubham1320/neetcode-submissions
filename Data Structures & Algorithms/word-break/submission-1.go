func wordBreak(s string, wordDict []string) bool {

    memo := make(map[int]bool)
    n := len(s)

    var dfs func(int) bool

    if len(s) == 0 {
        return false
    }

    dfs = func(i int) bool {

        if i == n {
            return true
        } 

        if _, ok := memo[i]; ok {
            return memo[i]
        }

        ans := false
        for _, word := range wordDict {
            if i+len(word) <= n && s[i:i+len(word)] == word {
                ans = ans || dfs(i+len(word))
            }
        }
        memo[i] = ans
        return memo[i]
    }

    return dfs(0)
}
