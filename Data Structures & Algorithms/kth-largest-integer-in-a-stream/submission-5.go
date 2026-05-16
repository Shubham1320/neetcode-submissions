type MinHeap struct {
    Nums []int
    Capacity int
    CurrentSize int
}

func (heap *MinHeap) insert(val int) {
    if heap.CurrentSize == heap.Capacity {
        if val > heap.Nums[0] {
            heap.pop()
        } else {
            return
        }
    }

    heap.Nums = append(heap.Nums, val)
    index := heap.CurrentSize

    for index > 0{

        var parent int
        if index % 2 == 0 {
            parent = (index-2)/2
        } else {
            parent = (index-1)/2
        }
 
        if heap.Nums[parent] <= heap.Nums[index] {
            break
        }

        temp := heap.Nums[parent]
        heap.Nums[parent] = heap.Nums[index]
        heap.Nums[index] = temp
        index = parent
    }

    heap.CurrentSize++
    fmt.Println(heap.Nums, heap.CurrentSize, heap.Capacity)
}


func (heap *MinHeap) pop() int {
    
    ans := heap.Nums[0]

    heap.Nums[0] = heap.Nums[heap.CurrentSize-1]
    heap.Nums = heap.Nums[0:heap.CurrentSize-1]
    heap.CurrentSize--

    index := 0
    for {
        minIndex := 2*index+1
        if 2*index+1 <= heap.CurrentSize-1 && 2*index+2 <= heap.CurrentSize-1{
            if heap.Nums[2*index+2] < heap.Nums[2*index+1] {
                minIndex = 2*index + 2
            }
        } else if 2*index+1 <= heap.CurrentSize-1 {
            minIndex = 2*index+1
        } else {
            break
        }

        fmt.Println("min index", heap.Nums, minIndex)

        if heap.Nums[index] <= heap.Nums[minIndex] {
            break
        }

        temp := heap.Nums[index]
        heap.Nums[index] = heap.Nums[minIndex]
        heap.Nums[minIndex] = temp 
        index = minIndex
    } 

    fmt.Println(heap.Nums, heap.CurrentSize, heap.Capacity)
    return ans
}


type KthLargest struct {
    heap *MinHeap  
}


func Constructor(k int, nums []int) KthLargest {

    kthLargest := KthLargest{}

    heap := &MinHeap {
        Capacity: k,
        CurrentSize: 0,
    }

    for _, val := range nums {
        heap.insert(val)
    }

    kthLargest.heap = heap
    return kthLargest
}


func (this *KthLargest) Add(val int) int {
    this.heap.insert(val)
    return this.heap.Nums[0]
}
