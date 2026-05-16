# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:

    maxLevel = 0

    def traverse(self, node, level):

        if node == None:
            return

        print(level)
        print(self.maxLevel)
        if level > self.maxLevel:
            self.maxLevel = level

        if node.left != None:
            self.traverse(node.left, level + 1)
            
        if node.right != None:
            self.traverse(node.right, level + 1)

    def maxDepth(self, root: Optional[TreeNode]) -> int:

        self.traverse(root, 1)

        return self.maxLevel