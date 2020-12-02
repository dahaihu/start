package cache

import (
	"container/list"
	"fmt"
	"github.com/golang/groupcache/consistenthash"
	"sort"
	"strconv"
	"testing"
)

func TestCache(t *testing.T) {
	cacheExap()
}

func TestListExamp(t *testing.T) {
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func TestStringAdd(t *testing.T) {
	fmt.Println("123" + "234")
}


func TestStudingSingleFlight(t *testing.T) {
	StudyingSingleFlight()
}

func TestBinarySearch(t *testing.T) {
	mark := []int{2, 5, 7}
	fmt.Println(sort.Search(len(mark), func(i int) bool {return mark[i] >= 6}))
}

func TestConsistency(t *testing.T) {
	hash1 := consistenthash.New(2, nil)
	hash1.Add("www.baidu.com", "www.google.com")
	for i := 0; i < 10; i++ {
		fmt.Println(hash1.Get(strconv.Itoa(i)))
	}

}