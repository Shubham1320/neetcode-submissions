


func maxProduct(nums []int) int {

    res := math.MinInt
    currMin := 1
    currMax := 1

    for i := 0; i < len(nums); i++ {

        if nums[i] == 0 {
            currMin = 1
            currMax = 1
        }

        temp := currMax * nums[i]
        currMax = max(currMax*nums[i], max(nums[i], currMin*nums[i]))
        currMin = min(temp, min(nums[i], currMin*nums[i]))
        res = max(res, currMax)
    }
    return res

}

