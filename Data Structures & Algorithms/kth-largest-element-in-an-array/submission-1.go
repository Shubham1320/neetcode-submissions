import (
    myheap "container/heap"
)

type MinHeap []int

func (m *MinHeap) Push(val interface{}) {
    *m = append(*m, val.(int))
}

func (m *MinHeap) Pop() interface{} {
    n := len(*m)-1
    x := (*m)[n]
    *m = (*m)[:n]
    return x
}

func (m MinHeap) Len() int {
    return len(m)
}

func (m MinHeap) Swap(i,j int) {
    m[i], m[j] = m[j], m[i]
}

func (m MinHeap) Less(i,j int) bool {
    return m[i] < m[j]
}

func findKthLargest(nums []int, k int) int {

    minHeap := &MinHeap{}

    for _, val := range nums {

        if minHeap.Len() == k {
            if val > (*minHeap)[0] {
                myheap.Pop(minHeap)
            } else {
                continue
            }
        } 

        myheap.Push(minHeap, val)
    }
    return (*minHeap)[0]

}
