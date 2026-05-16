# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:

    perimeter = 0

    def traverse(self, node):
        if node == None:
            return 0

        leftDepth = self.traverse(node.left)
        rightDepth = self.traverse(node.right)

        print(node.val, leftDepth, rightDepth)

        if leftDepth + rightDepth + 1 > self.perimeter:
            self.perimeter = leftDepth + rightDepth + 1

        return max(leftDepth, rightDepth) + 1

    def diameterOfBinaryTree(self, root: Optional[TreeNode]) -> int:
        
        self.traverse(root)
        return self.perimeter - 1