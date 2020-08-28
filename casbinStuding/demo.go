package casbinStuding

import (
	"fmt"
	"github.com/casbin/casbin"
)

func demo() {
	e, err := casbin.NewEnforcer("auth_model.conf", "policy.csv")
	fmt.Println("initiate err is ", err)
	if err == nil {
		ok ,_ := e.Enforce("hushichang", "/login", "admin")
		fmt.Println("ok is ", ok)
	}
	fmt.Println(e.GetRolesForUser("hushichang"))
	fmt.Println(e.GetImplicitPermissionsForUser("hushichang"))
	//fmt.Println(e.GetRolesForUser(""))
}