package _035

func searchInsert(nums []int, target int) int {
	min := 0
	max := len(nums) - 1
	for {
		if min > max {
			break
		}
		mid := min + (max - min) / 2
		if nums[min] == target {
			return mid
		}
		
		if nums[mid] > target {
			max = mid - 1
		} else {
			min = mid + 1
		}
	}

	return min
}