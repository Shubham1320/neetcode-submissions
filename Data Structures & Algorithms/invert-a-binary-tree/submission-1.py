# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:
    def invertTree(self, root: Optional[TreeNode]) -> Optional[TreeNode]:
        
        def invert(node):

            if node == None:
                return

            if node.left == None and node.right == None:
                return
            
            if node.left == None:
                invert(node.right)
                node.left = node.right
                node.right = None
            elif node.right == None:
                invert(node.left)
                node.right = node.left
                node.left = None
            else:
                invert(node.left)
                invert(node.right)
                left = node.left
                node.left = node.right
                node.right = left
            
            return

        invert(root)
        return root
            