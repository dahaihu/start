package anything

import (
	"encoding/json"
	"fmt"
)

/**
* @Author: 胡大海
* @Date: 2020-02-09 21:16
* A programmer who subconsciously views himself as an artist will enjoy what he does and will do it better ​
 */

type Response1 struct {
	Page   int
	Fruits []string
}

type A struct {
	C string `json:"c,omitempty"`
}
type B struct {
	D string `json:"d,omitempty"`
}
type Response2 struct {
	A
	B
}

func NewResponse2() *Response2 {
	return &Response2{}
}

func NewResponse1() *Response1 {
	return new(Response1)
}

func TestJson() {
	response2 := NewResponse2()
	response2.C = "10"
	//response2.D = "10"
	d, _ := json.Marshal(response2)
	fmt.Println(string(d))
	resp2 := NewResponse2()
	json.Unmarshal([]byte(d), resp2)
	fmt.Println(resp2.C)
	fmt.Println(resp2.D == "")
}
