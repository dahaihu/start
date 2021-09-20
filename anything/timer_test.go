package anything

import (
	"fmt"
	"testing"
	"time"
)

func TestTimerTest(t *testing.T) {
	TimerTest()

	now := time.Now()

	last := now.Add(time.Second)
	fmt.Println("时间差是：", now.Sub(last))
}
