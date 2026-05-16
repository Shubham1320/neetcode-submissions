
type Tuple struct {
    num1 int
    num2 int
    num3 int
}

func twoSum(nums []int, pos int , target int) ([][]int) {
    i := 0
    j := len(nums) - 1
    
    var ans [][]int

    for {
        if i >= j {
            break
        }
        
        if i == pos {
            i++
            continue
        }

        if j == pos {
            j--
            continue
        }

        sum := nums[i] + nums[j]
        if sum > target {
            j--
        } else if sum < target {
            i++
        } else {
            ans = append(ans, []int{-1 * target, nums[i], nums[j]})
            i++
            j--
        }
    }
    return ans
}

func threeSum(nums []int) [][]int {

    sort.Ints(nums)
    mapA := make(map[Tuple]bool)
    var ans [][]int

    for pos, num := range nums {
        pairs := twoSum(nums, pos, -1 * (num))
        fmt.Println(pairs)
        for _, pair := range pairs {
            res := []int{pair[0], pair[1], pair[2]}
            sort.Ints(res)
            key := Tuple{res[0], res[1], res[2]}
            fmt.Println(key, mapA)
            if _, ok := mapA[key]; !ok {
                ans = append(ans, res)
                mapA[key] = true
            }
        }
    }
    return ans

}
