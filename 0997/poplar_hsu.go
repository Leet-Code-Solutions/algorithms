package poplar_hsu

func sortedSquares(nums []int) []int {
	l := 0
	r := len(nums) - 1
	m := 0

	// 先找到平方后最小值的下标
	for i, v := range nums {
		if v > 0 {
			m = i
			break
		}
	}

	arr := make([]int, 0)
	for  {
		if(l >= m && r < m) || l > r {
			break
		}
		l2 := nums[l] * nums[l]
		r2 := nums[r] * nums[r]
		if(l2 < r2) {
			r --
			arr = append([]int{r2}, arr...)
		} else {
			l ++
			arr = append([]int{l2}, arr...)
		}
	}
	return arr
}