
class TreeNode(object):
    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
 
def GenerateBinaryTreeFromArray(array, index):
   
    if index >= len(array):
        return None

    if array[index] is None:
        return None

    left = GenerateBinaryTreeFromArray(array, 2*index+1)
    right = GenerateBinaryTreeFromArray(array, 2*index+2)
    node = TreeNode(array[index], left, right)

    return node

def PrintBinaryTreeWithDeepthFirstTraverse(root):
    nodes =[root] 
    level = 1 
    while len(nodes) > 0:
        levelNodes = []
        for index in range(level):
            if len(nodes) != 0:
                node = nodes[0] 
                nodes = nodes[1:]
                if node is None:
                    levelNodes.append("Null")
                    break
                levelNodes.append(node.val)
                nodes.append(node.left)
                nodes.append(node.right)
        print levelNodes
        level += 1
