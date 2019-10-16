package anything

import "fmt"

/**
* @Author: 胡大海
* @Date: 2019-10-09 09:53
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

func ArrayExp() {
	var a [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			a[i][j] = i + j
		}
	}
	fmt.Println(a)
}
