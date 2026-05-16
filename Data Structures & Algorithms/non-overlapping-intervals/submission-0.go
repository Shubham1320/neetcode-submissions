func eraseOverlapIntervals(intervals [][]int) int {
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][1] < intervals[j][1]
    })
    n := len(intervals)
    memo := make([]int, n)

    var dfs func(i int) int
    dfs = func(i int) int {
        if memo[i] != 0 {
            return memo[i]
        }

        res := 1
        for j := i + 1; j < n; j++ {
            if intervals[i][1] <= intervals[j][0] {
                res = max(res, 1+dfs(j))
            }
        }
        memo[i] = res
        return res
    }

    return n - dfs(0)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}