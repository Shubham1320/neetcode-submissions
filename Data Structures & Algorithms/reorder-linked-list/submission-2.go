/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {

    if head.Next == nil {
        return
    }


    prevPtr := head
    slowPtr := head
    fastPtr := head

    for fastPtr != nil {
        prevPtr = slowPtr
        slowPtr = slowPtr.Next
        if fastPtr.Next != nil {
            fastPtr = fastPtr.Next.Next
        } else {
            fastPtr = fastPtr.Next
        }
    }

    //fmt.Println(prevPtr.Val)
    //fmt.Println(slowPtr.Val)

    prevPtr.Next = nil

    prevPtr = nil
    nextPtr := slowPtr.Next
    slowPtr.Next = nil

    for nextPtr != nil {
        slowPtr.Next = prevPtr
        prevPtr = slowPtr
        slowPtr = nextPtr
        nextPtr = slowPtr.Next
    }

    slowPtr.Next = prevPtr

    //fmt.Println(slowPtr.Val)

    dummyNode := &ListNode{}
    node := dummyNode

    ptr1 := head
    ptr2 := slowPtr

    fmt.Println(ptr1.Val)
    fmt.Println(ptr2.Val)

    for ptr1 != nil && ptr2 != nil {
        node.Next = ptr1
        ptr1 = ptr1.Next
        node = node.Next
        node.Next = ptr2
        ptr2 = ptr2.Next
        node = node.Next
    }

    node.Next = ptr1
    if ptr2 != nil {
        node.Next = ptr2
    }

    head = dummyNode.Next

}
