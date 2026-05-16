func climbStairs(n int) int {

    var dfs func(int)

    ans := 0

    dfs = func(currentStep int) {

        if currentStep > n {
            return
        }

        if currentStep == n {
            ans++
        }

        dfs(currentStep + 1)
        dfs(currentStep + 2)

    }

    dfs(0)
    return ans


}
      




