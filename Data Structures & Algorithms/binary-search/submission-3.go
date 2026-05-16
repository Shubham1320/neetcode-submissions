
func binarySearch(nums []int, i int, j int, target int) int {

    if i > j {
        return -1
    }

    mid := int((i + j) / 2)
    fmt.Println(i, j, nums[mid])

    if nums[mid] == target {
        return mid
    } else if target < nums[mid] {
        return binarySearch(nums, i, mid-1, target)
    } else {
        return binarySearch(nums, mid+1, j, target)
    }

}

func search(nums []int, target int) int {
    return binarySearch(nums, 0, len(nums)-1, target)
}
