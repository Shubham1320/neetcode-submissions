func binSearch(nums []int, i int, j int, target int) int {

    if i > j {
        return -1
    }

    mid := int((i+j)/2)

    if nums[mid] == target {
        return mid
    }

    if nums[i] <= nums[mid] {
        if target > nums[mid] || target < nums[i] {
            return binSearch(nums, mid+1 , j, target)
        } else {
            return binSearch(nums, i, mid-1, target)
        }
    } else {
        if target < nums[mid] || target > nums[j] {
            return binSearch(nums, i, mid-1, target)
        } else {
            return binSearch(nums, mid+1 , j, target)
        }
    }
}


func search(nums []int, target int) int {
    return binSearch(nums, 0, len(nums)-1, target)
}
