
import (
    myheap "container/heap"
)

type MaxHeap []int

func (m *MaxHeap) Push(val interface{}) {
    *m = append(*m, val.(int))
}

func (m *MaxHeap) Pop() interface{} {
    n := len(*m)-1
    x := (*m)[n]
    *m = (*m)[:n]
    return x
}

func (m MaxHeap) Len() int {
    return len(m)
}

func (m MaxHeap) Swap(i, j int) {
    m[i], m[j] = m[j], m[i]
}

func (m MaxHeap) Less(i, j int) bool {
    return m[i] > m[j]
}

type QueueVal struct {
    freq int
    heapInsertTime int
}

func leastInterval(tasks []byte, n int) int {

    maxHeap := &MaxHeap{}
    count := make(map[byte]int)

    for _, val := range tasks {
        if _, ok := count[val]; ok {
            count[val]++
        } else {
            count[val] = 1
        }
    }

    for _, val := range count {
        myheap.Push(maxHeap, val)
    }

    time := 0
    queue := []QueueVal{}

    for len(queue) != 0 || maxHeap.Len() != 0 {
        time++
        if maxHeap.Len() != 0 {
            top := myheap.Pop(maxHeap)
            topInt := top.(int)
            if topInt - 1 != 0 {
                queue = append(queue, QueueVal{freq: topInt-1, heapInsertTime: time+n})
            }
        }

        if len(queue) != 0 {
            if queue[0].heapInsertTime == time {
                myheap.Push(maxHeap, queue[0].freq)
                queue = queue[1:]
            }
        }
        fmt.Println(queue, *maxHeap)
    }
    return time
}

