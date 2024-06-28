
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

# Problems

### DAG 

| Problems | Key Points | Possible Solutions | Comments |
| :- |:- |:- | :- | 
| [797. All Paths From Source to Target](https://leetcode.com/problems/all-paths-from-source-to-target/description/) | DFS | [code](graph_lc797.go) | | 

### Bipartite Graph 

##### Bipartite Graph Principles

**Usage**
1. It's useful with `m:n` model. For example, reference between `movies` and `acters`. 
2. Ford–Fulkerson algorithm


**Code Frame**
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


# References
1. https://labuladong.online/algo/data-structure/graph-traverse/