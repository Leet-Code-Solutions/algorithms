// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。

// 示例 1:
// 输入: nums = [0,1,0,3,12]
// 输出: [1,3,12,0,0]

func moveZeroes(nums []int)  {
    j := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
          tmp := nums[i]
          nums[i] = nums[j]
          nums[j] = tmp 
          j++
        }
    }
}

// 左指针j，右指针i
// i: 遍历数组， i-j中间都为0
// j: 左边都为非0数


func moveZeroes(nums []int)  {
    j := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
          nums[j] = nums[i]
          j++
        }
    }

    for i := j; i< len(nums); i ++ {
        nums[i] = 0
    }
}