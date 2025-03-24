package queue

type Queue []int //FIFO
func (q *Queue) Push(v int) {
	(*q) = append((*q), v)
}

func (q *Queue) Pop() int {
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}
