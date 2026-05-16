# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Codec:
    
    # Encodes a tree to a single string.
    def serialize(self, root: Optional[TreeNode]) -> str:

        treeArr = []

        if root == None:
            return ""

        q = []
        q.append(root)
        treeArr.append(root.val)

        while len(q) != 0:
            node = q.pop(0)

            if node.left:
                q.append(node.left)
                treeArr.append(node.left.val)
            else:
                treeArr.append('null')
            
            if node.right:
                q.append(node.right)
                treeArr.append(node.right.val)
            else:
                treeArr.append('null')

        print(treeArr)
        ans_str = ""
        for item in treeArr:
            ans_str += "{}|".format(item)

        return ans_str[:-1]

        
    # Decodes your encoded data to tree.
    def deserialize(self, data: str) -> Optional[TreeNode]:

        print(data)

        if data == "":
            return None

        data = data.split('|')

        index = 1

        q = []

        root = TreeNode(val=int(data[0]))
        q.append(root)

        while len(q) != 0:
            node = q.pop(0)

            if data[index] != 'null':
                node.left = TreeNode(val=int(data[index]))
                q.append(node.left)
            else:
                node.left = None
            index += 1
            
            if data[index] != 'null':
                node.right = TreeNode(val=int(data[index]))
                q.append(node.right)
            else:
                node.right = None
            index += 1

        return root
