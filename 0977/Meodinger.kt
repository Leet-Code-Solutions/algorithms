import kotlin.math.abs

class Solution {
    fun sortedSquares(nums: IntArray): IntArray {
        var i = nums.size - 1
        var v1: Int
        var v2: Int

        var l = 0
        var r = i

        val arr = IntArray(i + 1)

        while (l <= r) {
            v1 = abs(nums[l])
            v2 = abs(nums[r])
            if (v1 < v2) {
                r--
                arr[i--] = v2 * v2
            } else {
                l++
                arr[i--] = v1 * v1
            }
        }

        return arr
    }
}