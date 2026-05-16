func climbStairs(n int) int {

    var dfs func(int) int

    cache := make([]int, n+1)
    for i := 0; i <= n; i++ {
        cache[i] = -1
    }

    dfs = func(currentStep int) int {

        if currentStep > n {
            return 0
        }

        if currentStep == n {
            return 1
        }

        if cache[currentStep] != -1 {
            return cache[currentStep]
        }

        cache[currentStep] = dfs(currentStep + 1) + dfs(currentStep + 2)
        return cache[currentStep]
    }

    return dfs(0)


}
      




