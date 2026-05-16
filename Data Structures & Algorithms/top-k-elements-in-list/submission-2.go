func topKFrequent(nums []int, k int) []int {

    countMap := make(map[int]int)
    freqMap := make(map[int][]int)
    maxFreq := 0

    for _, num := range nums {
        if _, ok := countMap[num]; ok {
            countMap[num]++
        } else {
            countMap[num] = 1
        }
    }

    for key, value := range countMap {
        if _, ok := freqMap[value]; ok {
            freqMap[value] = append(freqMap[value], key)
        } else {
            freqMap[value] = []int{key}
        }
        if value > maxFreq {
            maxFreq = value
        }
    }

    fmt.Println(freqMap, maxFreq)

    ans := []int{}
    for {
        if res, ok := freqMap[maxFreq]; ok {
            if len(res) <= k {
                ans = append(ans, res...)
                k = k - len(res)
            } else {
                remaining := res[0:k]
                ans = append(ans, remaining...)
                k = 0
            }
            if k == 0 {
                break
            }
        }
        maxFreq--
    }
    return ans
}
