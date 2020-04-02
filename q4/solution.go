package q4

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	var left float64
	var right float64
	if m > n {
		m, n, nums1, nums2 = n, m, nums2, nums1
	}
	imin, imax, halflen := 0, m, (m+n+1)/2
	for imin <= imax {
		i := (imin + imax) / 2
		j := halflen - i
		if i > 0 && nums1[i-1] > nums2[j] {
			// check nums1[imin, i-1]
			imax = i - 1
		} else if i < m && nums2[j-1] > nums1[i] {
			imin = i + 1
		} else {

			if i == 0 {
				left = float64(nums2[j-1])
			} else if j == 0 {
				left = float64(nums1[i-1])
			} else {
				left = math.Max(float64(nums1[i-1]), float64(nums2[j-1]))
			}

			// If m+n is odd, exactly we get a num
			if (m+n)%2 == 1 {
				return left
			}

			// Or we have to find "right"
			if i == m {
				right = float64(nums2[j])
			} else if j == n {
				right = float64(nums1[i])
			} else {
				right = math.Min(float64(nums1[i]), float64(nums2[j]))
			}
			return (left + right) / 2
		}
	}
	return (left + right) / 2
}
