func isAnagram(s string, t string) bool {

    mapA := make(map[rune]int)
    fmt.Println(reflect.TypeOf(mapA))
    for _, val := range s {
        if _, ok := mapA[val]; ok {
            mapA[val]++
        } else {
            mapA[val] = 1
        }
    }

    for _, val := range t {
        if _, ok := mapA[val]; ok {
            mapA[val]--
            if mapA[val] == 0 {
                delete(mapA, val)
            } 
        } else {
            return false
        }
    }

    if len(mapA) != 0 {
        return false
    }
    
    return true
}
