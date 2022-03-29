func twoSum(nums []int, target int) []int {
    maps := make(map[int]int)
    for i:=0; i<len(nums); i++ {
        t := target - nums[i]
        maps[t] = i
    }

    for i:=0;i<len(nums); i++ {
       index, ok := maps[nums[i]]
       if ok && index != i {
           return []int{index, i}
       }
    }

    return nil
} // 还可以简化一下成一个循环