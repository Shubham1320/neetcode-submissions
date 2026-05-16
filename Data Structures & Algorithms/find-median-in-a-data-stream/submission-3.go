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

func (m MinHeap) Swap(i, j int) {
    m[i], m[j] = m[j], m[i]
}

func (m MinHeap) Less(i, j int) bool {
    return m[i] < m[j]
}

func (m MinHeap) Len() int {
    return len(m)
}

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

func (m MaxHeap) Swap(i, j int) {
    m[i], m[j] = m[j], m[i]
}

func (m MaxHeap) Less(i, j int) bool {
    return m[i] > m[j]
}

func (m MaxHeap) Len() int {
    return len(m)
}


type MedianFinder struct {
    maxHeap *MaxHeap
    minHeap *MinHeap
}


func Constructor() MedianFinder {
    return MedianFinder {
        maxHeap: &MaxHeap{},
        minHeap: &MinHeap{},
    }
}


func (this *MedianFinder) AddNum(num int)  {

    if this.maxHeap.Len() == 0 {
        myheap.Push(this.maxHeap, num)
    } else if num < (*this.maxHeap)[0] {
        myheap.Push(this.maxHeap, num)
    } else {
        myheap.Push(this.minHeap, num)
    }

    fmt.Println(this.maxHeap, this.minHeap)

    if this.maxHeap.Len() == this.minHeap.Len() {
        return
    } else {

        for this.maxHeap.Len() != this.minHeap.Len() + 1 {

            if this.maxHeap.Len() > this.minHeap.Len() {
                val := myheap.Pop(this.maxHeap).(int)
                myheap.Push(this.minHeap, val)
            } else if this.maxHeap.Len() < this.minHeap.Len() {
                val := myheap.Pop(this.minHeap).(int)
                myheap.Push(this.maxHeap, val)
            }

            if this.maxHeap.Len() == this.minHeap.Len() {
                break
            }
        }

    }
    fmt.Println(this.maxHeap, this.minHeap)
}


func (this *MedianFinder) FindMedian() float64 {
    if this.minHeap.Len() == this.maxHeap.Len() {
        return (float64((*this.minHeap)[0] + (*this.maxHeap)[0]))/2
    } else {
        return float64((*this.maxHeap)[0])
    }
}
