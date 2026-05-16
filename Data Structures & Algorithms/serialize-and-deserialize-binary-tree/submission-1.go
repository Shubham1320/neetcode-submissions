/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct{}

func Constructor() Codec {
    return Codec{}
}

// Encodes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    if root == nil {
        return "N"
    }
    var res []string
    queue := []*TreeNode{root}

    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]

        if node == nil {
            res = append(res, "N")
        } else {
            res = append(res, strconv.Itoa(node.Val))
            queue = append(queue, node.Left)
            queue = append(queue, node.Right)
        }
    }

    return strings.Join(res, ",")
}

// Decodes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
    vals := strings.Split(data, ",")
    if vals[0] == "N" {
        return nil
    }

    rootVal, _ := strconv.Atoi(vals[0])
    root := &TreeNode{Val: rootVal}
    queue := []*TreeNode{root}
    index := 1

    for len(queue) > 0 && index < len(vals) {
        node := queue[0]
        queue = queue[1:]

        if vals[index] != "N" {
            leftVal, _ := strconv.Atoi(vals[index])
            node.Left = &TreeNode{Val: leftVal}
            queue = append(queue, node.Left)
        }
        index++

        if index < len(vals) && vals[index] != "N" {
            rightVal, _ := strconv.Atoi(vals[index])
            node.Right = &TreeNode{Val: rightVal}
            queue = append(queue, node.Right)
        }
        index++
    }

    return root
}
