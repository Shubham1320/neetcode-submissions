# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:

    def traverse(self, node, p, q):
        
        if node == None:
            return None
        
        if p == node.val or q == node.val:
            return node
        elif node.val < p and node.val < q:
            return self.traverse(node.right, p, q)
        elif p < node.val and q < node.val:
            return self.traverse(node.left, p, q)
        elif p < node.val and q > node.val:
            return node
        elif q < node.val and p > node.val:
            return node


    def lowestCommonAncestor(self, root: TreeNode, p: TreeNode, q: TreeNode) -> TreeNode:
        return self.traverse(root, p.val, q.val)