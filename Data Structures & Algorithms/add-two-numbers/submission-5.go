/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

    carryforward := 0
    var newHead *ListNode
    var prevNode *ListNode

    for l1 != nil && l2 != nil {

        newNode := &ListNode {
            Val: int((l1.Val + l2.Val + carryforward)%10),
        }

        carryforward = int((l1.Val + l2.Val + carryforward) / 10)
        if newHead == nil {
            newHead = newNode
        }

        if prevNode != nil {
            prevNode.Next = newNode
        }        
        prevNode = newNode
        l1 = l1.Next
        l2 = l2.Next
    }

    ptr := l1
    if l1 == nil {
        ptr = l2
    }

    for ptr != nil {

        fmt.Println(ptr.Val, carryforward)
        newNode := &ListNode {
            Val: int((ptr.Val + carryforward)%10),
        }

        carryforward = int((ptr.Val + carryforward)/ 10)
        if newHead == nil {
            newHead = newNode
        }

        if prevNode != nil {
            prevNode.Next = newNode
        } 

        prevNode = newNode
        ptr = ptr.Next
    }

    if carryforward > 0 {
        newNode := &ListNode {
            Val:  carryforward,
        }
        prevNode.Next = newNode
    }

    return newHead

}
