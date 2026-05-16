
func binSearch(nums []int, i int, j int) int {

    if i == j {
        return nums[i]
    }

    mid := int((i+j)/2)

    if nums[mid] <= nums[j] && nums[mid] >= nums[i] {
        return binSearch(nums, i, mid)
    } else if nums[mid] <= nums[j] && nums[mid] <= nums[i] {
        return binSearch(nums, i, mid)
    } else if nums[mid] >= nums[j] && nums[mid] >= nums[i] {
        return binSearch(nums, mid+1, j)
    } else {
        return binSearch(nums, mid+1, j)
    }

}


func findMin(nums []int) int {
    return binSearch(nums, 0, len(nums)-1)
}
