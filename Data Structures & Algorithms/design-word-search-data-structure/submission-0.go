type Node struct {
    children map[byte]*Node
    endOfWord bool
}

type WordDictionary struct {
    root *Node
}

func Constructor() WordDictionary {
    return WordDictionary {
        root: &Node{children: make(map[byte]*Node)},
    }
}

func (this *WordDictionary) AddWord(word string)  {
    var insert func(*Node, int, string)

    insert = func(node *Node, i int, word string) {

        if i == len(word) {
            node.endOfWord = true
            return
        }

        if _, ok := node.children[byte(word[i])]; ok {
            insert(node.children[byte(word[i])], i+1, word)    
        } else {
            newNode := &Node{children: make(map[byte]*Node)}
            node.children[byte(word[i])] = newNode
            insert(newNode, i+1, word)
        }
    }

    insert(this.root, 0, word)
}

func (this *WordDictionary) Search(word string) bool {

    var find func(*Node, int, string) bool

    find = func(node *Node, i int, word string) bool {

        if i == len(word) {
            return node.endOfWord
        }

        if word[i] == byte('.') {
            ans := false
            for _, val := range node.children {
                ans = ans || find(val, i+1, word)
                if ans == true {
                    break
                }
            }
            return ans
        } else if _, ok := node.children[byte(word[i])]; ok {
            return find(node.children[byte(word[i])], i+1, word)
        }
        return false
    }    

    return find(this.root, 0, word)
}
