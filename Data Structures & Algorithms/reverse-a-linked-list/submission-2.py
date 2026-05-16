# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next


#0.     1.      2.     3.   

#head   i.      j

#       head.   i.      j

#               head.   i.     j


class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        
        if head == None:
            return head

        if head.next == None:
            return head

        i = head.next

        if i.next == None:
            i.next = head
            head.next = None
            head = i
            return head

        j = i.next



        head.next = None
        
        while i != None:
            i.next = head
            head = i
            i = j
            if i is not None:
                j = i.next

        return head