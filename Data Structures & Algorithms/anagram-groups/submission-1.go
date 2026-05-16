
func generateKey(str string) [26]int {

    var key [26]int

    for _, value := range str {
        key[value-'a']++
    }
    return key
}

func groupAnagrams(strs []string) [][]string {

    mapA := make(map[[26]int][]string)

    for _, str := range strs {
        fmt.Println(str)
        signatureMap := generateKey(str)
        fmt.Println(signatureMap)
        if _, ok := mapA[signatureMap]; ok {
            mapA[signatureMap] = append(mapA[signatureMap], str)
        } else {
            mapA[signatureMap] = []string{str}
        }
    }

    ans := [][]string{}
    for _, val := range mapA {
        ans = append(ans, val)
    }
    return ans

}
