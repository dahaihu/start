package leetcode

import (
	"fmt"
	"testing"
)

func TestExist(t *testing.T) {
	fmt.Println(exist([][]byte{
	{'A','B','C','E'},
	{'S','F','C','S'},
	{'A','D','E','E'}}, "EEDA"))

}