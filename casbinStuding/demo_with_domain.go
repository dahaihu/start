package casbinStuding

import (
	"fmt"
	"github.com/casbin/casbin"
)

func demoWithDomain() {
	e, err := casbin.NewEnforcer("auth_model_with_domain.conf", "policy_with_domain.csv")
	fmt.Println("initiate err is ", err)
	if err == nil {
		ok ,_ := e.Enforce("hushichang", "zhihu", "/login", "admin")
		fmt.Println("ok is ", ok)
	}
	fmt.Println(e.GetRolesForUser("hushichang", "zhihu"))
	fmt.Println(e.GetImplicitPermissionsForUser("hushichang", "zhihu"))
	//fmt.Println(e.GetRolesForUser(""))
}