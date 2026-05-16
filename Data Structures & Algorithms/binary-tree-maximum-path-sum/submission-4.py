# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:

    ans = float("-inf")

    def maxPathSum1(self, root):
        if root == None:
            return 0
            
        leftMaxPath = self.maxPathSum1(root.left)
        rightMaxPath = self.maxPathSum1(root.right)

        print(root.val, leftMaxPath , rightMaxPath)

        if root.val + leftMaxPath + rightMaxPath > self.ans:
            self.ans = root.val + leftMaxPath + rightMaxPath

        maxPath = max(root.val+leftMaxPath+rightMaxPath, root.val, root.val+leftMaxPath, root.val+rightMaxPath)

        if maxPath > self.ans:
            self.ans = maxPath

        return max(root.val, root.val+leftMaxPath, root.val+rightMaxPath)

    def maxPathSum(self, root: Optional[TreeNode]) -> int:

        ret_value = self.maxPathSum1(root)

        return max(ret_value, self.ans)



