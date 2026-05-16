# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:

    def traverse(self, node, left, right):
        if not node:
            return True

        if not (node.val < right and node.val > left):
            return False

        return self.traverse(node.left, left, node.val) and \
                self.traverse(node.right, node.val, right)

    def isValidBST(self, root: Optional[TreeNode]) -> bool:
        return self.traverse(root, float("-inf"), float("inf"))
        