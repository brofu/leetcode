
**Principles**

1. The `diameter` is the largest value of `sum(max_depth_of_left_subtree, max_depth_of_right_subtree)` of all of nodes of the tree

**Traverse**
1. BF
    * Usually based on `Queue`
    * Can also use recursive
2. DF
    * PreOrder
    * InOrder
    * PostOrder


**Thinking Patterns**
1. Tree Traverse. `Traverse` 
2. Recursive with sub trees. `Sub Tasks`

> 二叉树解题的思维模式分两类：
1、是否可以通过遍历一遍二叉树得到答案？如果可以，用一个 traverse 函数配合外部变量来实现，这叫「遍历」的思维模式。
2、是否可以定义一个递归函数，通过子问题（子树）的答案推导出原问题的答案？如果可以，写出这个递归函数的定义，并充分利用这个函数的返回值，这叫「分解问题」的思维模式。
无论使用哪种思维模式，你都需要思考：
如果单独抽出一个二叉树节点，它需要做什么事情？需要在什么时候（前/中/后序位置）做？其他的节点不用你操心，递归函数会帮你在所有节点上执行相同的操作。


**Problems**
* LC116. Populating Next Right Pointers in Each Node
    * Why the normal `traverse` approach doesn't work with this problem? O(N), O(1)
    * Util BFS `traverse`. O(N), O(N)
    * LC117. Populating Next Right Pointers in Each Node II
        * BFS approach. O(N), O(N)
        * BFS approach with cursive. O(N), O(N)
        * Multiple pointers approach & sentinel. O(N), O(Cons). And the approach also work with LC116.
* LC114. Flatten Binary Tree to Linked List
    * Traverse the tree, construct the link.
    * Flat recursively

***Construction***

* LC105. Construct binary tree from preorder and inorder array
    * Recursively.
    * Key points
      * Bound of array
* LC106. Construct binary tree from preorder and inorder array
    * Recursively.
    * Key points
      * Bound of array
* LC654. Maximum Binary Tree
    * Recursively
* LC889. Construct binary tree from preorder and postorder array  
    * Recursively
    * Key Points
        * 2 situations

