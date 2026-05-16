type Stack struct {
    data []int
    head int
}

func (s *Stack) insert(val int) {
    s.data = append(s.data, val)
    s.head++
}

func (s *Stack) top() int{
    return s.data[s.head]
}

func (s *Stack) pop() {
    s.data = s.data[0:s.head]
    s.head--
}

func evalRPN(tokens []string) int {

    stk := &Stack{head: -1}
    for _, tok := range tokens {

        if val, err := strconv.Atoi(tok); err == nil {
            stk.insert(val)
            continue
        }

        num2 := stk.top()
        stk.pop()
        num1 := stk.top()
        stk.pop()

        if tok == "+" {
            stk.insert(num1+num2)
        } else if tok == "-" {
            stk.insert(num1-num2)
        } else if tok == "*" {
            stk.insert(num1*num2)
        } else {
            stk.insert(num1/num2)
        }

    }
    return stk.data[stk.head]
}
