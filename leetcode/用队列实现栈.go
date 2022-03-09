package leetcode

type Queue struct {
	nums []int
}

func (q *Queue) Push(num int) {
	q.nums = append(q.nums, num)
}

func (q *Queue) Peek() int {
	return q.nums[0]
}

func (q *Queue) Pop() int {
	num := q.nums[0]
	q.nums = q.nums[1:]
	return num
}

func (q *Queue) Size() int {
	return len(q.nums)
}

func (q *Queue) IsEmpty() bool {
	return q.Size() == 0
}

type MyStack struct {
	usedQueue Queue
	helpQueue Queue
}

func (this *MyStack) Push(x int) {
	this.helpQueue.Push(x)
	for this.usedQueue.Size() != 0 {
		num := this.usedQueue.Pop()
		this.helpQueue.Push(num)
	}
	this.usedQueue, this.helpQueue = this.helpQueue, this.usedQueue
}

func (this *MyStack) Pop() int {
	return this.usedQueue.Pop()
}

func (this *MyStack) Top() int {
	return this.usedQueue.Peek()
}

func (this *MyStack) Empty() bool {
	return this.usedQueue.IsEmpty()
}
