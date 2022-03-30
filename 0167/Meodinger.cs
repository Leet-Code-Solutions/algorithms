public class Solution {
    public int[] TwoSum(int[] numbers, int target) {        
        int a = 0, b = numbers.Length - 1;
        while (a < b) {
            if (numbers[a] + numbers[b] == target) return new int[] { a+1, b+1 };
            if (numbers[a] + numbers[b] > target) {
                b--;
            } else {
                a++;
            }
        }
        return new int[0] {}; // for pass compile
    }
}