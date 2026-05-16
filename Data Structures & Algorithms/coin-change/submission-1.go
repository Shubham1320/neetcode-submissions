func coinChange(coins []int, amount int) int {
    
    memo := make(map[int]int)

    fmt.Println(memo)

    var dfs func(int) int 

    dfs = func(amount int) int {

        fmt.Println(amount)

        if amount == 0 {
            return 0
        }

        if _, ok := memo[amount]; ok{
            return memo[amount] 
        }

        res := math.MaxInt32
        for _, coin := range coins {
            if amount - coin >= 0 {
                res = min(res, 1 + dfs(amount-coin))
            }
        }
        memo[amount] = res
        return memo[amount]

    }
    ans := dfs(amount)
    if ans >= math.MaxInt32 {
        return -1
    }
    return ans
}      