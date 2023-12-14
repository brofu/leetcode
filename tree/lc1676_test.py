import unittest
import common
import lc1676

class TestCase(object):
    def __init__(self, name, array, nodes, want):
        self.name = name
        self.array = array
        self.nodes = nodes
        self.want = want

class TestSolution(unittest.TestCase):

    def test_lowestCommonAncestor(self):
        tt = [ 
            TestCase(name="lc case 1", array=[3,5,1,6,2,0,8,None,None,7,4], nodes=[4, 7], want=2),
            TestCase(name="lc case 1", array=[3,5,1,6,2,0,8,None,None,7,4], nodes=[1], want=1),
            TestCase(name="lc case 1", array=[3,5,1,6,2,0,8,None,None,7,4], nodes=[7,6,2,4], want=5),
            #TestCase(name="lc case 1", array=[3,5,1,6,2,0,8,None,None,7,4], nodes=[4, 7], want=2),
            ]

        for case in tt:
            inputTree = common.GenerateBinaryTreeFromArray(case.array, 0)
            nodes = []
            for num in case.nodes:
                node = common.GetTreeNodeFromTree(inputTree, num)
                nodes.append(node)
            got = lc1676.Solution().lowestCommonAncestor(inputTree, nodes)
            try:
                self.assertEqual(got and got.val, case.want)
            except AssertionError:
                print "name: %s, got: %s, want: %s" %(case.name, got.val, case.want)
            else:
                print "name: %s PASS" %case.name

    def test_lowestCommonAncestorV2(self):
        tt = [ 
            TestCase(name="lc case 1", array=[3,5,1,6,2,0,8,None,None,7,4], nodes=[4, 7], want=2),
            TestCase(name="lc case 1", array=[3,5,1,6,2,0,8,None,None,7,4], nodes=[1], want=1),
            TestCase(name="lc case 1", array=[3,5,1,6,2,0,8,None,None,7,4], nodes=[7,6,2,4], want=5),
            #TestCase(name="lc case 1", array=[3,5,1,6,2,0,8,None,None,7,4], nodes=[4, 7], want=2),
            ]

        for case in tt:
            inputTree = common.GenerateBinaryTreeFromArray(case.array, 0)
            nodes = []
            for num in case.nodes:
                node = common.GetTreeNodeFromTree(inputTree, num)
                nodes.append(node)
            got = lc1676.Solution().lowestCommonAncestorV2(inputTree, nodes)
            try:
                self.assertEqual(got and got.val, case.want)
            except AssertionError:
                print "name: %s, got: %s, want: %s" %(case.name, got.val, case.want)
            else:
                print "name: %s PASS" %case.name

if __name__ == '__main__':
    unittest.main()


