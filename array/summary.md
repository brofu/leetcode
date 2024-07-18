
# Problems

### N-Sum Problems
| Types | Problems | Key Points | Possible Solutions | Comments |
| :- | :- | :- |:- | :- | 
| nSum Problems | [259](https://leetcode.com/problems/3sum-smaller/description/) | * Double pointer <br>* Remove the duplication | [code](array_lc259.go) | | 


### LIS Problems

| Problems | Possible Solutions | Key Points | Code | Comments |
| :- | :- | :- |:- | :- | 
| [674. Longest Continuous Increasing Subsequence](https://leetcode.com/problems/longest-continuous-increasing-subsequence/description/) | | Edged case: the array ALL is increasing | [code](array_lc674.go) | |


### MergeSort Related Problems

**Key Principles**

Break down the problems into sub problems and the merge. `Divide and Conquer` 

| Problems | Possible Solutions | Key Points | Code | Comments |
| :- | :- | :- |:- | :- | 
| [315. Count of Smaller Numbers After Self](https://leetcode.com/problems/count-of-smaller-numbers-after-self/description/) |* Merge Sort<br>* Binary Indexed Tree<br> * Segment Tree | Principle why merge sort can work (R1) | [code](array_lc315.go) | |
| [493. Reverse Pairs](https://leetcode.com/problems/reverse-pairs/description/) |* Merge Sort<br>* Binary Indexed Tree<br>* Segment Tree | How to insert the logic into MergeSort | [code](array_lc493.go) | |
| [327. Count of Range Sum](https://leetcode.com/problems/count-of-range-sum/description/) |* Merge Sort<br>* Binary Indexed Tree<br>* Segment Tree | * Convert the problem to presum problem <br>* merge sort<br>* Refer to the code | [code](array_lc327.go) | |


Reference:
R1. https://labuladong.online/algo/practice-in-action/merge-sort/

### PreSum Problems

**Key Principles**

| Problems | Possible Solutions | Key Points | Code | Comments |
| :- | :- | :- |:- | :- | 
| [303. Range Sum Query - Immutable](https://leetcode.com/problems/range-sum-query-immutable/description/) | Presum | * If the `left` num is included | [code](array_lc303.go) | |
| [304. Range Sum Query 2D - Immutable](https://leetcode.com/problems/range-sum-query-2d-immutable/description/) | Presum | * if `left` is included<br> * the set calculation | [code](array_lc304.go) | |


Reference:
R1. https://labuladong.online/algo/practice-in-action/merge-sort://labuladong.online/algo/data-structure/prefix-sum/ 


### BitMask Problems

| Problems | Possible Solutions | Key Points | Code | Comments |
| :- | :- | :- |:- | :- | 
| [2506. Count Pairs Of Similar Strings](https://leetcode.com/problems/count-pairs-of-similar-strings/description/) | bitmask | should contain SAME characters | [code](array_lc2506.go) | |
