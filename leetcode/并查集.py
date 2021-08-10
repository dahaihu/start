class Solution:
    def solve(self, board):
        """
        Do not return anything, modify board in-place instead.
        """

        def find(x):
            if x not in mark:
                mark[x] = x
                return x
            if mark[x] != x:
                mark[x] = find(mark[x])
            return mark[x]

        def union(x, y):
            mark[find(x)] = find(y)

        mark = dict()
        m, n = len(board), len(board[0])

        dummy = m * n + 1
        for i in range(m):
            for j in range(n):
                if board[i][j] == 'X':
                    continue
                # the following is talking about 'O'
                if i == 0 or i == m - 1 or j == 0 or j == n - 1:
                    union(i * n + j, dummy)
                else:
                    for add_x, add_y in [(-1, 0), (1, 0), (0, -1), (0, 1)]:
                        if board[i + add_x][j + add_y] == 'O':
                            union(i * n + j, (i + add_x) * n + j + add_y)

        for i in range(m):
            for j in range(n):
                if board[i][j] == 'O' and find(i * n + j) != find(dummy):
                    board[i][j] = 'X'


def dfs(board, i, m, j, n):
    if i < 0 or i >= m or j < 0 or j >= n or board[i][j] != 'O':
        return
    board[i][j] = 'A'
    dfs(board, i - 1, m, j, n)
    dfs(board, i + 1, m, j, n)
    dfs(board, i, m, j - 1, n)
    dfs(board, i, m, j + 1, n)


def solve(board):
    if len(board) == 0 or len(board[0]) == 0:
        return
    m, n = len(board), len(board[0])
    for i in range(m):
        if board[i][0] == 'O':
            dfs(board, i, m, 0, n)


if __name__ == '__main__':
    s = Solution()
    board = [["O", "X", "X", "O", "X"],
             ["X", "O", "O", "X", "O"],
             ["X", "O", "X", "O", "X"],
             ["O", "X", "O", "O", "O"],
             ["X", "X", "O", "X", "O"]]

    s.solve(board)
    for line in board:
        print(line)
