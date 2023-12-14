class Solution(object):

    def __init__(self):
        self.foundP = False 
        self.foundQ = False

    def lowestCommonAncestorV2(self, root, p, q):

        result = self.traverseV2(root, p, q) 

        if not result or result.val != p.val and result.val != q.val:
            return result 
        if result.val == p.val and not self.traverseV2(root, q, q) or result.val == q.val and not self.traverseV2(root, p, p):
            return None
        return result

    def traverseV2(self, root, p, q):
        if not root or root.val == p.val or root.val == q.val:
            return root

        left = self.traverseV2(root.left, p, q)
        right = self.traverseV2(root.right, p, q)
        
        if not left:
            return right
        if not right:
            return left
        return root

    def lowestCommonAncestor(self, root, p, q):
        """
        :type root: TreeNode
        :type p: TreeNode
        :type q: TreeNode
        :rtype: TreeNode
        """
        result = self.traverse(root, p, q)
        if self.foundP and self.foundQ:
            return result
        return None

    def traverse(self, root, p, q):
        if not root:
            return None

        left = self.traverse(root.left, p, q)
        right = self.traverse(root.right, p, q)

        if root.val == p.val:
            self.foundP = True
            return root

        if root.val == q.val:
            self.foundQ = True
            return root
  

        if not left:
            return right

        if not right:
            return left

        return root

