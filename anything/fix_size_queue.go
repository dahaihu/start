package anything

import "fmt"

type fixedSizeQueue struct {
	queue []interface{}
	start int
	end   int
	size  int
	cap   int
}

func NewFixSizeQueue(cap int) *fixedSizeQueue {
	var q fixedSizeQueue
	q.queue = make([]interface{}, cap, cap)
	q.cap = cap
	return &q
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

func (q *fixedSizeQueue) pop() (ele interface{}, ok bool) {
	if q.length() == 0 {
		return
	}
	ele = q.queue[q.start]
	q.start += 1
	if q.start >= len(q.queue) {
		q.start = 0
	}
	q.size -= 1
	ok = true
	return
}

func (q *fixedSizeQueue) String() string {
	return fmt.Sprintf("queue not is at (%d, %d)", q.start, q.end)
}
