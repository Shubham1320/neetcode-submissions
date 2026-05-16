# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next


class Solution:
    def hasCycle(self, head: Optional[ListNode]) -> bool:
        
        
        
        i = head
        if head.next:
            j = head.next.next
        else:
            return False

        while j != None:
            if i == j:
                return True
            i = i .next
            j = j.next.next

        return False

        