function search(nums: number[], target: number): number {
  let start = 0;
  let end = nums.length - 1;

  while (start <= end) {
    const mid = start + ~~((end - start) / 2); // get mid pointer
    const cur = nums[mid];
    if (cur === target) { // if mid value equal to target return the mid
      return mid;
    }
    if (cur < target) { // if mid value is less than target should move the left pointer to mid+1
      start = mid + 1;
    } else { // else if mid value is greater than target should move the right pointer to mid-1
      end = mid - 1;
    }
  }

  return -1; // if can't find target return -1
}
