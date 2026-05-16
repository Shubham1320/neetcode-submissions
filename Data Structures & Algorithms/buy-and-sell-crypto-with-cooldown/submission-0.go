func maxProfit(prices []int) int {
    dp := make(map[[2]int]int) // key is [i, buying], value is max profit

    var dfs func(i int, buying bool) int
    dfs = func(i int, buying bool) int {
        if i >= len(prices) {
            return 0
        }

        key := [2]int{i, boolToInt(buying)}
        if val, found := dp[key]; found {
            return val
        }

        cooldown := dfs(i + 1, buying)
        if buying {
            buy := dfs(i + 1, false) - prices[i]
            dp[key] = max(buy, cooldown)
        } else {
            sell := dfs(i + 2, true) + prices[i]
            dp[key] = max(sell, cooldown)
        }

        return dp[key]
    }

    return dfs(0, true)
}

func boolToInt(b bool) int {
    if b {
        return 1
    }
    return 0
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
