class Solution:
    def numIslands(self, grid):
        def dfs(grid, x, y):
            if x < 0 or x >= len(grid) \
                    or y < 0 or y >= len(grid[0]) \
                    or grid[x][y] == '0':
                return
            grid[x][y] = '0'
            dfs(grid, x - 1, y)
            dfs(grid, x + 1, y)
            dfs(grid, x, y + 1)
            dfs(grid, x, y - 1)

        nums = 0
        for x in range(len(grid)):
            for y in range(len(grid[0])):
                if grid[x][y] == '1':
                    nums += 1
                dfs(grid, x, y)
        return nums


print(Solution().numIslands(
    [["1", "1", "0", "0", "0"],
     ["1", "1", "0", "0", "0"],
     ["0", "0", "1", "0", "0"],
     ["0", "0", "0", "1", "1"]]))
