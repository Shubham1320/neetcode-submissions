type Stack struct {
    head int
    data []int
    index []int
}

func (s *Stack) insert(index int, val int) {
    s.data = append(s.data, val)
    s.index = append(s.index, index)
    s.head++
}

func (s *Stack) top() (int, int) {
    return s.index[s.head], s.data[s.head]
}

func (s *Stack) pop() {
    s.data = s.data[0:s.head]
    s.index = s.index[0:s.head]
    s.head--
}

func dailyTemperatures(temperatures []int) []int {

    stk := &Stack{head: -1}

    ans := make([]int, len(temperatures))

    for j := len(temperatures) - 1; j >= 0; j-- {
        if stk.head == -1 {
            stk.insert(j, temperatures[j])
            ans[j] = 0
            continue
        }

        topIndex := -1
        topVal := -1

        for true{
        
            topIndex, topVal = stk.top()
            
            if temperatures[j] >= topVal {
                stk.pop()
                if stk.head == -1 {
                    ans[j] = 0
                    topIndex = - 1
                    topVal = - 1
                    break
                }
            } else {
                break
            }
        }

        stk.insert(j, temperatures[j])

        if topIndex != -1 {
            ans[j] = topIndex - j
        }

        fmt.Println(stk.data)
        fmt.Println(stk.index)
        fmt.Println(ans)
        fmt.Println("\n")

    }

    return ans

}
