// Package sort implements some classic sort algorithms.
package sort

// MergeSort implements a merge sort algorithm.
func MergeSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	mid := (l + r) >> 1
	MergeSort(nums, l, mid)
	MergeSort(nums, mid+1, r)
	i, j := l, mid+1
	var tmp []int
	for i <= mid && j <= r {
		if nums[i] <= nums[j] {
			tmp = append(tmp, nums[i])
			i++
		} else {
			tmp = append(tmp, nums[j])
			j++
		}
	}
	for i <= mid {
		tmp = append(tmp, nums[i])
		i++
	}
	for j <= r {
		tmp = append(tmp, nums[j])
		j++
	}
	j = 0
	for i := l; i <= r; i++ {
		nums[i] = tmp[j]
		j++
	}
}
