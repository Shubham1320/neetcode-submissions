func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
    prices := make([]int, n)
    for i := range prices {
        prices[i] = math.MaxInt32
    }
    prices[src] = 0

    for i := 0; i <= k; i++ {
        tmpPrices := make([]int, n)
        copy(tmpPrices, prices)

        for _, flight := range flights {
            s, d, p := flight[0], flight[1], flight[2]
            if prices[s] == math.MaxInt32 {
                continue
            }
            if prices[s] + p < tmpPrices[d] {
                tmpPrices[d] = prices[s] + p
            }
        }
        prices = tmpPrices
    }

    if prices[dst] == math.MaxInt32 {
        return -1
    }
    return prices[dst]

}

