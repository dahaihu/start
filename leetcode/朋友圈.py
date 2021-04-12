class Solution:
    def __init__(self):
        self.mark = dict()
        self.nums = 0

    def find(self, x):
        if x not in self.mark:
            self.nums += 1
            self.mark[x] = x
            return x
        while self.mark[x] != x:
            x = self.mark[x]
        return x

    def union(self, x, y):
        x_ancestor = self.find(x)
        y_ancestor = self.find(y)
        if x_ancestor != y_ancestor:
            self.mark[y_ancestor] = x_ancestor
            self.nums -= 1

    def findCircleNum(self, isConnected):
        for x in range(len(isConnected)):
            for y in range(len(isConnected[0])):
                if isConnected[x][y] == 1:
                    self.union(x, y)
        return self.nums


if __name__ == '__main__':
    print(Solution().findCircleNum([[1, 0, 0], [0, 1, 0], [0, 0, 1]]))
