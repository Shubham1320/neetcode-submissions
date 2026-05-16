# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:

    def rightSideView(self, root: Optional[TreeNode]) -> List[int]:
        
        q = []
        tempQ = []
        ans = []

        if root == None:
            return ans

        q.append(root)

        while len(q) != 0:
            level = []
            for i in q:
                level.append(i.val)
                if i.left:
                    tempQ.append(i.left)
                if i.right:
                    tempQ.append(i.right)
            ans.append(level[-1])
            q = tempQ
            tempQ = []

        return ans