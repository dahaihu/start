package anything

import "testing"

/**
* @Author: 胡大海
* @Date: 2020-10-04 10:21
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */


func TestRWLockExp(t *testing.T) {
	for i := 0; i <= 10; i++ {
		RWLockExp()
	}
}
