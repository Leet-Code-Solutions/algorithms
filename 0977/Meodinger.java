public class Solution {
    public int[] sortedSquares(int[] nums) {
        int v1, v2;
        int l = 0, r = nums.length - 1;
        int[] array = new int[nums.length];

        while (l <= r) {
            v1 = Math.abs(nums[l]);
            v2 = Math.abs(nums[r]);
            if (v1 < v2) {
                r--;
                arr[i--] = v2 * v2;
            } else {
                l++;
                arr[i--] = v1 * v1;
            }
        }

        return array;
    }
}
