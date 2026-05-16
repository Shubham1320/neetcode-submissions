
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
    traverse(index+1, subset, sum+nums[index], nums, target)
 
    subset = subset[:len(subset)-1]

    for index+1 < len(nums) && nums[index] == nums[index+1] {
        index++
    }
    traverse(index+1, subset, sum, nums, target)

}


func combinationSum2(candidates []int, target int) [][]int {
    ans = [][]int{}

    sort.Ints(candidates)

    traverse(0, []int{}, 0, candidates, target)

    return ans
}
