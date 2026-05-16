func getSum(a int, b int) int {
    mask := 0xFFFFFFFF
    maxInt := 0x7FFFFFFF

    for b != 0 {
        carry := (a & b) << 1
        a = (a ^ b) & mask
        b = carry & mask
    }

    if a <= maxInt {
        return a
    }
    return ^(a ^ mask)
}