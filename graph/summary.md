
# Core Principles

* How to `traverse` a `graph`?    
    1. Similar as `tree` 
    2. But need to avoid repeated-visiting, usually by a `visited` array.

* Code frame

    ```
    // 记录被遍历过的节点
    var visited []bool
    // 记录从起点到当前节点的路径
    var onPath []bool

    func traverse(graph Graph, s int) {
        if visited[s] {
            return
        }
        // 经过节点 s，标记为已遍历
        visited[s] = true
        // 做选择：标记节点 s 在路径上
        onPath[s] = true
        for _, neighbor := range graph.neighbors(s) {
            traverse(graph, neighbor)
        }
        // 撤销选择：节点 s 离开路径
        onPath[s] = false
    }
    ```
    1. The difference of `visisted` and `onPath`.
    2. Operations on `onPath` is very similar to `backtrack`?
        >这个 onPath 数组的操作很像前文 回溯算法核心套路 中做「做选择」和「撤销选择」，区别在于位置：回溯算法的「做选择」和「撤销选择」在 for 循环里面，而对 onPath 数组的操作在 for 循环外面。为什么有这个区别呢？这就是回溯算法和 DFS 算法的区别所在：回溯算法关注的不是节点，而是树枝。
      
        The difference on code
        ```
        // dfs. Attention on the nodes.
        func traverse(root *TreeNode) {
            if root == nil {
                return
            }
            fmt.Printf("进入节点 %s", root)
            for _, child := range root.children {
                traverse(child)
            }
            fmt.Printf("离开节点 %s", root)
        }
        // backtrack. Attention on the edges
        func backtrack(root *TreeNode) {
            if root == nil {
                return
            }
            for _, child := range root.children {
                fmt.Printf("从 %s 到 %s", root, child) // choose
                backtrack(child) // next layer
                fmt.Printf("从 %s 到 %s", child, root) // cancel choose
            }
        }
        ```

# Categories

### DAG 

| Problems | Key Points | Possible Solutions | Comments |
| :- |:- |:- | :- | 
| [797. All Paths From Source to Target](https://leetcode.com/problems/all-paths-from-source-to-target/description/) | DFS | [code](graph_lc797.go) | | 

### Bipartite Graph 

##### Problems

| Problems | Key Points | Possible Solutions | code| Comments |
| :- |:- |:- | :- | :-- |
| [785. Is Graph Bipartite?](https://leetcode.com/problems/is-graph-bipartite/description/) | It's possible that the graph is NOT connected | DFS, BFS | [code](graph_lc785.go) | | 
| [886. Possible Bipartition](https://leetcode.com/problems/possible-bipartition/description/) | How to transfer the problem to the Bipartite Graph Problem | DFS, BFS| [code](graph_lc886.go) | | 

###### Usage
1. It's useful with `m:n` model. For example, reference between `movies` and `acters`. 
2. Ford–Fulkerson algorithm


##### Code Frame
```
func traverse(graph Graph, visited []bool, v int) {
	visited[v] = true
    // 遍历节点 v 的所有相邻节点 neighbor
	for _, neighbor := range graph.Neighbors(v) {
		if !visited[neighbor] { // not visted yet, visit it, and color it with different color.
			traverse(graph, visited, neighbor)
		} else { // visted, compare if the color is different. if NO, the graph is not bipartite
		}
	}
}
```

### Ring Check

##### Problems

| Problems | Key Points | Possible Solutions | code| Comments |
| :- |:- |:- | :- | :-- |
| [207. Course Schedule](https://leetcode.com/problems/course-schedule/description/) | * Difference of `visited` and `onPath`<br>* Construct graph | DFS, BFS | [code](graph_lc207.go) | | 


### Topological Sorting

##### Concepts

What's `Topological Sorting`?
>直观地说就是，让你把一幅图「拉平」，而且这个「拉平」的图里面，所有箭头方向都是一致的，比如上图所有箭头都是朝右的。很显然，如果一幅有向图中存在环，是无法进行拓扑排序的，因为肯定做不到所有箭头方向一致；反过来，如果一幅图是「有向无环图」，那么一定可以进行拓扑排序。As for problem 210, 如果把课程抽象成节点，课程之间的依赖关系抽象成有向边，那么这幅图的拓扑排序结果就是上课顺序。

##### Problems

| Problems | Key Points | Possible Solutions | code| Comments |
| :- |:- |:- | :- | :-- |
| [210. Course Schedule II](https://leetcode.com/problems/course-schedule-ii/description/) | * Directions mean `Depending on` or `Depended by`? <br>* Pre-order or Post-order to collect the current node? | DFS, BFS | [code](graph_lc210.go) | | 


### Connectivity Problems - Union Find 

##### Concepts

>Connectivity problems of graph
>自反性,对称性,传递性


##### Usage

* Connectivity of graph
* Satisfiability of Equality Equations
* Minimum Spanning Tree

##### Skills 

>使用森林（若干棵树）来表示图的动态连通性，用数组来具体实现这个森林。
>
>1. 用 parent 数组记录每个节点的父节点，相当于指向父节点的指针，所以 parent 数组内实际存储着一个森林（若干棵多叉树）。
>2. 用 size 数组记录着每棵树的重量，目的是让 union 后树依然拥有平衡性，保证各个 API 时间复杂度为 O(logN)，而不会退化成链表影响操作效率。
>3. 在 find 函数中进行路径压缩，保证任意树的高度保持在常数，使得各个 API 时间复杂度为 O(1)。使用了路径压缩之后，可以不使用 size 数组的平衡优化。

Code frame and optimization refer to [code](./common.go)

##### Problems

| Problems | Key Points | Possible Solutions | code| Comments |
| :- |:- |:- | :- | :-- |
| [323. Number of Connected Components in an Undirected Graph](https://leetcode.com/problems/number-of-connected-components-in-an-undirected-graph/description/) | | Union Find | [code](graph_lc323.go) | | 
| [130. Surrounded Regions](https://leetcode.com/problems/surrounded-regions/description/) |  | * DFS(islands problem) <br>* Union Find | [code](graph_lc130.go) | | 
| [990. Satisfiability of Equality Equations](https://leetcode.com/problems/satisfiability-of-equality-equations/description/) | | Union Find | [code](graph_lc990.go) | | 
| [261. Graph Valid Tree](https://leetcode.com/problems/graph-valid-tree/description/) | * How to separate the graph and tree? | Union Find | [code](graph_lc261.go) | | 


### Kruskal Algorithm

##### Concepts

* Minimum Spanning Tree

  > 先说「树」和「图」的根本区别：树不会包含环，图可以包含环.
如果一幅图没有环，完全可以拉伸成一棵树的模样。树就是「无环连通图」。
那么什么是图的「生成树」呢，就是在图中找一棵包含图中的所有节点的树。生成树是含有图中所有顶点的「无环连通子图」。
一幅图可以有很多不同的生成树，对于加权图，每条边都有权重，所以每棵生成树都有一个权重和。
权重和最小的那棵生成树就叫「最小生成树」。

##### Skills 

* How to generate a MST?

  > 1、包含图中的所有节点。
2、形成的结构是树结构（即不存在环）。
3、权重和最小。
前两条其实可以很容易地利用 Union-Find 算法做到，关键在于第 3 点，如何保证得到的这棵生成树是权重和最小的。
这里就用到了贪心思路：
将所有边按照权重从小到大排序，从权重最小的边开始遍历，如果这条边和 mst 中的其它边不会形成环，则这条边是最小生成树的一部分，将它加入 mst 集合；否则，这条边不是最小生成树的一部分，不要把它加入 mst 集合。

In short, `Kruskal MST` can be considered as `Union Find` + `Greedy`


##### Problems

| Problems | Key Points | Possible Solutions | code| Comments |
| :- |:- |:- | :- | :-- |
| [1135. Connecting Cities With Minimum Cost](https://leetcode.com/problems/connecting-cities-with-minimum-cost/description/) | | * Kruskal MST | [code](graph_lc1135.go) | | 
| [1584. Min Cost to Connect All Points](https://leetcode.com/problems/min-cost-to-connect-all-points/description/) | | * Kruskal MST <br> * [For-Loop](graph_lc1584.go#L47) | [code](graph_lc1584.go) | | 

# References
1. [Graph Traverse](https://labuladong.online/algo/data-structure/graph-traverse/)
2. [Bipartite Problems](https://labuladong.online/algo/data-structure/bipartite-graph/)
3. [Ring Check Problems](https://labuladong.online/algo/data-structure/topological-sort/)
4. [Topological Sorting](https://labuladong.online/algo/data-structure/topological-sort/)
5. [Union Find](https://labuladong.online/algo/data-structure/union-find/)
6. [Kruskal Algorithm](https://labuladong.online/algo/data-structure/kruskal/)
7. []
