func findTargetSumWays(nums []int, target int) int {

    memo := make(map[int]map[int]int)
    n := len(nums)

    for idx, _ := range nums {
        memo[idx] = make(map[int]int)
    }

    var dfs func(int, int) int

    dfs = func(i int, curr_sum int) int {

        if i == n  {
            if curr_sum == target {
                return 1
            }
            return 0
        }

        if _, ok := memo[i][curr_sum]; ok {
            return memo[i][curr_sum]
        }

        memo[i][curr_sum] = dfs(i+1, curr_sum + nums[i]) + dfs(i+1, curr_sum - nums[i]) 

        return memo[i][curr_sum]
    }

    return dfs(0, 0)

}                   