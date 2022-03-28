function search(nums: number[], target: number): number {
  let start = 0;
  let end = nums.length - 1;

  while (start <= end) {
    const mid = start + ~~((end - start) / 2);
    const cur = nums[mid];
    if (cur === target) {
      return mid;
    }
    if (cur < target) {
      start = mid + 1;
    } else {
      end = mid - 1;
    }
  }

  return -1;
}
