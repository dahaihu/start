package anything

import (
	"fmt"
	"strconv"
)

/**
* @Author: 胡大海
* @Date: 2019-09-25 13:48
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

func TestStrconv() {
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	fmt.Println(sample)
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%v %x\n", sample[i], sample[i])
	}
	fmt.Println()
	fmt.Println(strconv.Atoi("123123"))
}
