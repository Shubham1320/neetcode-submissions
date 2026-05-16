func twoSum(nums []int, target int) []int {

    mapA := make(map[int]int)

    for index, value := range nums {
        findVal := target - value
        if index1, ok := mapA[findVal]; ok {
            return []int{index1, index}
        }
        mapA[value] = index
    }
    return []int{}
}
