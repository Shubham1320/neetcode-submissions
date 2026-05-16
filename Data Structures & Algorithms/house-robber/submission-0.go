func rob(nums []int) int {

    n := len(nums)
    cache := make([]int, n)

    for i := 0; i < n; i++ {
        cache[i] = -1
    }

    var dfs func(int) int

    dfs = func(i int) int {

        if i >= n {
            return 0
        }

        if cache[i] != -1 {
            return cache[i]
        }

        cache[i] = max(nums[i] + dfs(i+2), dfs(i+1))

        fmt.Println(cache[i])

        return cache[i]
    }

    return dfs(0)

}
