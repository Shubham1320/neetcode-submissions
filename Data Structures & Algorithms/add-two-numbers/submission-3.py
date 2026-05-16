# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        
        head = None
        i = None
        carryOn = 0

        while True:
            if l1 == None or l2 == None:
                break
            
            first_num = int(l1.val)
            second_num = int(l2.val)

            ans = (first_num + second_num + carryOn) % 10
            carryOn = int((first_num + second_num + carryOn) / 10)
            print(carryOn)

            newNode = ListNode(val=ans)

            if i != None:
                i.next = newNode
                i = newNode

            if head == None:
                head = newNode
                i = head

            l1 = l1.next
            l2 = l2.next

        while l1 != None:
            ans = (l1.val + carryOn) % 10
            carryOn = int((l1.val + carryOn) / 10)
            print(ans, carryOn)
            newNode = ListNode(val=ans)
            i.next = newNode
            i = newNode
            l1= l1.next

        while l2 != None:
            ans = (l2.val + carryOn) % 10
            carryOn = int((l2.val + carryOn) / 10)
            newNode = ListNode(val=ans)
            i.next = newNode
            i = newNode
            l2 = l2.next

        if carryOn != 0:
            newNode = ListNode(val=carryOn)
            i.next = newNode

        print(carryOn)

        return head