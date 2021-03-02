package sort

// QuickSort ...
func QuickSort(arr []int, left, right int) {
	base := arr[left]
	l, r := left, right
	for true {
		for l < right && arr[l] >= base {
			l++
		}
		for r > left && arr[r] <= base {
			r--
		}
		if l >= r {
			break
		}
		arr[l], arr[r] = arr[r], arr[l]
	}
	arr[left], arr[r] = arr[r], arr[left]
	QuickSort(arr, left, r-1)
	QuickSort(arr, r+1, right)
}
