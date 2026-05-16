/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {

    nodeMap := make(map[*Node]*Node)

    var newHead *Node
    var prevNode *Node

    for head != nil {

        var newNode *Node

        newNode, ok := nodeMap[head]

        //fmt.Println(head.Val)

        if !ok {
            newNode = &Node{
                Val: head.Val,
            }
            nodeMap[head] = newNode
        }

        //fmt.Println(newNode.Val)

        if newHead == nil {
            newHead = newNode
        }

        //fmt.Println(newHead.Val)

        if prevNode != nil {
            prevNode.Next = newNode
        }

        if head.Random != nil {
            newRandomNode, ok := nodeMap[head.Random]
            if !ok {
                newRandomNode = &Node {
                    Val: head.Random.Val,
                }
                nodeMap[head.Random] = newRandomNode
            }
            newNode.Random = newRandomNode
        }

        prevNode = newNode
        head = head.Next

    }

    return newHead

}
