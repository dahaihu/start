package cache

import (
	"container/list"
	"fmt"
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