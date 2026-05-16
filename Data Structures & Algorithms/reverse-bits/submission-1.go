func reverseBits(n int) int {

    n1 := uint32(n)
    fmt.Println(n1)

    var ans uint32
    ans = 0

    for i := 0; i < 32; i++ {
        bit := (n1 & 1) << (31-i)
        n1 >>= 1
        ans |= bit
    }
    return int(ans)

}



//000101 & 000001 = 000001 << 5 = 100000

//100010 & 000001 = 000000 << 4 = 000000

//100001 & 000001 = 000001 << 3 = 001000



