
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

### Ancestor Problems

| Problems | Solutions | Key Points | code| Comments |
| :- |:- |:- | :- | :-- |
| [236. Lowest Common Ancestor of a Binary Tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/description/) | 1. post-order DFS<br>2. Track the path from root to target node and compare the slice | 1. Only 2 situations: one left, on right and one is parent of another <br> 2. The loop order of the track slice| [code](tree_lc236.go) | | 
| [235. Lowest Common Ancestor of a Binary Search Tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/description/) | Same solutions as 236 | 1. Utilize the character of BST | [code](tree_lc235.go) | | 
| [1650. Lowest Common Ancestor of a Binary Tree III](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree-iii/description/) | Linked list | 1. Check the public node 2 linked list | [code](tree_lc1650.go) | | 
| [1257. Smallest Common Region](https://leetcode.com/problems/smallest-common-region/description/) | Same solutions as 1650 | | [code](tree_lc1257.go) | | 


### BST Problems

| Problems | Solutions | Key Points | code| Comments |
| :- |:- |:- | :- | :-- |
| [108. Convert Sorted Array to Binary Search Tree](https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/description/) | | | [code](tree_lc108.go) | | 
| [1382. Balance a Binary Search Tree](https://leetcode.com/problems/balance-a-binary-search-tree/description/) | in-order traverse + convert ordered array to BST | | [code](tree_lc1382.go) | | 

### BST Problems - Construction
| [96. Unique Binary Search Trees](https://leetcode.com/problems/unique-binary-search-trees/description/) | DFS post-order | How to set memo effectively | [code](tree_lc96.go) | | 
| [95. Unique Binary Search Trees II](https://leetcode.com/problems/unique-binary-search-trees-ii/description/) | DFS post-order | Edged case of `low > high` and `low == high` | [code](tree_lc95.go) | | 

Reference
1. https://labuladong.online/algo/data-structure/bst-part3/

### Other Problems
| Problems | Solutions | Key Points | code| Comments |
| :- |:- |:- | :- | :-- |
| [1530. Number of Good Leaf Nodes Pairs](https://leetcode.com/problems/number-of-good-leaf-nodes-pairs/description/) | DFS | * Count the pairs of left sub-tree, right sub-tree and current node <br> * How to calculate the distance? Utilized height | [code](tree_lc1530.go) | | 
| [993. Cousins in Binary Tree](https://leetcode.com/problems/cousins-in-binary-tree/description/) | DFS, BFS | | [code](tree_lc993.go) | | 

**Binary Tree**

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
        * Common solution.
* LC652. Find Duplicate Subtrees
    * DFS. Post order
    * Key Points
        * Generate a `signature` for each sub tree
            * Serilize the tree or
            * generate an id for it. Pay attention to the id of `nil node`
        * Time: O(n). O(n), O(1)
        * Space: O(n)
* LC315. Count of Smaller Numbers After Self
    * Merge Sort
    * BIT
    * Complexity:
      * T:
        * Copy middle list: O(N)
        * Merge Sort: O(N*LogN) 
        * Setup Index Map: O(N)
        * Query and Update BIT: O(N*LogN)*2
      * Space
        * Middle list: O(N)
        * Index Map: O(N)
        * BIT: O(N)
* LC493. Reverse Pairs
    * Similar as LC315.
      * Reverse BIT
    * Solution 2. Only with Merge Sort
      * Key Points
        * Count first and then sort
        * count += middle-i+1
        * Complexity
          * T: 
            * Compare and Sort: 2(N*logN)
          * S:
            * O(N)
      * Why LC315 can not resolved with only Merge Sort?

* LC1664. Lowest Common Ancestor of a Binary Tree II.
  * Key Points.
    * Keep the idea of pre-order, in-order, post-order traverse...
* LC1676. Lowest Common Ancestor of a Binary Tree IV
  * Key Point.
    * Difference from LC1664??

**Complete Binary Tree**
  * Principle
    > According to Wikipedia, every level, except possibly the last, is completely filled in a complete binary tree, and all nodes in the last level are as far left as possible. It can have between 1 and 2h nodes inclusive at the last level h.

    * For any node, if it has children, then, there must be at least 1 `Perfect Binary Tree` in it's children. 
    
**Perfect Binary Tree**
  * Principle
    * Every level of the tree is full
    
**Full Binary Tree**
  * Principle
    * Every node has either has both of the 2 children or none of the 2 children.
    
**Binary Search Tree**

* Key Points
  * Use closure to reduce the complexity of passing parameters
  *
* LC230. Kth Smallest Element in a BST
    * Inorder traverse
    * Follow Up
      * Extend the BST with `NumberSumOfSubtree`
      * The operation of `Insert/Delete` elements into/from BST

* LC538. LC1038. Convert BST to Greater Tree (same as 1038)
    * DFT and right sub tree first (in order)


* LC98. Validate Binary Search Tree
  * Key Points
    * Not only the left and right child. But also all the left and right subtree
* LC700. Search in a Binary Search Tree.
* LC701. Insert into a Binary Search Tree.
* LC450. Delete Node in a BST

* LC95. Unique Binary Search Trees II


**Binary Indexed Tree**

![](../images/tree_bit.png?raw=true)
