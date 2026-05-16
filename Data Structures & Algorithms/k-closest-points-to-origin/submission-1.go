
import (
    myheap "container/heap"
)

type Point struct {
    x int
    y int
    distanceFromOrg float64
}

type MinHeap []Point

func (m *MinHeap) Push(val interface{}) {
    *m = append(*m, val.(Point))
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
    return m[i].distanceFromOrg < m[j].distanceFromOrg
}

func kClosest(points [][]int, k int) [][]int {

    minHeap := &MinHeap{}
    var ans [][]int

    for _, val := range points {
        x := val[0]
        y := val[1]

        temp := math.Pow(float64(x),2.0) + math.Pow(float64(y),2.0)
        distance := math.Pow(temp, 0.5)

        point := Point {
            x: x,
            y: y,
            distanceFromOrg: distance,
        }
        fmt.Println(point)
        myheap.Push(minHeap, point)
        fmt.Println(minHeap)
    }

    for k != 0 {
        
        temp := myheap.Pop(minHeap)
        point := temp.(Point)
        ans = append(ans, []int{point.x, point.y})
        k--
    }

    return ans
}
