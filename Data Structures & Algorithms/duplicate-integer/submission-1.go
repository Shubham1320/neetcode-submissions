func hasDuplicate(nums []int) bool {

    mapA := make(map[int]bool)
    for _, value := range nums {
        fmt.Println(value)
        if _, ok := mapA[value]; ok {
            return true
        }
        mapA[value] = true
    }
    return false
}
