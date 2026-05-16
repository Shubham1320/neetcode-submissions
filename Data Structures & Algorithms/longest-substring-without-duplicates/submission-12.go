func lengthOfLongestSubstring(s string) int {

    mapA := make(map[byte]int)
    var ans int
    ans = 1

    if len(s) <= 1 {
        return len(s)
    }

    mapA[s[0]] = 0

    for i, j := 0, 1; j != len(s); {
        if _, ok := mapA[s[j]]; ok {
            fmt.Println(i,j)
            for {
                delete(mapA, s[i])
                fmt.Println(mapA)
                if s[i] == s[j] {
                    break
                }
                i++
            }
            i++
            mapA[s[j]] = j
            j++
        } else {
            mapA[s[j]] = j
            if len(mapA) > ans {
                ans = len(mapA)
            }
            j++
            fmt.Println(mapA)
        }
    }   
    return ans
}
