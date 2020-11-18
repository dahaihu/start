package leetcode

/**
* @Author: 胡大海
* @Date: 2020-11-18 14:49
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

func existValInMultiArray(array [][]int, target int) bool {
	xIdx, yIdx := 0, len(array)-1
	for xIdx <= len(array) - 1 && yIdx >= 0 {
		if array[xIdx][yIdx] == target {
			return true
		} else if array[xIdx][yIdx] > target {
			yIdx -= 1
		} else {
			xIdx += 1
		}
	}
	return false
}
