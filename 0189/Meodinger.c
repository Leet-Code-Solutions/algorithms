void reverse(int* nums, int l, int r) {
    int tmp;
    while (l < r) {
        tmp = nums[l];
        nums[l] = nums[r];
        nums[r] = tmp;
        l++;
        r--;
    }
}

void rotate(int* nums, int numsSize, int k) {
    k %= numsSize;

    if (k == 0) return;

    reverse(nums, 0, numsSize - 1);
    reverse(nums, 0, k - 1);
    reverse(nums, k, numsSize - 1);
}