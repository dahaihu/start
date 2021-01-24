package casbinStuding

import (
	"fmt"
	"github.com/casbin/casbin"
)

func demo() {
	e:= casbin.NewEnforcer("auth_model.conf", "policy.csv")
	ok := e.Enforce("bob", "data2", "write")
	fmt.Println("ok is ", ok)
}