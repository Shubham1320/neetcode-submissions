import (
    myheap "container/heap"
)

type MinHeap []Edge

func (m MinHeap) Len() int {
    return len(m)
}

func (m MinHeap) Swap(i int, j int) {
    m[i], m[j] = m[j], m[i]
}

func (m MinHeap) Less (i int, j int) bool {
    return m[i].Time < m[j].Time
}

func (m *MinHeap) Push(val interface{}) {
    *m = append(*m, val.(Edge))
}

func (m *MinHeap) Pop() interface{} {
    n := len(*m) - 1
    x := (*m)[n]
    *m = (*m)[:n]
    return x
}

type Edge struct {
    Time int
    Vi int
}

func networkDelayTime(times [][]int, n int, k int) int {

    minHeap := &MinHeap{}

    adjList := make(map[int][]Edge)

    for i := 0; i < len(times); i++ {
        ui, vi, time := times[i][0], times[i][1], times[i][2]
        if _, ok := adjList[ui]; ok {
            adjList[ui] = append(adjList[ui], Edge{Time: time, Vi: vi})
        } else {
            adjList[ui] = []Edge{Edge{Time: time, Vi: vi}}
        }
    }

    myheap.Push(minHeap, Edge{0, k})
    visited := make(map[int]bool)

    time := 0

    for minHeap.Len() != 0 {

        top := myheap.Pop(minHeap)
        ui := top.(Edge).Vi

        if _, ok := visited[ui]; ok {
            continue
        }
        time = top.(Edge).Time
        visited[ui] = true

        for _, val := range adjList[ui] {
            vi := val.Vi
            curTime := val.Time
            if _, ok := visited[vi]; ok {
                continue
            }
            myheap.Push(minHeap, Edge{Time: curTime+time, Vi: vi})
        }
        fmt.Println(minHeap)
    }

    fmt.Println(visited)
    fmt.Println(adjList)
    if n == len(visited) {
        return time
    } 
    return -1
}
