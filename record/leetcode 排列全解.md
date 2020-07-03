# leetcode 排列全解

这次题解有四题，全是排列相关的。

## 全排列

### 原题

给定一个 **没有重复** 数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

### 题解

列表 A 用于保存一个可能的排列的部分值，A 的初始值是一个空列表。列表 B 用于保存列表 A 还可以选择的值，B 的初始值是题目中的输入。在 B 中还有值得时候，A 的下一个取值就可以是 B 中的任何一个，这个在代码里面就可以表现为顺序选择 B 的每一个元素。结束的时机就在数组 B 中的元素个数为 1 的时候，把 B 中剩下的一个唯一的元素添加到 A 中就好了，此时的 A 就是一个可能的排列。

剩下的就是扣细节了。由于切片的底层数组公用，所以在 A 从 B 中任意选择一个数字并添加到 A 之前，都需要创建一个新的然后再添加到其中，避免相互影响。

全部代码如下

```go
func Permute(nums []int) [][]int {
	results := make([][]int, 0)
	var _permute func([]int, []int)
	_permute = func(nums []int, internalResult []int) {
         // 结束的时机
		if len(nums) == 1 {
			internalResult = append(internalResult, nums[0])
			results = append(results, internalResult)
			return
		}
		for i := 0; i < len(nums); i++ {
             // 创建新的
			tmpInternalResult := make([]int, len(internalResult))
			copy(tmpInternalResult, internalResult)
             // 把选择的数字加入到新的
			tmpInternalResult = append(tmpInternalResult, nums[i])
             // 此时 A 和 B 已经变化了，A 中添加了一个元素， B 中减少了一个元素
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

##  全排列 II

### 原题

给定一个可包含重复数字的序列，返回所有不重复的全排列。

示例:

输入: [1,1,2]
输出:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

### 题解

此题和上面的题目大同小异，唯一的区分就是如何分辨排列是重复的。如果此时对于数组 B 中待选择列表为`[b, b, ...]`也就是数组 A 从 B 中选择了第一个元素`b`之后，还可以选择`b`吗？答案是否定的，因为选择了第一个`b`之后产生的所有可能排列，和选择了第二个`b`之后产生的所有排列是完全相同的，exp: 选择了第一个`b`之后，A 是[a1, a2, ... , b]，B 是[b, ....]，选择了第二个`b`的时候，A  还是[a1, a2, ... , b]，B 也还是[b, ....]，两次产生的排列会重复一遍。所以在选择了下一个元素和前一个元素相同的时候，应该过滤掉。

当然，相同大小的元素位置相互挨着的条件是数组是排序过的，所以操作前要对数组进行排序。这个不是考点，用第三方提供的工具来操作就好了。

```go
import "sort"
func permuteUnique(nums []int) [][]int {
     // 排序
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
             // 过滤重复的排列
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


func popItem(nums []int, idx int) []int {
	result := make([]int, idx)
	copy(result, nums[:idx])
	result = append(result, nums[idx+1:]...)
	return result
}
```

## 下一个全排列

### 原题


实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须原地修改，只允许使用额外常数空间。

以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/next-permutation
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

### 题解

如何找到下一个比此排列大的排列呢？如果仅仅思考这个问题是不好思考到结果的，而如果从实例来出发，思路比较好的找出了。我找了一个比较好的例子, `1,3,2`→ `2,1,3`。

显然，1开头的排列排完了，如果还是以 1 开头，后面的`3,2`无论怎么变化，整体的结果都不会比`1,3,2`更大。所以需要换开头，哪一个开头比较好，是选择`2`还是选择`3`？肯定是选择`2`，因为选择`3`了，确实之前的排列大，但是大的太多了，两者之间还有其他的可能的排列。选择`2`之后就完了吗？肯定不是的呀，剩下的 `1` 和 `3`应该怎么放呀？你可能会说屁话，当然是`3`放在`1`后面，你是对的！

其实整体思路就是和上面的一样，区别就是牵扯的数字变多了，唯一的例外是列表中数字的顺序是从大往小排列的，此时就是最大的一个排列，那么久需要把列表逆序排列下就好了。

剩下的我也不知道怎么写了，就把官方的解题里面的一个图贴下来给大家看看吧

![ Next Permutation ](https://pic.leetcode-cn.com/dd4e79b184b1922429d8cda6148a3f0b7579869e85626e04ba29ba88e8052729-file_1555696116786)

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

## 第 k 个排列

### 原题

给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。

按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：

"123"
"132"
"213"
"231"
"312"
"321"
给定 n 和 k，返回第 k 个排列。

说明：

给定 n 的范围是 [1, 9]。
给定 k 的范围是[1,  n!]。
示例 1:

输入: n = 3, k = 3
输出: "213"
示例 2:

输入: n = 4, k = 9
输出: "2314"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutation-sequence
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

### 题解

总的思路来说，也是如何从列表 B`[1, ..., n]`中一个一个的选择数字，来组成第 K 个排列。

`n! = n * (n-1)!`，这个公式比较重要，在前面列出来，含义就是：列表`[1, ..., n]`中的`n`个数字，每个数字开头的排列都有`(n-1)!`个。

对于要找第 K 大的数，那么如何从`[1, ..., n]`里面选择第一个数字呢？可以通过 `idx, k = k/(n-1)!, k%(n-1)!`来选择第一个数，如果余数`k`为0， 表示的是 `k`是`(n-1)!`的整数倍`k = idx * (n-1)!`，此时在列表中`[1, ..., n]`应该选择的是数字的索引就是`idx-1`，而第`idx-1`个数开头的最大的排列，就需要把剩下的数字逆序的加入到结果之中即可，因为数组`B`中的第`idx`个数组成的最大排列就是第`idx * (n-1)!`个排列。如果不是0，那么应该选择的数字的索引就是`idx`了，因为前`idx * (n-1)!`个数不够，如果余数是 1， 则表示的是选择完`idx`这个值之后，把列表 B 中的第一个排列加入到结果就好了，如果不是则需要进行递归的操作。

```go
func GetPermutation(n int, k int) string {
	res := bytes.Buffer{}
	mark := make([]int, n)
    // 把 1 到 n-1 的每个阶乘存入到一个切片中，避免每次都要重新计算
	mark[0] = 1
	for i := 1; i < n; i++ {
		mark[i] = mark[i-1] * i
	}
	items := make([]int, n)
	for i := 0; i < n; i++ {
		items[i] = i + 1
	}
	var idx, item int
	for true {
		idx, k = k/mark[len(items)-1], k%mark[len(items)-1]
		if k == 0 {
			items, item = popKItem(items, idx-1)
			res.WriteString(strconv.Itoa(item))
			for i := len(items) - 1; i >= 0; i-- {
				res.WriteString(strconv.Itoa(items[i]))
			}
			break
		}
		items, item = popKItem(items, idx)
		res.WriteString(strconv.Itoa(item))
		if k == 1 {
			for i := 0; i < len(items); i++ {
				res.WriteString(strconv.Itoa(items[i]))
			}
			break
		}

	}
	return res.String()
}
func popKItem(items []int, idx int) ([]int, int) {
	item := items[idx]
	newItems := make([]int, idx)
	copy(newItems, items[:idx])
	newItems = append(newItems, items[idx+1:]...)
	return newItems, item
}

```

