package anything

import (
	"testing"
	"time"
)

func TestAutoLoad(t *testing.T) {
	autoLoad(time.Second)
	time.Sleep(time.Second * 100)
}
