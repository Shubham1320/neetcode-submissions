
var ans [][]int

func recursive(index int, subset []int, nums []int) {

    if index >= len(nums) {
        temp := make([]int, len(subset))
        copy(temp, subset)
        ans = append(ans, temp)
        fmt.Println(subset)
        return
    }
    subset = append(subset, nums[index])
    recursive(index+1, subset, nums)
    subset = subset[:len(subset)-1]
    recursive(index+1, subset, nums)
}

func subsets(nums []int) [][]int {
    ans = [][]int{}
    
    recursive(0, []int{}, nums)

    return ans
}
