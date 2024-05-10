# leetcode

Golang version of Leetcode problems

## Getting started

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`.

Running it then should be as simple as:

```console
$ make
$ ./bin/leetcode
```

### Testing

``make test``


## Algorithm

### Sort

| Scenarios | Problems |
| :- | :- | 
| Sort | * 912 |
| The K largest number in array | * 215 |

| algorithm |  Time Complexity | Space Complexity | Scenarios | Problems | References |
| :- | :- | :- | :- | :- | :- |
| Quick Sort | O(N*lgN) | O(1) <br> O(lgN) considering recursion| | | https://labuladong.online/algo/practice-in-action/quick-sort/ |
| Merge Sort| O(N*lgN) | O(1) <br> O(N*lgN) considering recursion| | | * https://labuladong.online/algo/practice-in-action/merge-sort/ <br> * https://www.geeksforgeeks.org/merge-sort/ |
| Heap Sort| O(N*lgN) | O(1) | | | https://www.geeksforgeeks.org/heap-sort/ |


### Dynamic Programming

| Scenarios | Problems |
| :- | :- | 
| | * 32. Longest Valid Parentheses |

## Data Structures

### Stack

| Scenarios | Problems |
| :- | :- | 
| Valid Parentheses Problems| * 30  | 


### Binary Heap

| Scenarios | Problems |
| :- | :- | 
| K Merge Problem | * 23  | 
| Priority Queue Problems | * 295 Get the median of an ordered stream | 
| Graph Algorithms | ?  | 

### List

### Array & String

#### Skills

##### Quick Select 

* Problems

  * Get the K largest number without sorting the whole array

* Examples

  * 215

##### Sliding Window

* Problem

    * Sub-string problems

* Examples

    * 3
    * 76
    * 438
    * 567

* References

    * https://labuladong.online/algo/essential-technique/sliding-window-framework/

### Tree

#### Skills

##### BFS

* Problems

    1. Min depth of binary tree
    2. Min steps to open lock
    3. Shortest distance from the `start` to the `target` node in a `graph`


* Examples

    * 103
    * 111
    * 752 

* References

    * https://labuladong.online/algo/essential-technique/bfs-framework/


#### DFS

  * Binary Search Tree
  * Binary Indexed Tree
  * AVL Tree
      * References: https://www.geeksforgeeks.org/introduction-to-avl-tree/
  * Red-Black Tree
      * [ ] References: https://geeksforgeeks.org/introduction-to-red-black-tree


#### BackTrack

* Problems

    1. Permute number
    2. Queen Problems
    3. Restore IP address
    4. Sudoku

* Examples

    * 46
    * 51
    * 52

* References

    * https://labuladong.online/algo/essential-technique/backtrack-framework/
    
### Graph

#### Heap
  * Presentation
    * Array. A complete binary tree
  * Applications
    * Heap Sort
    * Priority Queue
      * K-group Merge 
    * Graph Algorithims?
      * Dijkstra
      *
  * References
    * https://labuladong.github.io/algo/di-yi-zhan-da78c/shou-ba-sh-daeca/er-cha-dui-1a386/
    * https://www.geeksforgeeks.org/heap-sort/
    * https://www.geeksforgeeks.org/difference-between-min-heap-and-max-heap/



### LinkedHashMap

#### Skills

##### LinkdedHashMap

* Problems
    * LRU Cache
    * LFU Cache

* Examples
    * 146
    * 460

* References
    * https://labuladong.online/algo/data-structure/lru-cache/
