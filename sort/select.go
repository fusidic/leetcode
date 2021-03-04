package sort

// SelectSort ...
func SelectSort(arr []int) []int {
	min := 0

	for i := 0; i < len(arr); i++ {
		min = i
		for j := i; j < len(arr); j++ {
			if arr[j] < min {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}
