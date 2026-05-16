func minCostClimbingStairs(cost []int) int {

    cache := make([]int, len(cost))
    for i := 0; i < len(cache); i++ {
        cache[i] = -1
    }

    var dfs func(int) int

    n := len(cost)

    dfs = func(i int) int {

        if i >= n {
            return 0
        }

        if cache[i] != -1 {
            return cache[i]
        }

        cache[i] = min(dfs(i+1), dfs(i+2)) + cost[i]
        fmt.Println(cache[i])
        return cache[i]
    }

    return min(dfs(0), dfs(1))

}
