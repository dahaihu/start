package anything

/**
* @Author: 胡大海
* @Date: 2019-10-10 09:31
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

// 对于go来说，为组建方法提供方法，所以比较简陋
// 这个里面算是学习了，参数传递的时候，可以传递函数
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string {
	ret := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			ret = append(ret, v)
		}
	}
	return ret
}

func Map(vs []string, f func(string) string) []string {
	ret := make([]string, len(vs))
	for i, v := range vs {
		ret[i] = f(v)
	}
	return ret
}
