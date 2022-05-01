class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        lookup = set()
        left = 0
        curLen, maxLen = 0, 0
        for ele in s:
            while ele in lookup:
                lookup.remove(s[left])
                left += 1
            lookup.add(ele)
            curLen = len(lookup)
            if curLen > maxLen:
                maxLen = curLen
        return maxLen


if __name__ == '__main__':
    print(Solution().lengthOfLongestSubstring( "abcabcbb"))