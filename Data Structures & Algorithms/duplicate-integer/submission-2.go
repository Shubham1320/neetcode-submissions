func hasDuplicate(nums []int) bool {

    map1 := make(map[int]bool)
    for _, val := range nums {
        if _, ok := map1[val]; ok {
            return true
        }
        map1[val] = true
    }
    return false
}




