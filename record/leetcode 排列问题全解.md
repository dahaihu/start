# leetcode 排列问题全解

### 排列问题
1. 全排列
```go
func Permute(nums []int) [][]int {
	results := make([][]int, 0)
	var _permute func([]int, []int)
	_permute = func(nums []int, internalResult []int) {
		if len(nums) == 1 {
			internalResult = append(internalResult, nums[0])
			results = append(results, internalResult)
			return
		}
		for i := 0; i < len(nums); i++ {
			tmpInternalResult := make([]int, len(internalResult))
			copy(tmpInternalResult, internalResult)
			tmpInternalResult = append(tmpInternalResult, nums[i])
			_permute(popItem(nums, i), tmpInternalResult)
		}
	}
	internalResult := make([]int, 0)
	_permute(nums, internalResult)
	return results
}

func popItem(nums []int, idx int) []int {
	result := make([]int, idx)
	copy(result, nums[:idx])
	result = append(result, nums[idx+1:]...)
	return result
}
```
2. 全排列2
```go
func PermuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	results := make([][]int, 0)
	var _permute func([]int, []int)
	_permute = func(nums []int, internalResult []int) {
		if len(nums) == 1 {
			internalResult = append(internalResult, nums[0])
			results = append(results, internalResult)
			return
		}
		for i := 0; i < len(nums); i++ {
			if i >= 1 && nums[i] == nums[i-1] {
				continue
			}
			tmpInternalResult := make([]int, len(internalResult))
			copy(tmpInternalResult, internalResult)
			tmpInternalResult = append(tmpInternalResult, nums[i])
			_permute(popItem(nums, i), tmpInternalResult)
		}
	}
	internalResult := make([]int, 0)
	_permute(nums, internalResult)
	return results
}
```
3. 下一个全排列
```go
func nextPermutation(nums []int)  {
	margin := len(nums)
	for i:=len(nums)-1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			margin = i
			break
		}
	}
	if margin == len(nums) {
		swap(nums, 0, len(nums)-1)
		return
	}

	for idx := len(nums)-1; idx >= margin; idx-- {
		if nums[idx] > nums[margin-1] {
			nums[margin-1], nums[idx] = nums[idx], nums[margin-1]
			swap(nums, margin, len(nums)-1)
			return
		}
	}
}

func swap(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start += 1
		end -= 1
	}
}
```