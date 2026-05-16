func numDecodings(s string) int {

    n := len(s)
    memo := make([]int, n)

    if s[0] != '0' {
        memo[0] = 1
    }

    for i := 1; i < n; i++ {
        ans := 0
        if i - 2 >= 0 { 
            if s[i-1] != '0' {
                str := string(s[i-1]) + string(s[i])
                if val, _ := strconv.Atoi(str); val <= 26 {
                    ans += memo[i-2]
                }
            }
        } else if s[i-1] != '0' {
            str := string(s[i-1]) + string(s[i])
            if val, _ := strconv.Atoi(str); val <= 26 {
                ans += 1
            }
        }

        if s[i] != '0' {
            ans += memo[i-1]
        }
        memo[i] = ans
    }

    return memo[n-1]
}