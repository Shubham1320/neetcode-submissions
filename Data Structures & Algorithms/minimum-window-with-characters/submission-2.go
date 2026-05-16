func compareMaps(mapA map[byte]int, mapB map[byte]int) bool {
    for key, value := range mapA {
        if _, ok := mapB[key]; !ok {
            return false
        } else if value > mapB[key]{
            return false
        }
    } 
    return true
}



func minWindow(s string, t string) string {

    ans := ""

    mapT := make(map[byte]int)
    for _, char := range t {
        if _, ok := mapT[byte(char)]; ok {
            mapT[byte(char)]++
        } else {
            mapT[byte(char)] = 1
        }
    }

    currentMap := make(map[byte]int)

    i := 0
    j := 0
    for {
        if i == len(s) || j == len(s) {
            return ans
        } 
        if _, ok := mapT[s[i]]; ok {
            break
        }
        i++
        j++
    }
    fmt.Println(mapT, i, j)
    currentMap[s[j]] = 1
    for {
        fmt.Println(i,j)

        if i == len(s) || j == len(s) {
            return ans
        }

        if compareMaps(mapT, currentMap) {
                
            if len(s[i:j+1]) < len(ans) || ans == ""{
                ans = s[i:j+1]
                fmt.Println(ans)
            }

            if _, ok := currentMap[s[i]]; ok{
                currentMap[s[i]]--
                if currentMap[s[i]] == 0 {
                    delete(currentMap, s[i])
                }
            }
            i++
            if i == len(s) {
                return ans
            }
            for {
                if _, ok := mapT[s[i]]; ok {
                    break
                }
                i++
                if i == len(s) {
                    return ans
                }
            }
        } else {
            j++
            if j == len(s) {
                return ans
            }
            if _, ok := mapT[s[j]]; ok {
                if _, ok := currentMap[s[j]]; ok {
                    currentMap[s[j]]++
                } else {
                    currentMap[s[j]] = 1
                }
            }
        }
        fmt.Println(currentMap, i, j, ans)
    }
    //fmt.Println("hello")
    return ans
}
