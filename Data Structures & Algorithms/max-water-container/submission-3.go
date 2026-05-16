func maxArea(heights []int) int {

    i := 0
    j := len(heights) - 1
    max := 0

    for {
        if i >= j {
            break
        }
        area := min(heights[i],heights[j]) * (j-i)
        if area > max {
            max = area
        }

        if heights[i] < heights[j] {
            i++
        } else {
            j--
        }
    }

    return max
}
