# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:
    
    currentNum = 0
    ans = 0
    k = 0

    def traverse(self, node):

        if node == None:
            return

        print(node.val, self.k, self.currentNum)

        self.traverse(node.left)
        print('hello')
        if self.currentNum + 1 == self.k:
            print('hellloo11')
            self.ans = node.val
        self.currentNum += 1
        self.traverse(node.right)
    
    def kthSmallest(self, root: Optional[TreeNode], k: int) -> int:
        self.k = k
        self.traverse(root)
        return self.ans