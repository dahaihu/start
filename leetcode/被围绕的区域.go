package leetcode

func solve(board [][]byte) {
	for i := 0; i < len(board[0]); i++ {
		dfs(board, 0, i)
		dfs(board, len(board)-1, i)
	}
	for i := 1; i < len(board)-1; i++ {
		dfs(board, i, 0)
		dfs(board, i, len(board[0])-1)
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == 'o' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(board [][]byte, i, j int) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != 'O' {
		return
	}
	board[i][j] = 'o'
	dfs(board, i-1, j)
	dfs(board, i+1, j)
	dfs(board, i, j-1)
	dfs(board, i, j+1)
}
