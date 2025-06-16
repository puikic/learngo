package main

import (
	"bufio"
	"fmt"
	"learngo/src/DS_ALG/HeapSort"
	"os"
	"strconv"
	"strings"
)

func main() {
	//ACM模式
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	n, _ := strconv.Atoi(line)
	h := make(HeapSort.Heap, n)
	scanner.Scan()
	line = scanner.Text()
	nums := strings.Fields(line)
	for i := range nums {
		h[i], _ = strconv.Atoi(nums[i])
	}
	h.Init()
	for i := 0; i < n; i++ {
		fmt.Println(h.Pop())
	}
}
