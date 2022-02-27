package leetcode

func reviewExist(board [][]byte, word string) bool {
	xLength, yLength := len(board), len(board[0])
	for x := 0; x < xLength; x++ {
		for y := 0; y < yLength; y++ {
			if helper(x, y, board, 0, word) {
				return true
			}
		}
	}
	return false
}

func helper(x, y int, board [][]byte, matchIdx int, word string) bool {
	if matchIdx == len(word) {
		return true
	}
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) || board[x][y] != word[matchIdx] {
		return false
	}
	tmp := board[x][y]
	board[x][y] = '0'
	defer func() {
		board[x][y] = tmp
	}()
	return helper(x-1, y, board, matchIdx+1, word) ||
		helper(x+1, y, board, matchIdx+1, word) ||
		helper(x, y+1, board, matchIdx+1, word) ||
		helper(x, y-1, board, matchIdx+1, word)
}
