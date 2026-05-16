# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

class Solution:

    def reverseList(self, start, end):
        i = start
        if i == end:
            return start, end, i.next
        j = i.next
        if i.next == end:
            newNext = j.next
            j.next = i
            i.next = None
            return j, i, newNext 
        k = j.next
        
        i.next = None

        while k != end.next:
            #print(i.val, j.val, k.val)
            j.next = i
            i = j
            j = k
            k = j.next
        j.next = i

        #print(j.val, i.val,start.val)

        l = j
        while l != start:
            #print(l.val)
            l = l.next
        #print(l.val)

        return j, start, k


    def reverseKGroup(self, head: Optional[ListNode], k: int) -> Optional[ListNode]:
        
        i = head
        j = head

        firstPass = True
        reversedLists = []

        while i != None:
            iterator = 1
            done = False
            while iterator < k:
                if j.next == None:
                    #print(i.val, j.val)
                    reversedLists.append([i,j])
                    done = True
                    break
                j = j.next
                iterator += 1
            if done:
                break
            prevFirst, prevLast, nextNode = self.reverseList(i,j)
            if firstPass:
                head = prevFirst
                firstPass = False
            reversedLists.append([prevFirst, prevLast])

            i = nextNode
            j = i
            
        i = 1
        while i < len(reversedLists):
            #print(reversedLists[i][0].val, reversedLists[i][1].val)
            reversedLists[i-1][1].next = reversedLists[i][0]
            i += 1

        return head
            
        
            
