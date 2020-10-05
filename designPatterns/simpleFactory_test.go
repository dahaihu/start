package designPatterns

import "testing"

/**
* @Author: 胡大海
* @Date: 2020-10-05 11:35
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

func TestSimpleFactory(t *testing.T) {
	factory("product1").create()
	factory("product2").create()
	//factory("1").create()
}
