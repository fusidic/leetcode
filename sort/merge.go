package sort

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	res := make([]int, 0)
	l1, l2 := 0, len(left)
	r1, r2 := 0, len(right)
	for l1 < l2 && r1 < r2 {
		if left[l1] > right[r1] {
			res = append(res, right[r1])
			r1++
			continue
		}
		res = append(res, left[l1])
		l1++
	}
	res = append(res, left[l1:]...)
	res = append(res, right[r1:]...)
	return res
}
