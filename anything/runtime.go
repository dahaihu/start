package anything

import (
	"fmt"
	"path/filepath"
	"runtime"
)

/**
* @Author: 胡大海
* @Date: 2019-10-15 11:01
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */


func RuntimeCallerExp() {
	for skip := 0; ; skip++ {
		pc, file, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}
		file = filepath.Base(file)
		fmt.Printf("file is %s \n", file)
		f := runtime.FuncForPC(pc)
		fmt.Printf("aaa is %s\n", filepath.Base(f.Name()))
		fmt.Printf("skip = %v, pc = %v, file = %v, line = %v, funcname = %s, f.entry() = %v \n", skip, pc, file, line, f.Name(), f.Entry())
	}
}
