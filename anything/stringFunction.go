package anything

import (
	"fmt"
	"strings"
)

/**
* @Author: 胡大海
* @Date: 2019-10-10 09:45
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

func StringFunctionExp() {
	fmt.Println("contains: ", strings.Contains("hushichang", "shi"))
	fmt.Println("count: ", strings.Count("hushichang", "shi"))
	fmt.Println("split: ", strings.Split("hushichang", "shi"))
	fmt.Println("upper: ", strings.ToUpper("hushichang"))
	fmt.Println("repeat: ", strings.Repeat("hu ", 5))
}
