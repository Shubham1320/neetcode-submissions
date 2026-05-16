/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverse(head *ListNode) *ListNode {

    if head == nil {
        return nil
    }

    newHead := head
    if head.Next != nil {
        newHead = reverse(head.Next)
        head.Next.Next = head
    }
    head.Next = nil
    return newHead
}


func reverseList(head *ListNode) *ListNode {
    
    if head == nil {
        return nil
    }

    ans := reverse(head)
    return ans
}
