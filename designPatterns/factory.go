package designPatterns

import "fmt"

/**
* @Author: 胡大海
* @Date: 2020-10-05 16:46
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

type BunShopInterface interface {
	Generate(t string) Bun
}

type Bun interface {
	create()
}

type GDBaozi struct {
}

func (baozi GDBaozi) create() {
	fmt.Println("create 广东包子")
}

type GDJiaozi struct{}

func (jiaozi GDJiaozi) create() {
	fmt.Println("create 广东饺子")
}

type HBBaozi struct {
}

func (baozi HBBaozi) create() {
	fmt.Println("create 湖北包子")
}

type HBJiaozi struct{}

func (jiaozi HBJiaozi) create() {
	fmt.Println("create 湖北饺子")
}

type GDBunShopFactory struct {
}

func (bunFactory GDBunShopFactory) Generate(s string) Bun {
	switch s {
	case "baozi":
		return GDBaozi{}
	case "jiaozi":
		return GDJiaozi{}
	default:
		panic("unknown type")
	}

}

type HBBunShopFactory struct {
}

func (bunFactory HBBunShopFactory) Generate(s string) Bun {
	switch s {
	case "baozi":
		return HBBaozi{}
	case "jiaozi":
		return HBJiaozi{}
	default:
		panic("unknown type")
	}
}
