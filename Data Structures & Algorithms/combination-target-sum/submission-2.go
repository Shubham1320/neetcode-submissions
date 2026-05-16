var ans [][]int

func traverse(index int, subset []int, sum int, nums []int, target int) {

    if sum == target {
        fmt.Println(sum, target, subset)
        temp := make([]int, len(subset))
        copy(temp, subset)
        fmt.Println(ans)
        ans = append(ans, temp)
        fmt.Println(ans)
        return
    } else if index >= len(nums) || sum > target{
        return
    }

    
    subset = append(subset, nums[index])
    traverse(index, subset, sum+nums[index], nums, target)
 
    subset = subset[:len(subset)-1]
    traverse(index+1, subset, sum, nums, target)

}


func combinationSum(nums []int, target int) [][]int {
    ans = [][]int{}

    traverse(0, []int{}, 0, nums, target)

    return ans
}
