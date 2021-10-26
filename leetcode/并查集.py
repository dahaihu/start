# 130. 被围绕的区域
# 给你一个 m x n 的矩阵 board ，由若干字符 'X' 和 'O' ，找到所有被 'X' 围绕的区域，并将这些区域里所有的 'O' 用 'X' 填充。
#  
#
# 示例 1：
#
#
# 输入：board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]
# 输出：[["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
# 解释：被围绕的区间不会存在于边界上，换句话说，任何边界上的 'O' 都不会被填充为 'X'。 任何不在边界上，或不与边界上的 'O' 相连的 'O' 最终都会被填充为 'X'。如果两个元素在水平或垂直方向相邻，则称它们是“相连”的。
# 示例 2：
#
# 输入：board = [["X"]]
# 输出：[["X"]]
#  
#
# 提示：
#
# m == board.length
# n == board[i].length
# 1 <= m, n <= 200
# board[i][j] 为 'X' 或 'O'


class Solution:
    def __init__(self):
        self.mark = dict()

    def find(self, n):
        if n not in self.mark:
            self.mark[n] = n
            return n
        init = n
        while self.mark[n] != n:
            n = self.mark[n]
        self.mark[init] = n
        return n

    def union(self, m, n):
        m_ancestor = self.find(m)
        n_ancestor = self.find(n)
        if m_ancestor != n_ancestor:
            self.mark[m_ancestor] = n_ancestor

    def same(self, m, n):
        m_ancestor = self.find(m)
        n_ancestor = self.find(n)
        return m_ancestor == n_ancestor

    def solve(self, board):
        if len(board) == 0 or len(board[0]) == 0:
            return
        length_x, length_y = len(board), len(board[0])
        dummy = length_x * length_y
        self.find(dummy)
        for x in range(length_x):
            for y in range(length_y):
                if board[x][y] == 'X':
                    continue
                if x == 0 or x == length_x - 1 or y == 0 or y == length_y - 1:
                    self.union(x * length_y + y, dummy)
                else:
                    for gap_x, gap_y in ((0, 1), (0, -1), (1, 0), (-1, 0)):
                        next_x, next_y = x + gap_x, y + gap_y
                        if next_x < 0 or next_x >= length_x or \
                                next_y < 0 or next_y >= length_y or \
                                board[next_x][next_y] == 'X':
                            continue
                        self.union(next_x * length_y + next_y, x * length_y + y)
        for x in range(length_x):
            for y in range(length_y):
                if board[x][y] == 'O' and not self.same(x * length_y + y,
                                                        dummy):
                    board[x][y] = 'X'


def dfs(board, x, y, margin_x, margin_y):
    if x >= margin_x or x < 0 or y >= margin_y or y < 0 or board[x][y] != 'O':
        return
    board[x][y] = 'A'
    dfs(board, x - 1, y, margin_x, margin_y)
    dfs(board, x + 1, y, margin_x, margin_y)
    dfs(board, x, y - 1, margin_x, margin_y)
    dfs(board, x, y + 1, margin_x, margin_y)


def solve(board):
    margin_x = len(board)
    if margin_x <= 1:
        return
    margin_y = len(board[0])
    if margin_y <= 1:
        return
    for i in range(margin_x):
        if board[i][0] == 'O':
            dfs(board, i, 0, margin_x, margin_y)
        if board[i][margin_y-1] == 'O':
            dfs(board, i, margin_y-1, margin_x, margin_y)
    for i in range(1, margin_y-1):
        if board[0][i] == 'O':
            dfs(board, 0, i, margin_x, margin_y)
        if board[margin_x-1][i] == 'O':
            dfs(board, margin_x-1, i, margin_x, margin_y)
    for i in range(margin_x):
        for j in range(margin_y):
            if board[i][j] == 'O':
                board[i][j] = 'X'
    for i in range(margin_x):
        for j in range(margin_y):
            if board[i][j] == 'A':
                board[i][j] = 'O'

if __name__ == '__main__':
    s = Solution()
    board = [["X", "X", "X", "X"],
             ["X", "O", "O", "X"],
             ["X", "X", "O", "X"],
             ["O", "X", "X", "X"]]

    # s.solve(board)
    solve(board)
    for line in board:
        print(line)
