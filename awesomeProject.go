package main

import (
	"awesomeProject/anything"
	"awesomeProject/playball"
	"awesomeProject/runner"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"time"
)

type Contact struct {
	Name    string `json:"name"`
	Title   string `json:"title,omitempty"`
	Contact struct {
		Home string `json:"home"`
		Cell string `json:"cell"`
	} `json:"contact"`
}

var JSON = `{
	"name":"Gopher",
	"title":"",
	"contact":{
		"home":"123123123",
		"cell":"456456456"
		}
	}`

func hehe() {
	log.Println("Starting work.")
	r := runner.New(3 * time.Second)
	r.Add(runner.CreateTask(), runner.CreateTask(), runner.CreateTask())
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("terminating due to timeout")
		case runner.ErrInterrupt:
			log.Println("terminating due to interrupt")

		}
	}
	a := map[string]string{"hushichang": "hudahai"}
	fmt.Println("invalid result is ", a["hudahai"])
	array := [3]*string{new(string), new(string), new(string)}
	*array[0] = "Red"
	*array[1] = "Blue"
	*array[2] = "Green"
	for _, value := range array {
		fmt.Println(*value)
	}
	// new(T)为一个T类型的空间，并将空间初始化T的零值
	p1 := new(int)
	fmt.Println(p1)
	fmt.Println(*p1)
	// 接码json
	var c Contact
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}
	fmt.Println(c.Contact)
	fmt.Println(c.Title)
	fmt.Println(c.Name)

	// 编码json
	// 类型Contact中的title的设置为 `json:"title,omitempty"`的作用出来了
	// 编码为json的时候，没有这个字段
	data, _ := json.Marshal(c)
	fmt.Println(string(data))

	// 不知道数据结构的时候解析json
	var tmp map[string]interface{}
	_ = json.Unmarshal([]byte(JSON), &tmp)
	fmt.Println("Name:", tmp["name"])
	fmt.Println("Title:", tmp["title"])
	fmt.Println("Contact:", tmp["contact"])
	fmt.Println("Home:", tmp["contact"].(map[string]interface{})["home"])
	fmt.Println("Cell:", tmp["contact"].(map[string]interface{})["cell"])

}

func f() {
	for skip := 0; ; skip++ {
		pc, file, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}
		fmt.Printf("%v\n", filepath.Base(runtime.FuncForPC(pc).Name()))
		fmt.Printf("%v\n", filepath.Base(file))
		fmt.Printf("%v\n", runtime.FuncForPC(pc).Name())
		fmt.Printf("skip=%v, pc=%v, file=%v, line=%v\n", skip, pc, file, line)
	}
}

func testFindStringSubmatch() {
	text := "Hello 世界！123 Go."
	reg := regexp.MustCompile("([a-z]+)")
	// findStringSubmatch是用来找一个一个匹配的pattern，并把其中分组的信息也放在里面一起处理
	fmt.Printf("%q\n", reg.FindStringSubmatch(text))
}

func passRightToLeft(left, right chan int) {
	left <- 1 + <-right
}

func daisyChain() {
	const n = 10000
	leftmost := make(chan int)
	right := leftmost
	left := leftmost
	for i := 0; i < n; i++ {
		right = make(chan int)
		go passRightToLeft(left, right)
		left = right
	}
	go func(c chan int) { c <- 1 }(right)
	fmt.Println(<-leftmost)

}

func testBuffer() {
	var buf bytes.Buffer
	var kv = map[string]string{"1": "2", "2": "3"}
	for key, value := range kv {
		buf.WriteString(key)
		buf.WriteString("=")
		buf.WriteString(value)
	}
	fmt.Println("buf is %v\n", string(buf.Bytes()))
	fmt.Println("buf is %v\n", buf.Bytes())
}

func boring(msg string) <-chan string { // Returns receive-only channel of strings. 我觉得这个返回的应该是一个只out-only of strings
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function.
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			//<- time.After(1 * time.Second)
		}
	}()
	return c // Return the channel to the caller.
}

type QueryArgs map[string]interface{}

func (q QueryArgs) Condition() (sql string, args []interface{}) {
	builder := strings.Builder{}
	cnt := 0
	for k, v := range q {
		if cnt == 0 {
			builder.WriteString("where ")
		} else {
			builder.WriteString(" and ")
		}
		cnt += 1

		builder.WriteString(fmt.Sprintf("`%s`=", k))
		args = append(args, v)
		switch v.(type) {
		case int64:
			builder.WriteString("?")
		case string:
			builder.WriteString("'?'")
		default:
			builder.WriteString("?")
		}
	}
	return builder.String(), args
}

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) ReflectCallfunc() {
	fmt.Println("Allen.Wu ReflectCallFunc")
}

func DoFieldAndMethod(input interface{}) {
	//获取输入参数的类型
	getType := reflect.TypeOf(input)
	fmt.Println("get Type is :", getType.Name())
	getValue := reflect.ValueOf(input)
	fmt.Println("get all fields is:", getValue)
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("%s: %v=%v\n", field.Name, field.Type, value)
	}

}

func ttest(names ...string) {
	fmt.Println(names)
}
func test(names ...string) string {
	ttest(names...)
	return strings.Join(names, " ")
}
func main() {
	anything.SelectExp()
}

func testFallin() {
	c := fallin(boring("joe"), boring("ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

func fallin(input1, input2 <-chan string) chan string {
	c := make(chan string)
	//go func() { for {c <- <-input1} }()
	//go func() { for {c <- <-input2} }()
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

func testRunner() {
	log.Println("Starting work.")
	r := runner.New(7 * time.Second)
	r.Add(runner.CreateTask(), runner.CreateTask(), runner.CreateTask(), runner.CreateTask())
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("terminating due to timeout")
		case runner.ErrInterrupt:
			log.Println("terminating due to interrupt")

		}
	}
}

func testPlayball() {
	playball.Wg.Add(2)
	court := make(chan int)
	go playball.Play("胡世昌", court)
	go playball.Play("李漫漫", court)
	court <- 1
	playball.Wg.Wait()

}
