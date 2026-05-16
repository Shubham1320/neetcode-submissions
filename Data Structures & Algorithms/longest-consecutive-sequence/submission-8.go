func longestConsecutive(nums []int) int {
    mapA := make(map[int]int)

    for _, num := range nums {
        if _, ok := mapA[num+1]; ok {
            mapA[num] = mapA[num+1]
            delete(mapA, num+1)
        } else {
            if _, ok := mapA[num]; !ok {
                mapA[num] = num
            }
        }
        fmt.Println(mapA)
    }

    for key, value := range mapA {
        fmt.Println(key, value)
        fmt.Println(mapA)
        for {
            if _, ok := mapA[value+1]; ok {
                mapA[key] = mapA[value+1]
                delete(mapA, value+1)
                value = mapA[key]
            } else if key == value {
                break
            } else if _, ok := mapA[value]; ok {
                mapA[key] = mapA[value]
                delete(mapA, value)
                value = mapA[key]
            } else {
                break
            }
        }
    }

    ans := 0
    for key, value := range mapA {
        if value - key + 1 > ans {
            ans = value - key + 1
        }
    }

    return ans
}
