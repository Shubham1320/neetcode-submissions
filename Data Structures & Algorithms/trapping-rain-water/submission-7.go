func trap(height []int) int {
    length := len(height)
    left := make([]int, length)
    right := make([]int, length)

    leftMax := height[0]
    rightMax := height[length - 1]

    for i := 1; i < len(height); i++ {
        left[i] = leftMax
        if height[i] > leftMax {
            leftMax = height[i]
        } 
    }

    for j := len(height) - 2; j >= 0; j-- {
        right[j] = rightMax
        if height[j] > rightMax {
            rightMax = height[j]
        }
    }

    fmt.Println(left, right)
    var ans int
    for pos, h := range height {
        minimum := min(left[pos], right[pos])
        if minimum != 0 && minimum - h > 0{
            ans = ans + (minimum - h)
        }
    }
    return ans
}
