class Solution:
    def numIslands(self, grid):
        def dfs(grid, x, y):
            if x < 0 or x >= len(grid) \
                    or y < 0 or y >= len(grid[0]) \
                    or grid[x][y] == '0':
                return 0
            grid[x][y] = '0'
            return 1 + \
                   dfs(grid, x - 1, y) + \
                   dfs(grid, x + 1, y) + \
                   dfs(grid, x, y + 1) + \
                   dfs(grid, x, y - 1)

        res = 0
        for x in range(len(grid)):
            for y in range(len(grid[0])):
                res = max(res, dfs(grid, x, y))
        return res


print(Solution().numIslands(
    [["1", "1", "0", "0", "0"],
     ["1", "1", "0", "0", "0"],
     ["0", "0", "1", "1", "0"],
     ["0", "0", "1", "1", "1"]]))
