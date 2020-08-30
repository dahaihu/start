package casbinStuding

import (
	"fmt"
	"github.com/casbin/casbin"
)

func demo() {
	e:= casbin.NewEnforcer("auth_model.conf", "policy.csv")
	ok := e.Enforce("hushichang", "/login", "admin")
	fmt.Println("ok is ", ok)
	fmt.Println(e.GetRolesForUser("hushichang"))
	fmt.Println(e.GetImplicitPermissionsForUser("hushichang"))
	//fmt.Println(e.GetRolesForUser(""))
}