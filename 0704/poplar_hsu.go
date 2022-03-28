func search(nums []int, target int) int {
    max := len(nums) - 1
    min := 0

    for {
        if max < min {
            return -1
        }

        mid := (max+min)/2
        if nums[mid] == target {
            return mid
        }

        // 为了避免死循环max重新取值需要-1
         if nums[mid] > target {
            max = mid-1
        }

        // 同理
         if nums[mid] < target {
            min = mid+1
        }
    }

    return -1
}
