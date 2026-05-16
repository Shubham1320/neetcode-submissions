func maxProfit(prices []int) int {
    ans := 0
    for i, j := 0, 1; j != len(prices); {
        if prices[j] > prices[i] {
            ans = max(ans, prices[j]-prices[i])
            j++
        } else {
            i = j
            j = i + 1
        }
    } 
    return ans

}
