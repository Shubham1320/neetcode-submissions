func twoSum(numbers []int, target int) []int {

    i := 0
    j := len(numbers) - 1

    for {
        sum := numbers[i] + numbers[j]
        if sum > target {
            j--
        } else if sum < target {
            i++
        } else {
            return []int{i+1, j+1}
        }
    }

}
