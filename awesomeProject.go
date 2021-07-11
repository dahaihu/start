package main

import "start/anything"

type Foo struct {
	A string
	B int64
}

type fixedSizeQueue struct {
	queue []interface{}
	start int
	end   int
	size  int
	cap   int
}

func (q *fixedSizeQueue) length() int {
	return q.size
}

func (q *fixedSizeQueue) capacity() int {
	return q.cap
}

func (q *fixedSizeQueue) add(val interface{}) (ok bool) {
	if q.length() == q.capacity() {
		return
	}
	q.queue[q.end] = val
	q.end += 1
	if q.end >= len(q.queue) {
		q.end = 0
	}
	q.size += 1
	ok = true
	return
}

func (q *fixedSizeQueue) pop() (interface{}, bool) {
	return nil, true
}

func main() {
	anything.Fallthrough(false)
}
