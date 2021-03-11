package sort

func Merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(result, right...)
		}
		if len(right) == 0 {
			return append(result, left...)
		}
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	return result
}

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	middle := len(arr) / 2

	left := MergeSort(arr[:middle])
	right := MergeSort(arr[middle:])
	res := Merge(left, right)
	return res
}

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
	res := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(res, right...)
		}
		if len(right) == 0 {
			return append(res, left...)
		}
		if left[0] <= right[0] {
			res = append(res, left[0])
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}
	return res
}

func mergeSort0311(nums []int) []int {
	l := len(nums)
	if l <= 1 {
		return nums
	}
	mid := l / 2
	left := mergeSort0311(nums[:mid])
	right := mergeSort0311(nums[mid:])
	return merge0311(left, right)
}

func merge0311(left, right []int) []int {
	res := []int{}
	l1, l2 := 0, len(left)
	r1, r2 := 0, len(right)
	for l1 < l2 && r1 < r2 {
		if left[l1] <= right[r1] {
			res = append(res, left[l1])
			l1++
		} else {
			res = append(res, right[r2])
			r1++
		}
	}
	res = append(res, left[l1:]...)
	res = append(res, right[r1:]...)
	return res
}
