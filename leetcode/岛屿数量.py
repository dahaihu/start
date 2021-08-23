class Solution:
    def numIslands(self, grid):
        nums = 0
        for i in range(len(grid)):
            for j in range(len(grid[0])):
                if grid[i][j] == '1':
                    nums += 1
                    self.traverse(grid, i, j)
        return nums

    def traverse(self, grid, i, j):
        if i < 0 or i >= len(grid) \
                or j < 0 or j >= len(grid[0]) \
                or grid[i][j] == '0':
            return
        grid[i][j] = '0'
        self.traverse(grid, i + 1, j)
        self.traverse(grid, i - 1, j)
        self.traverse(grid, i, j + 1)
        self.traverse(grid, i, j - 1)


if __name__ == '__main__':
    print(Solution().numIslands(
        [["1", "1", "0", "0", "0"],
         ["1", "1", "1", "0", "0"],
         ["0", "0", "0", "1", "0"],
         ["0", "0", "1", "1", "1"]]))
