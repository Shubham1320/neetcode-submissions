func hammingWeight(n int) int {

    result := 0

    for n != 0 {
        bit := (n & 1)
        if bit == 1 {
            result++
        }
        n >>=1 
    }
    return result

}
