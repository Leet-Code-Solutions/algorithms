package poplar_hsu

func rotate(nums []int, k int) {
	if k == 0 {
		return
	}
	k %= len(nums)                // 找到翻转后的中间位置
	reverse(nums, 0, len(nums)-1) // 第一次将整个数组翻转过来
	reverse(nums, 0, k-1)         // 翻转左边
	reverse(nums, k, len(nums)-1) //翻转右边
}

func reverse(nums []int, l, r int) {
	for {
		if l > r {
			break
		}
		tmp := nums[l]
		nums[l] = nums[r]
		nums[r] = tmp
		l++
		r--
	}
}
