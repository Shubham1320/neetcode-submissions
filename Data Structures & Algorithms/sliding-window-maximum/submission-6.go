type Node struct {
    Data int
    Index int
}

var maxHeap []*Node

func heapify(curIndex int) {
    i := curIndex
    for true {
        if i == 0 {
            break
        }

        var root int
        if i % 2 == 0 {
            root = i/2 - 1
        } else {
            root = int(i/2)
        }
        
        if maxHeap[root].Data >= maxHeap[i].Data {
            break
        } else {
            temp := maxHeap[i]
            maxHeap[i] = maxHeap[root]
            maxHeap[root] = temp
            i = root
        }
    }
}

func insert(num int, index int) { 
    newNode := &Node{Data: num, Index: index}
    if maxHeap[0] == nil {
        maxHeap[0] = newNode
        return
    }

    i := 0
    for true {
        if maxHeap[2*i + 1] == nil {
            i = 2*i + 1
            break
        } else if maxHeap[2*i + 2] == nil {
            i = 2*i + 2
            break
        } else {
            i++
        }
    }
    maxHeap[i] = newNode
    heapify(i)
}

func pop() (int, int) {
    ansData := maxHeap[0].Data
    ansIndex := maxHeap[1].Index
    if maxHeap[1] == nil {
        return ansIndex, ansData
    }

    for index, node := range maxHeap {
        if node == nil {
            maxHeap[0] = maxHeap[index-1]
            maxHeap[index-1] = nil
            break
        }
    }

    i := 0
    j := 0
    for true {
        var max int
        j = i
        if maxHeap[2*i + 1] != nil {
            max = maxHeap[2*i + 1].Data
            j = 2*i + 1
        }

        if maxHeap[2*i + 2] != nil {
            if maxHeap[2*i + 2].Data > max {
                max = maxHeap[2*i + 2].Data
                j = 2*i + 2
            }
        }

        if i == j {
            break
        }

        if maxHeap[i].Data >= max {
            break
        }

        temp := maxHeap[i]
        maxHeap[i] = maxHeap[j]
        maxHeap[j] = temp
        i = j
    }
    return ansIndex, ansData
}

func display() {
    var ans []int
    var index []int
    for _, node := range maxHeap {
        if node != nil {
            ans = append(ans, node.Data)
            index = append(index, node.Index)
        }
    }
    fmt.Println(ans)
    fmt.Println(index)
    fmt.Println("\n")
}

func maxSlidingWindow(nums []int, k int) []int {

    maxHeap = make([]*Node, 1000)
    var ans []int

    for i := 0; i < k; i++ {
        insert(nums[i], i)
    }


    for i,j := 0,k; j <= len(nums); i,j = i+1, j+1{
        fmt.Println(i,j)
        for true {
            index, data := maxHeap[0].Index, maxHeap[0].Data
            if index >= i {
                ans = append(ans, data)
                break
            }
            pop()
        }
        if j < len(nums){
            insert(nums[j], j)
        }
        display()
    }

    return ans
}
