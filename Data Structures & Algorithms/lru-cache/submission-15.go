type Node struct {
    Key int
    Val int
    Prev *Node
    Next *Node
}

type LRUCache struct {
    head *Node
    last *Node
    capacity int
    noOfNodes int
    hashMap map[int]*Node
}

func Constructor(capacity int) LRUCache {
    hashMap := make(map[int]*Node)
    return LRUCache{capacity: capacity, hashMap: hashMap}
}

func (this *LRUCache) Get(key int) int {

    if node,ok := this.hashMap[key]; ok {

        if this.last != node {
            if this.head == node {
                this.head = node.Next
                node.Next.Prev = nil
            } else {
                node.Prev.Next = node.Next
                node.Next.Prev = node.Prev
            }

            this.last.Next = node
            node.Prev = this.last
            node.Next = nil
            this.last = node
        }

        fmt.Println("GET", this.head, this.last, this.hashMap)

        return node.Val
    }
    return -1
}

func (this *LRUCache) Put(key int, value int) {

    if node, ok := this.hashMap[key]; ok {
        node.Val = value

        if this.last != node {
            if this.head == node {
                this.head = node.Next
                node.Next.Prev = nil
            } else {
                node.Prev.Next = node.Next
                node.Next.Prev = node.Prev
            }

            this.last.Next = node
            node.Prev = this.last
            node.Next = nil
            this.last = node
        }

    } else {
        node = &Node{
            Key: key,
            Val: value,
        }

        if this.head == nil {
            this.head = node
            this.last = node
        } else {
            this.last.Next = node
            node.Prev = this.last
            this.last = node
        }

        if this.noOfNodes < this.capacity {
            this.noOfNodes++
        } else {
            delete(this.hashMap, this.head.Key)
            this.head.Next.Prev = nil
            this.head = this.head.Next
        }
        this.hashMap[key] = node
    }
    fmt.Println("PUT", this.head, this.last, this.hashMap) 
}
