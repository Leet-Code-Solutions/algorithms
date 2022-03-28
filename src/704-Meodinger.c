int biSearch(int target, int* nums, int l, int r) {
    int m, v;
    while (l <= r) {
        m = l + (r - l) / 2;
        v = nums[m];
 
        if (v == target) return m;
 
        if (v < target)
            l = m + 1;
         else
            r = m - 1;
    }
    return -1;
} 

int search(int* nums, int numsSize, int target) {
    return biSearch(target, nums, 0, numsSize - 1);
}