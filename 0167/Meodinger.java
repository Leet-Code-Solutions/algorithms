public class Solution {
    public int[] twoSum(int[] numbers, int target) {
        int a = 0, b = numbers.length - 1;
        while (a < b) {
            if (numbers[a] + numbers[b] == target)
                return new int[] { a + 1, b + 1 };
            if (numbers[a] + numbers[b] > target) {
                b--;
            } else {
                a++;
            }
        }
        return new int[0]; // for pass the compilation
    }
}
