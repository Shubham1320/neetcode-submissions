var ans [][]int

func traverse(index int , subset []int, nums []int){

    if index >= len(nums) {
        temp := make([]int, len(subset))
        copy(temp, subset)
        ans = append(ans, temp)
        return
    }

    subset = append(subset, nums[index])
    traverse(index+1, subset, nums)

    for index + 1 < len(nums) && nums[index] == nums[index+1] {
        index++
    }
    subset = subset[:len(subset)-1]
    traverse(index+1, subset, nums)
}

func subsetsWithDup(nums []int) [][]int {
    ans = [][]int{}
    sort.Ints(nums)
    fmt.Println(nums)
    traverse(0, []int{}, nums)

    return ans
}
