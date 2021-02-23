class Solution:
    def __init__(self):
        self.mark = dict()
        self.nums = 0

    def find(self, x):
        if x not in self.mark:
            self.mark[x] = x
            self.nums += 1
            return x
        while self.mark[x] != x:
            x = self.mark[x]
        return x

    def union(self, x, y):
        x_ancestor, y_ancestor = self.find(x), self.find(y)
        if x_ancestor != y_ancestor:
            self.nums -= 1
            self.mark[x_ancestor] = y_ancestor

    def findCircleNum(self, isConnected):
        for i in range(len(isConnected)):
            for j in range(len(isConnected[0])):
                if isConnected[i][j]:
                    if i == j:
                        self.find(i)
                    else:
                        self.union(i, j)
        return self.nums


if __name__ == '__main__':
    s = Solution()
    print(s.findCircleNum([[1, 0, 0], [0, 1, 0], [0, 0, 1]]))
