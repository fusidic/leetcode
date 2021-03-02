package array

func getLeastNumbers(arr []int, k int) []int {
	if arr == nil || k == 0 {
		return []int{}
	}
	left := 0
	right := len(arr) - 1
	idx := partition(arr, left, right)
	for idx != k-1 {
		if idx > k-1 {
			right = idx - 1
		} else {
			left = idx + 1
		}
		idx = partition(arr, left, right)
	}
	return arr[:idx+1]
}

func partition(arr []int, left, right int) int {
	base := arr[left]
	l := left + 1
	r := right
	for true {
		for l < right && arr[l] <= base {
			l++
		}
		for r > left && arr[r] >= base {
			r--
		}
		if l >= r {
			break
		}
		arr[l], arr[r] = arr[r], arr[l]
	}
	arr[left], arr[r] = arr[r], arr[left]
	return r
}
