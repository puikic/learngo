package quickSort

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	x := split(arr)
	quickSort(arr[0:x])
	quickSort(arr[x+1 : len(arr)])
}
func split(arr []int) int {
	x := arr[0]
	length := len(arr)
	l, r := 0, length-1
	for l < r {
		for arr[r] >= x && l < r {
			r--
		}
		arr[l] = arr[r]

		for arr[l] <= x && l < r {
			l++
		}
		arr[r] = arr[l]
	}
	arr[l] = x
	return l
}
