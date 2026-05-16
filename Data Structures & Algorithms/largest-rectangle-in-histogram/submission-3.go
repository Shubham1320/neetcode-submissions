type Data struct {
    height int
    idx int
}

type Stack struct {
    data []Data
    head int
}

func (s *Stack) insert(data Data) {
    s.data = append(s.data, data)
    s.head++
}

func (s *Stack) pop() {
    s.data = s.data[:s.head]
    s.head--
}

func (s *Stack) top() Data {
    return s.data[s.head]
}

func largestRectangleArea(heights []int) int {

    s := Stack{head: -1}
    ans := 0

    s.insert(Data{height: heights[0], idx: 0})
    ans = heights[0]


    for i := 1; i < len(heights); i++ {
        stkTop := s.top()
        if heights[i] >= stkTop.height {
            s.insert(Data{height: heights[i], idx: i})
        } else {
            for {
                area := (i - stkTop.idx) * stkTop.height                         
                if area > ans {
                    ans = area
                }
                index := stkTop.idx
                s.pop()
                if s.head != -1 {
                    stkTop = s.top()
                    if heights[i] >= stkTop.height {
                        s.insert(Data{height: heights[i], idx: index})
                        break
                    } 
                } else {
                    s.insert(Data{height: heights[i], idx: index})
                    break
                }
            }
        }
        fmt.Println(s.data)
    }

    for {
        if s.head != -1 {
            stkTop := s.top()
            area :=  (len(heights) - stkTop.idx) * stkTop.height
            if area > ans {
                ans = area
            }
            s.pop()
        } else {
            break
        }
    }

    return ans
}




 


