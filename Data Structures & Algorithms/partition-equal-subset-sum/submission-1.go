


func canPartition(nums []int) bool {

    n := len(nums)
    full := 0

    sort.Ints(nums)

    memo := make(map[int]map[int]bool)
    for i := 0; i < n; i++ {
        memo[i] = make(map[int]bool)
        full += nums[i]
    }

    if full % 2 != 0 {
        return false
    }

    half := full/2

    var dfs func(int, int) bool

    dfs = func(sum int, i int) bool {

        if i >= n {
            return false
        }

        if sum + nums[i] == half {
            return true
        }

        if _, ok := memo[i][sum]; ok {
            return memo[i][sum]
        }

        ans := false
        if sum + nums[i] < half {
            ans = ans || dfs(sum + nums[i], i+1)
        } else {
            return false
        }

        ans = ans || dfs(sum, i+1)
        memo[i][sum] = ans
        return ans
    }

    return dfs(0,0)

}




    
          
                                    
