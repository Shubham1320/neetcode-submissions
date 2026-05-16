func productExceptSelf(nums []int) []int {
    length := len(nums)
    prefix := make([]int, length) 
    suffix := make([]int, length)

    prefix[0] = nums[0]
    for i := 1; i < len(nums); i++ {
        prefix[i] = nums[i] * prefix[i-1]
    }

    suffix[len(nums)-1] = nums[len(nums)-1] 
    for j := len(nums)-2; j >= 0; j-- {
        suffix[j] = nums[j] * suffix[j+1]
    } 

    fmt.Println(prefix, suffix)

    ans := make([]int, length)
    ans[0] = suffix[1]
    ans[length-1] = prefix[length-2]

    for i := 1; i < length-1; i++ {
        ans[i] = prefix[i-1] * suffix[i+1]
    }

    return ans

}
