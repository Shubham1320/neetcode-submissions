
type Node struct {
    children map[byte]*Node
    endOfWord bool
}


type PrefixTree struct {
    root *Node
}

func Constructor() PrefixTree {
    return PrefixTree{root: &Node{children: make(map[byte]*Node), endOfWord: false}}
}

func (this *PrefixTree) Insert(word string) {

    var insert func(*Node, int, string)

    insert = func(currentNode *Node, i int, word string) {

        if i == len(word) {
            currentNode.endOfWord = true
            return
        }

        if nextNode, ok := currentNode.children[byte(word[i])]; ok {
            insert(nextNode, i+1, word) 
        } else {
            newNode := &Node{children: make(map[byte]*Node), endOfWord: false}
            currentNode.children[byte(word[i])] = newNode
            insert(newNode, i+1, word)
        }
    }

    insert(this.root, 0, word)
}

func (this *PrefixTree) Search(word string) bool {

    var find func(*Node, int, string) bool

    find = func(currentNode *Node, i int, word string) bool {

        fmt.Println(i, currentNode.children, currentNode.endOfWord)

        if i == len(word) {
            fmt.Println("found")
            return currentNode.endOfWord
        }

        if nextNode, ok := currentNode.children[byte(word[i])]; ok {
            return find(nextNode, i+1, word)
        } 
        return false

    }

    return find(this.root, 0, word)

}

func (this *PrefixTree) StartsWith(prefix string) bool {
    var find func(*Node, int, string) bool

    find = func(currentNode *Node, i int, word string) bool {

        fmt.Println(i, currentNode.children, currentNode.endOfWord)

        if i == len(word) {
            return true
        }

        if nextNode, ok := currentNode.children[byte(word[i])]; ok {
            return find(nextNode, i+1, word)
        } 
        return false

    }

    return find(this.root, 0, prefix)
}
