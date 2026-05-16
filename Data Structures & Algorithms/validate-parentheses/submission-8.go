type Stack struct {
    Head int
    Data []rune
}

func (s *Stack) insert(val rune) {
    s.Data = append(s.Data, val)
    s.Head++
}

func (s *Stack) top() rune {
    return s.Data[s.Head]
}

func (s *Stack) pop() {
    s.Data = s.Data[0:s.Head]
    s.Head--
}

func isValid(s string) bool {

    stk := &Stack{Head: -1}

    if len(s) <= 1 {
        return false
    }

    for _, char := range s {

        if char == '[' || char == '{' || char == '(' {
            stk.insert(char)
        } else {

            if stk.Head == - 1{
                return false
            }

            if (stk.top() == '[' && char == ']') || (stk.top() == '(' && char == ')') || (stk.top() == '{' && char == '}') {
                stk.pop()
            } else {
                return false
            }
        }
    }

    if stk.Head != -1 {
        return false
    }

    return true
}
