import (
    myheap "container/heap"
)

type MaxHeap []int

func (h *MaxHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
    
    n := len(*h) - 1
    x := (*h)[n]
    *h = (*h)[0:n]
    return x
}

func (h MaxHeap) Swap(i int, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h MaxHeap) Less(i, j int) bool {
    return h[i] > h[j]
}

func (h MaxHeap) Len() int {
    return len(h)
}

func lastStoneWeight(stones []int) int {
    maxHeap := &MaxHeap{}
    *maxHeap = stones
    myheap.Init(maxHeap)

    for maxHeap.Len() > 1 {

        firstMax := myheap.Pop(maxHeap)
        secondMax := myheap.Pop(maxHeap)

        first := firstMax.(int)
        second := secondMax.(int)

        if first > second {
            myheap.Push(maxHeap, first-second)    
        } else if second > first {
            myheap.Push(maxHeap, second-first)
        }
    }

    fmt.Println(*maxHeap)

    if maxHeap.Len() == 1 {
        return (*maxHeap)[0]
    }

    return 0
}
