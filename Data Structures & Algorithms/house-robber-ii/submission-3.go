func rob(nums []int) int {

    n := len(nums)

    if n == 1{
        return nums[0]
    }

    cache := make([][2]int, n)

    for i := 0; i < n; i++ {
        cache[i][0] = -1
        cache[i][1] = -1
    }

    var dfs func(int, int) int

    dfs = func(i int, flag int) int {

        if i >= n {
            return 0
        }

        if i == n-1 {
            if flag == 1 {
                return 0
            }
        }

        if cache[i][flag] != -1 {
            return cache[i][flag]
        }

        cache[i][flag] = max(dfs(i+1, flag), nums[i] + dfs(i+2, flag))

        return cache[i][flag]
    }

    return max(dfs(0, 1), dfs(1, 0))

}
