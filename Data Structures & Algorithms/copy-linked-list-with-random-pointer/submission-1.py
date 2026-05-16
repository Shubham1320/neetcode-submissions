"""
# Definition for a Node.
class Node:
    def __init__(self, x: int, next: 'Node' = None, random: 'Node' = None):
        self.val = int(x)
        self.next = next
        self.random = random
"""

class Solution:
    def copyRandomList(self, head: 'Optional[Node]') -> 'Optional[Node]':
        hashmap = {}

        i = head

        if head is None:
            return None

        head1 = Node(x=i.val)
        j = head1

        hashmap[i] = head1

        while i != None:
            if i.next and i.next not in hashmap.keys():
                newNode = Node(x=i.next.val)
                j.next = newNode
                hashmap[i.next] = newNode
            elif i.next:
                j.next = hashmap[i.next]

            if i.random != None:

                if i.random not in hashmap.keys():
                    newNode = Node(x=i.random.val)
                    j.random = newNode
                    hashmap[i.random] = newNode
                else:
                    j.random = hashmap[i.random]
            else:
                j.random = None

            i = i.next
            j = j.next

        return head1

