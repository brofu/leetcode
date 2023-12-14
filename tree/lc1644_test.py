import unittest
import common
import lc1644

class TestCase(object):
    def __init__(self, name, array, p, q, want):
        self.name = name
        self.array = array
        self.p = p
        self.q = q
        self.want = want

class TestSolution(unittest.TestCase):

    def test_lowestCommonAncestor(self):
        tt = [ 
            TestCase(name="lc case 1", array=[3,5,1,6,2,0,8,None,None,7,4], p=5, q=1, want=3),
            TestCase(name="lc case 2", array=[3,5,1,6,2,0,8,None,None,7,4], p=5, q=4, want=5),
            TestCase(name="lc case 3", array=[3,5,1,6,2,0,8,None,None,7,4], p=5, q=10, want=None),
            TestCase(name="lc case 4", array=[17,13,None,15,None,None,None,18,5,None,None,None,None,None,None,None,4,None,None,None,None,None,None,None,None,None,None,None,None,None,None,None,None,9], p=5, q=9, want=15),
            ]

        for case in tt:
            inputTree = common.GenerateBinaryTreeFromArray(case.array, 0)
            #common.PrintBinaryTreeWithDeepthFirstTraverse(inputTree)
            got = lc1644.Solution().lowestCommonAncestor(inputTree, common.TreeNode(case.p),common.TreeNode(case.q))
            try:
                self.assertEqual(got and got.val, case.want)
            except AssertionError:
                print "name: %s, got: %s, want: %s" %(case.name, got.val, case.want)
            else:
                print "name: %s PASS" %case.name

    def test_lowestCommonAncestorV2(self):
        tt = [ 
            #TestCase(name="lc case 1", array=[3,5,1,6,2,0,8,None,None,7,4], p=5, q=1, want=3),
            #TestCase(name="lc case 2", array=[3,5,1,6,2,0,8,None,None,7,4], p=5, q=4, want=5),
            #TestCase(name="lc case 3", array=[3,5,1,6,2,0,8,None,None,7,4], p=5, q=10, want=None),
            #TestCase(name="lc case 4", array=[17,13,None,15,None,18,5,None,4,None,None,9], p=5, q=10, want=13),
            ]

        for case in tt:
            inputTree = common.GenerateBinaryTreeFromArray(case.array, 0)
            got = lc1644.Solution().lowestCommonAncestorV2(inputTree, common.TreeNode(case.p),common.TreeNode(case.q))
            try:
                self.assertEqual(got and got.val, case.want)
            except AssertionError:
                print "name: %s, got: %s, want: %s" %(case.name, got, case.want)
            else:
                print "name: %s PASS" %case.name


if __name__ == '__main__':
    unittest.main()


