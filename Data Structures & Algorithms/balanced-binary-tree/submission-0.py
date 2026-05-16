# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:

    balanced = True

    def traverse(self, node):
        if node == None:
            return 0

        leftDepth = self.traverse(node.left)
        rightDepth = self.traverse(node.right)

        print(node.val, leftDepth, rightDepth)
        if abs(leftDepth - rightDepth) > 1:
            self.balanced = False
        
        return max(leftDepth, rightDepth) + 1

    def isBalanced(self, root: Optional[TreeNode]) -> bool:
        
        self.traverse(root)
        return self.balanced