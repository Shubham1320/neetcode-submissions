# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:

    def traverse(self, node, treeArr):

        if node == None:
            return

        if node.left != None:
            treeArr.append(node.left.val)
        else:
            treeArr.append('None')

        self.traverse(node.left, treeArr)

        if node.right != None:
            treeArr.append(node.right.val)
        else:
            treeArr.append('None')

        self.traverse(node.right, treeArr)


    def isSameTree(self, p: Optional[TreeNode], q: Optional[TreeNode]) -> bool:

        if p == None and q == None:
            return True
        if p == None or q == None:
            return False


        leftTree = [p.val]
        rightTree = [q.val]

        self.traverse(p, leftTree)
        self.traverse(q, rightTree)

        print(leftTree)
        print(rightTree)

        return leftTree == rightTree
