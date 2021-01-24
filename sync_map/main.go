package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"unsafe"
)

func syncMapFunc() {
	m := sync.Map{}
	m.Store("name", "zhangsan")
}

func loadPointer() {
	var a int64 = 0
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			p := unsafe.Pointer(&a)
			fmt.Println(*(*int64)(atomic.LoadPointer(&p)))
		}()
	}
	wg.Wait()
}

//type entry struct {
//	p unsafe.Pointer
//}

//func (e *entry) tryStore(i *interface{}) bool {
//	for {
//		p := atomic.LoadPointer(&e.p)
//		if atomic.CompareAndSwapPointer(&e.p, p, unsafe.Pointer(i)) {
//			return true
//		}
//	}
//}

//func newEntry(a *interface{}) *entry {
//	return &entry{p: unsafe.Pointer(a)}
//}

type entryMap struct {
	v map[interface{}]*entry
}

type syncMap struct {
	val atomic.Value
}

func (m *syncMap) Store(key, value interface{}) {
	eM := m.val.Load().(entryMap)
	if v, ok := eM.v[key]; ok {
		v.tryStore(&value)
	} else {
		fmt.Printf("%v not in map\n", key)
	}
}

func atomicValue() {
	var (
		a  atomic.Value
		wg sync.WaitGroup
	)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			a.Store(v)
		}(i)
	}
	wg.Wait()
	fmt.Println("loaded value", a.Load().(int))
}

type concurrentMap struct {
	m  map[string]string
	mx sync.RWMutex // or mx sync.Mutex
}

func runMap() {

}

func main() {
	m := Map{}
	a := 10
	m.Store("age", a)
	m.Store("age", 1)
	fmt.Println("a is ", a)
}
