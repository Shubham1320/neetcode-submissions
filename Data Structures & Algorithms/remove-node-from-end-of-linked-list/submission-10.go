/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {

    fastptr := head

    for i := 1; i < n; i++ {
        fastptr = fastptr.Next
    }

    if fastptr == head && head.Next == nil {
        return nil
    }

    fmt.Println(fastptr.Val)

    ptr := head
    prev := head

    for fastptr.Next != nil {
        prev = ptr
        ptr = ptr.Next
        fastptr = fastptr.Next
    }

    if prev == ptr {
        return prev.Next
    }

    prev.Next = ptr.Next

    return head

}
