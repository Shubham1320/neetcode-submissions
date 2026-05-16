# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:

    ans = []

    def traverse(self, node, prevGoodNode):
        if node == None:
            return
        if node.val >= prevGoodNode:
            print(self.ans)
            print(node.val, prevGoodNode)
            self.ans.append(node.val)
            prevGoodNode = node.val
        
        self.traverse(node.left, prevGoodNode)
        self.traverse(node.right, prevGoodNode)

    def goodNodes(self, root: TreeNode) -> int:
        self.ans = []
        self.ans.append(root.val)
        print(self.ans)
        self.traverse(root.left, root.val)
        self.traverse(root.right, root.val)

        return len(self.ans)