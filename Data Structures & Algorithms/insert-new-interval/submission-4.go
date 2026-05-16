func insert(intervals [][]int, newInterval []int) [][]int {

    i := 0
    for ; i < len(intervals); i++ {

        start := intervals[i][0]

        if newInterval[0] <= start {
            break
        }
    }

    fmt.Println(i)

    intervalsToRemove := []int{}

    j := i
    for ; j < len(intervals); j++ {
        if newInterval[1] < intervals[j][0] {
            break
        }
        intervalsToRemove = append(intervalsToRemove, j)
    }

    n := len(intervalsToRemove)
    if n > 0 {
        if newInterval[1] < intervals[intervalsToRemove[n-1]][1] {
            newInterval[1] = intervals[intervalsToRemove[n-1]][1]
        }
        intervals = append(intervals[0:intervalsToRemove[0]], intervals[intervalsToRemove[n-1]+1:]...)
    }

    fmt.Println(i, intervals, newInterval)

    if i > 0 {
        if newInterval[0] <= intervals[i-1][1] {
            newInterval[0] = intervals[i-1][0]
            if newInterval[1] < intervals[i-1][1] {
                newInterval[1] = intervals[i-1][1]
            }
            intervals = append(intervals[:i-1], intervals[i:]...)
            i--
        }
    }
    fmt.Println(i, intervals, newInterval)

    fmt.Println(intervalsToRemove)
    fmt.Println(newInterval)
    fmt.Println(intervals)

    if i == 0 {
        intervals = append([][]int{newInterval}, intervals...)
    } else if i == len(intervals) {
        intervals = append(intervals, newInterval)
    } else {
        intervals = append(intervals[0:i], append([][]int{newInterval}, intervals[i:]...)...)
    }

    return intervals

}
