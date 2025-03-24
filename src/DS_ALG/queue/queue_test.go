package queue

import "testing"

func TestQueue1(t *testing.T) {
	q := &Queue{}
	q.Push(3)
	q.Push(4)
	q.Push(5)
	//3 45
	x := q.Pop()
	t.Log(q, x)
}
