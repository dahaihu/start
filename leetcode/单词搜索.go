package leetcode

//给定一个二维网格和一个单词，找出该单词是否存在于网格中。
//
//单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
//
// 
//
//示例:
//
//board =
//[
//  ['A','B','C','E'],
//  ['S','F','C','S'],
//  ['A','D','E','E']
//]
//
//给定 word = "ABCCED", 返回 true
//给定 word = "SEE", 返回 true
//给定 word = "ABCB", 返回 false
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/word-search
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if checkExist(board, i, j, 0, word) {
				return true
			}
		}
	}
	return false
}

func checkExist(board [][]byte, x, y, idx int, word string) bool {
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) {
		return false
	}

	if board[x][y] != word[idx] {
		return false
	}

	if idx == len(word)-1 {
		return true
	}

	tmp := board[x][y]
	board[x][y] = '.'
	res := checkExist(board, x-1, y, idx+1, word)
	res = res || checkExist(board, x+1, y, idx+1, word)
	res = res || checkExist(board, x, y+1, idx+1, word)
	res = res || checkExist(board, x, y-1, idx+1, word)

	board[x][y] = tmp
	return res
}
