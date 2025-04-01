package HeapSort

// 小根堆
type Heap []int

func (h Heap) Len() int {
	return len(h)
}
func (h Heap) Init() {
	// 仅通过Push无法保证堆结构
	// 来一组数据，不能通过不断朝初始长度为0的切片push进行初始化
	// 应该调用该Init函数
	l := h.Len()
	for i := l / 2; i >= 0; i-- {
		h.Down(i)
	}
}

func (h *Heap) Push(val int) {
	(*h) = append((*h), val)
	(*h).Up(h.Len() - 1)
}

func (h *Heap) Pop() int {
	res := (*h)[0]
	(*h)[0] = (*h)[h.Len()-1]
	(*h) = (*h)[:h.Len()-1]
	h.Down(0)
	return res
}

func (h Heap) Down(i int) {
	n := h.Len()
	for {
		l, r := 2*i+1, 2*i+2
		small := i
		for l < n && h[l] < h[small] {
			small = l
		}
		for r < n && h[r] < h[small] {
			small = r
		}
		if small == i {
			break
		}
		h[small], h[i] = h[i], h[small]
		i = small
	}
}

func (h Heap) Up(i int) {
	for {
		parent := (i - 1) / 2
		if i == 0 || h[parent] <= h[i] {
			break
		}
		h[parent], h[i] = h[i], h[parent]
		i = parent
	}
}
