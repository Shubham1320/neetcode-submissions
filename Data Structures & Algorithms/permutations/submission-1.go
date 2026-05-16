import (
    "slices"
)

var ans [][]int

func traverse(nums []int, subset []int) {

    if len(subset) == len(nums) {
        for index, val := range subset {
            subset[index] = nums[val] 
        }

        ans = append(ans, subset)
        return
    }

    for index, _ := range nums {
        if slices.Contains(subset, index) {
            continue
        }

        subsetTemp := make([]int, len(subset))
        copy(subsetTemp, subset)
        subsetTemp = append(subsetTemp, index)

        traverse(nums, subsetTemp)

    }
}

func permute(nums []int) [][]int {
    ans = [][]int{}

    for index, _ := range nums {
        traverse(nums, []int{index})
    }

    return ans
}
