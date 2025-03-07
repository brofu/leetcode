package tree

import (
	"fmt"
	"strings"

	"github.com/brofu/leetcode/common"
)

/*
c
c b
0 a, a 0, a
*/

/*
Sub task version
Not work. Why? An example

root: e
  - sub left: "abab"
  - sub right: "ab"

with sub task approached, "ab" < "abab"
and we get "abe" as the final result, but "abe" > "ababe"
*/
func smallestFromLeaf(root *TreeNode) string {

	if root == nil {
		return ""
	}

	rootStr := string(rune(root.Val + 97))
	if root.Left == nil && root.Right == nil {
		return rootStr
	}

	strLeft := smallestFromLeaf(root.Left)
	strRight := smallestFromLeaf(root.Right)

	if strLeft != "" && strRight != "" {
		minStr := common.MinString(strLeft+rootStr, strRight+rootStr)
		fmt.Println("flag", rootStr, minStr)
		return minStr
	}
	if strLeft == "" {
		return strRight + rootStr
	}

	return strLeft + rootStr
}

/*
Traverse version.

Key Point
 1. Need to reverse the array and reverse it back
*/
func smallestFromLeafDFS(root *TreeNode) string {
	if root == nil {
		return ""
	}

	result := ""

	var traverse func(*TreeNode, []string)
	traverse = func(root *TreeNode, path []string) {

		temp := append(path, string(rune(root.Val+97)))

		fmt.Println(temp)

		if root.Left == nil && root.Right == nil {
			common.ReverseSlice(temp)
			s := strings.Join(temp, "")
			if result == "" {
				result = s
			} else {
				result = common.MinString(result, s)
			}
			common.ReverseSlice(temp)
			return
		}

		if root.Left != nil {
			fmt.Println("left-before", temp)
			traverse(root.Left, temp)
			fmt.Println("left-after", temp)
		}
		if root.Right != nil {
			traverse(root.Right, temp)
		}
	}

	traverse(root, []string{})
	return result
}
