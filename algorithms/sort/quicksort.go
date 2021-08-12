package sort

func QuickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	i, j, x := l-1, r+1, nums[(l+r)>>1]
	for i < j {
		for {
			i++
			if nums[i] >= x {
				break
			}
		}
		for {
			j--
			if nums[j] <= x {
				break
			}
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	QuickSort(nums, l, j)
	QuickSort(nums, j+1, r)
}
