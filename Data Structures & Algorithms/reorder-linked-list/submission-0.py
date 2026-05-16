# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def reorderList(self, head: Optional[ListNode]) -> None:
        
        i = head
        j = head.next

        while j != None:
            i = i.next
            if j.next:
                j = j.next.next
            else:
                j = j.next

        j = i.next
        i.next = None
        while j != None:
            print(i.val, j.val)
            j_nxt = j.next
            j.next = i
            i = j
            j = j_nxt

        j = i
        i = head
        while True:

            if i.next is None or j.next is None:
                break

            i_nxt = i.next
            j_nxt = j.next
            i.next = j
            j.next = i_nxt
            i = i_nxt
            j = j_nxt

        






        