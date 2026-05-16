/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */


func cloneGraph(node *Node) *Node {

    var dfs func(*Node)*Node

    oldToNew := make(map[*Node]*Node)

    dfs = func(node *Node) *Node {

        if node == nil {
            return nil
        }

        if val, ok := oldToNew[node]; ok {
            return val
        }

        newNode := &Node{
            Val: node.Val,
            Neighbors: []*Node{},
        }

        oldToNew[node] = newNode

        for _, val := range node.Neighbors {
            newNode.Neighbors = append(newNode.Neighbors, dfs(val))
        }

        return newNode
    }

    return dfs(node)
}
