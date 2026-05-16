# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right





class Solution:   

    found = False
 
    def compareTrees(self, tree1, tree2):
    
        if not tree1 and not tree2:
            return True

        if tree1 and tree2 and tree1.val == tree2.val:
            return self.compareTrees(tree1.left, tree2.left) and \
                    self.compareTrees(tree1.right, tree2.right)
        
        return False

    def traverse(self, node, subRoot):
        if node == None:
            return

        if node.val == subRoot.val:
            if self.compareTrees(node, subRoot):
                self.found = True

        self.traverse(node.left, subRoot)
        self.traverse(node.right, subRoot)

    def isSubtree(self, root: Optional[TreeNode], subRoot: Optional[TreeNode]) -> bool:
        self.traverse(root, subRoot)
        return self.found
