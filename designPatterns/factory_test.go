package designPatterns

import "testing"

/**
* @Author: 胡大海
* @Date: 2020-10-05 17:06
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

func TestFactory(t *testing.T) {
	gdF := GDBunShopFactory{}
	hbF := HBBunShopFactory{}
	gdF.Generate("baozi").create()
	hbF.Generate("baozi").create()

}
