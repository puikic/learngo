package queue

import "testing"

func TestQueue1(t *testing.T) {
	q := &Queue{}
	q.Push(3)
	q.Push(4)
	q.Push(5)
	//3 45
	t.Log((*q)[q.Size()-1])
	x := q.Pop()
	t.Log(q, x)
	t.Log(q.isEmpty())
}
