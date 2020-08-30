package casbinStuding

import (
	"fmt"
	"github.com/casbin/casbin"
)

func demoWithDomain() {
	e := casbin.NewEnforcer("auth_model_with_domain.conf", "policy_with_domain.csv")
	ok := e.Enforce("hushichang", "zhihu", "/login", "admin")
	fmt.Println("enforce result is ", ok)
	fmt.Println(e.GetImplicitRolesForUser("hushichang", "zhihu"))
	fmt.Println(e.GetImplicitPermissionsForUser("hushichang", "zhihu"))
	//fmt.Println(e.GetRolesForUser(""))
}