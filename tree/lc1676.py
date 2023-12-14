class Solution(object):

    found = None

    def lowestCommonAncestor(self, root, nodes):
        nodes = set(nodes)
        self.traverse(root, nodes)
        return self.found

    def traverse(self, root, nodes):
        if not root:
            return 0
        
        count = 0
        left = self.traverse(root.left, nodes)
        right = self.traverse(root.right, nodes)

        count += left + right

        if root in nodes:
            count += 1

        if count == len(nodes):
            self.found = self.found or root
        
        return count
       
    def lowestCommonAncestorV2(self, root, nodes):
        nodes = set(nodes)
        return self.traverseV2(root, nodes)

    def traverseV2(self, root, nodes):
        if not root:
            return None

        if root in nodes:
            return root

        left = self.traverseV2(root.left, nodes)
        right = self.traverseV2(root.right, nodes)

        if left and right:
            return root

        return left or right
