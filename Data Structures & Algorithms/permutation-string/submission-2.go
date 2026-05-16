func compareMaps(mapA map[byte]int, mapB map[byte]int) bool {

    for key, value := range mapA {
        if val, ok := mapB[key]; !ok || value != val {
            return false
        }
    }
    return true
}


func checkInclusion(s1 string, s2 string) bool {

    if len(s2) < len(s1) {
        return false
    }
    
    targetMap := make(map[byte]int)
    for _, s := range s1 {
        if _, ok := targetMap[byte(s)]; !ok {
            targetMap[byte(s)] = 1
        } else {
            targetMap[byte(s)]++
        }
    }

    mapA := make(map[byte]int)
    for i:=0; i<len(s1); i++ {
        if _, ok := mapA[s2[i]]; !ok {
            mapA[s2[i]] = 1
        } else {
            mapA[s2[i]]++
        }
    }

    if compareMaps(mapA, targetMap) {
        return true
    } 

    for i,j := 1, len(s1); j != len(s2); {
        mapA[s2[i-1]]--
        if mapA[s2[i-1]] == 0 {
            delete(mapA, s2[i-1])
        }

        if _, ok := mapA[s2[j]]; !ok {
            mapA[s2[j]] = 1
        } else {
            mapA[s2[j]]++
        }

        if compareMaps(mapA, targetMap) {
            return true
        }
        i++
        j++
    }
    return false 

}
