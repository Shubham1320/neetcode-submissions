

func lengthOfLIS(nums []int) int {

    memo := make(map[int]int)
    n := len(nums)
    var dfs func(int) int 

    dfs = func(i int) int {

        if i == n {
            return 0
        }

        if _, ok := memo[i]; ok {
            return memo[i]
        }

        ans := 1
        for j := i + 1; j < n; j++ {
            if nums[j] <= nums[i] {
                continue
            }
            ans = max(ans, 1+dfs(j))
        }
        memo[i] = ans
        return memo[i]

    }

    res := 1

    for i := 0; i < n; i++ {
        res = max(res, dfs(i))
    }
    return res

}


            
