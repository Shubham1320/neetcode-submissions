/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {

    slowPtr := head
    fastPtr := head

    for fastPtr != nil && slowPtr != nil {

        if fastPtr.Next != nil {
            slowPtr = slowPtr.Next
            fastPtr = fastPtr.Next.Next

            if slowPtr == fastPtr {
                return true
            }
        } else {
            break
        }
    }

    return false

}
