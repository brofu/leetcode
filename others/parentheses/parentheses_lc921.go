package parentheses

/**

KP
	1.	当 needRight == -1 的时候意味着什么？ 因为只有遇到右括号 ) 的时候才会 needRight--，needRight == -1 意味着右括号太多了，所以需要插入左括号。 比如说 s = "))" 这种情况，需要插入 2 个左括号，使得 s 变成 "()()"，才是一个有效括号串。
	2.	为什么返回 needLeft + needRight？ 因为 needLeft 记录的左括号的插入次数，needRight 记录了右括号的需求，当 for 循环结束后，若 needRight 不为 0，那么就意味着右括号还不够，需要插入。 比如说 s = "))(" 这种情况，插入 2 个左括号之后，还要再插入 1 个右括号，使得 s 变成 "()()()"，才是一个有效括号串。

*/

func minAddToMakeValid(s string) int {

	needRight, needLeft := 0, 0

	for _, o := range s {

		if o == '(' {
			needRight++
		}

		if o == ')' {
			needRight--
			if needRight == -1 {
				needRight = 0 // clear needRight and
				needLeft += 1 // add needLeft
			}
		}
	}

	return needLeft + needRight
}
