package sort

// QuickSort ...
func QuickSort(arr []int) []int {
	partition(arr, 0, len(arr)-1)
	return arr
}

func partition(arr []int, left, right int) {
	if left >= right {
		return
	}

	base := arr[left]
	l, r := left, right

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
	partition(arr, left, r-1)
	partition(arr, r+1, right)
}

func quickSort0310(nums []int) []int {
	partition0310(nums, 0, len(nums)-1)
	return nums
}

func partition0310(nums []int, left, right int) {
	if left >= right {
		return
	}

	base := nums[left]
	l, r := left, right
	for true {
		for l < right && nums[l] <= base {
			l++
		}
		for r > left && nums[r] >= base {
			r--
		}
		if l >= r {
			break
		}
		nums[l], nums[r] = nums[r], nums[l]
	}
	nums[left], nums[r] = nums[r], nums[left]
	partition0310(nums, left, r-1)
	partition0310(nums, r+1, right)
}
