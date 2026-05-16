


/*func change(amount int, coins []int) int {

    n := len(coins)
    memo := make(map[int]map[int]int)

    for i := 0; i <= amount; i++ {
        memo[i] = make(map[int]int)
    }

    fmt.Println(memo)

    var dfs func(int, int) int

    dfs = func(i int, s int) int {

        fmt.Println(i, s)

        if s < 0 {
            return 0
        }

        if s == 0 {
            return 1
        }

        ans := 0
        for j := i; j < n; j++ {
            coin := coins[j]
            if _, ok := memo[s][coin]; ok {
                ans += memo[s][coin]
            } else {
                memo[s][coin] = dfs(j, s - coin)
                ans += memo[s][coin]
            }
        }
        fmt.Println(memo, ans)

        return ans
    }
    return dfs(0, amount)
}*/

func change(amount int, coins []int) int {
    sort.Ints(coins)
    memo := make([][]int, len(coins) + 1)
    for i := range memo {
        memo[i] = make([]int, amount + 1)
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }

    var dfs func(i, a int) int
    dfs = func(i, a int) int {
        if a == 0 {
            return 1
        }
        if i >= len(coins) {
            return 0
        }
        if memo[i][a] != -1 {
            return memo[i][a]
        }

        res := 0
        if a >= coins[i] {
            res = dfs(i+1, a)
            res += dfs(i, a-coins[i])
        }
        memo[i][a] = res
        return res
    }

    return dfs(0, amount)
}


