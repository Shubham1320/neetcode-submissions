# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:
        
        i = head
        j = head
        if head.next is None:
            return None

        k = 0
        while k < n:
            j = j.next
            if j is None:
                head = i.next
                return head
            k += 1

        while j and j.next is not None:
            print(i.val, j.val)
            i = i.next
            j = j.next

        print(i.val)

        i.next = i.next.next
        

        return head