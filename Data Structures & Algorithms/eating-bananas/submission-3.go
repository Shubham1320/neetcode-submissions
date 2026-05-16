


func hoursNeeded(piles []int, k int) int {

    hrs := 0

    for _, val := range piles {
        hrs = hrs + int(math.Ceil(float64(val)/float64(k)))
    }

    return hrs
}


func binSearch(piles []int, i int, j int, h int) int {


    if i > j {
        return -1
    }

    mid := int((i+j)/2)

    noOfHours := hoursNeeded(piles, mid)

    fmt.Println(i, j, mid, noOfHours)

    if noOfHours > h {
        return binSearch(piles, mid+1, j, h)
    } else if noOfHours <= h {
        ans := binSearch(piles, i, mid-1, h)
        if ans == -1 {
            return mid
        }
        return ans
    }
    return -1
}


func minEatingSpeed(piles []int, h int) int {

    max := 0
    for _, val := range piles {
        if val > max {
            max = val
        }
    }

    return binSearch(piles, 1, max, h)
}




