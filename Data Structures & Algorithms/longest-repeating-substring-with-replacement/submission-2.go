func getMaxFreq(mapA map[byte]int) int {
    ans := 0
    for _, value := range mapA {
        ans = max(ans, value)
    }
    return ans
}


func characterReplacement(s string, k int) int {

    mapA :=  make(map[byte]int)
    ans := 0

    for i,j := 0,0 ; j != len(s); {
        if _, ok := mapA[s[j]]; !ok {
            mapA[s[j]] = 1 
        } else {
            mapA[s[j]]++
        }

        for {
            maxFreq := getMaxFreq(mapA)
            charsToReplace := j - i + 1 - maxFreq
            if charsToReplace <= k {
                if j - i + 1 > ans {
                    ans = j - i + 1
                }
                break
            }
            mapA[s[i]]--
            if mapA[s[i]] == 0 {
                delete(mapA, s[i])
            }
            i++
        }
        j++
    }
    return ans
}
