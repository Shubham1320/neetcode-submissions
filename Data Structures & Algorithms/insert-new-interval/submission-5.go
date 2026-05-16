func insert(intervals [][]int, newInterval []int) [][]int {

    ans := [][]int{}

    for i := 0; i < len(intervals); i++ {

        if newInterval[0] > intervals[i][1] {
            ans = append(ans, intervals[i])
        } else if newInterval[1] < intervals[i][0] {
            ans = append(ans, newInterval)
            return append(ans, intervals[i:]...)
        } else {
            newInterval[0] = min(intervals[i][0], newInterval[0])
            newInterval[1] = max(intervals[i][1], newInterval[1])
        }
    }
    ans = append(ans, newInterval)
    return ans
}
