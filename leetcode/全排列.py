class Permute:
    def __init__(self):
        self.result = []

    def permute(self, nums):
        self._permute(nums, [])
        return self.result

    def _permute(self, nums, internal_result):
        if len(nums) == 1:
            self.result.append(internal_result + [nums[0]])
            return
        for i in range(len(nums)):
            self._permute(nums[:i] + nums[i + 1:], internal_result + [nums[i]])


class Solution:
    def __init__(self):
        self.result = []

    def permuteUnique(self, nums):
        nums.sort()
        self._permute(nums, [])
        return self.result

    def _permute(self, nums, internal_result):
        if len(nums) == 1:
            self.result.append(internal_result + [nums[0]])
            return
        for i in range(len(nums)):
            if i > 0 and nums[i] == nums[i - 1]:
                continue
            self._permute(nums[:i] + nums[i + 1:], internal_result + [nums[i]])


if __name__ == '__main__':
    nums = [1, 3, 1]
    nums.sort()
    print(nums)
