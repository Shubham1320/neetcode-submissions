type MinStack struct {
    min []int
    data []int
    head int
}

func Constructor() MinStack {
    return MinStack{head: -1}
}

func (this *MinStack) Push(val int) {
    this.data = append(this.data, val)
    
    if this.head == - 1 {
        this.min = append(this.min, val)
    } else if val < this.min[this.head] {
        this.min = append(this.min, val)
    } else {
        this.min = append(this.min, this.min[this.head])
    }
    this.head++
}

func (this *MinStack) Pop() {

    this.min = this.min[0:this.head]
    this.data = this.data[0:this.head]
    this.head--
}

func (this *MinStack) Top() int {
    return this.data[this.head]
}

func (this *MinStack) GetMin() int {
    return this.min[this.head]
}
