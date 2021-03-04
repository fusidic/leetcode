package sort

// InsertSort ...
func InsertSort(arr []int) []int {
	l := len(arr)
	for i := 1; i < l; i++ {
		for j := 0; j < i; j++ {
			if arr[i] < arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
