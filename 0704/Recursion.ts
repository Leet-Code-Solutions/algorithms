function search(nums: number[], target: number): number {
  const compareFn = (start: number, end: number) => {
      if (start > end) return -1 // if start is greater than end it means that nums doesn't include target, we will return -1

      const mid = start === end ? start : start + Math.floor((end - start) / 2) // get mid pointer
      const cur = nums[mid]
      if (cur === target) return mid // if mid value is equal to target we will return mid
      // if mid value is greater than target we will move end pointer to end-1 or move start pointer to start+1 
      return cur > target ? compareFn(start, end -1) : compareFn(start + 1, end)
  }
  return compareFn(0, nums.length - 1)
};
