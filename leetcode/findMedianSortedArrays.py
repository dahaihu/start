class Solution:
    def findMedianSortedArrays(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: float
        """
        length1, length2 = len(nums1), len(nums2)
        idx, remainder = divmod(length1 + length2, 2)
        if remainder == 1:
            return self.findKth(nums1, 0, length1 - 1, nums2, 0, length2 - 1, idx + 1)
        return (self.findKth(nums1, 0, length1 - 1, nums2, 0, length2 - 1, idx) + self.findKth(nums1, 0, length1 - 1,
                                                                                               nums2, 0, length2 - 1,
                                                                                               idx + 1)) / 2

    def findKth(self, nums1, start1, end1, nums2, start2, end2, k):
        length1 = end1 - start1 + 1
        length2 = end2 - start2 + 1
        if length1 > length2:
            return self.findKth(nums2, start2, end2, nums1, start1, end1, k)
        if length1 == 0:
            return nums2[start2 + k - 1]
        if k == 1:
            return min(nums1[start1], nums2[start2])
        idx1 = start1 + min(k // 2, length1) - 1
        idx2 = start2 + min(k // 2, length2) - 1
        if nums1[idx1] > nums2[idx2]:
            return self.findKth(nums1, start1, end1, nums2, idx2 + 1, end2, k - (idx2 - start2 + 1))
        return self.findKth(nums1, idx1 + 1, end1, nums2, start2, end2, k - (idx1 - start1 + 1))


class Solutions:
    def findMedianSortedArrays(self, nums1, nums2) -> float:
        n1 = len(nums1)
        n2 = len(nums2)
        if n1 > n2:
            return self.findMedianSortedArrays(nums2, nums1)
        k = (n1 + n2 + 1) // 2
        left = 0
        right = n1
        while left < right:
            m1 = left + (right - left) // 2
            m2 = k - m1
            print('m1 is {}, m2 is {} '.format(m1, m2))
            print('left is {}, right is {} '.format(left, right))
            if nums1[m1] < nums2[m2 - 1]:
                left = m1 + 1
            else:
                right = m1
        print('left is ', left)
        m1 = left
        m2 = k - m1
        c1 = max(nums1[m1 - 1] if m1 > 0 else float("-inf"), nums2[m2 - 1] if m2 > 0 else float("-inf"))
        if (n1 + n2) % 2 == 1:
            return c1
        c2 = min(nums1[m1] if m1 < n1 else float("inf"), nums2[m2] if m2 < n2 else float("inf"))
        return (c1 + c2) / 2

if __name__ == '__main__':
    nums1 = [1, 2, 10, 13]
    nums2 = [5, 6, 7, 8]
    s = Solutions()
    print(s.findMedianSortedArrays(nums1, nums2))