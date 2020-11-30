class Solution(object):
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        res = 0
        mark = ""
        for ele in s:
            if ele not in mark:
                mark = mark + ele
            else:
                res = max(res, len(mark))
                idx = mark.index(ele)
                mark = mark[idx + 1:] + ele
        res = max(res, len(mark))
        return res


if __name__ == '__main__':
    a = [1, 2, 3]
    b = a.copy()
    print(a.pop())
    print(a)
    print(b)
