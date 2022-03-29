/**
 * The knows API is defined in the parent class Relation.
 * isBadVersion(version: number): boolean {
 *     ...
 * };
 */

var solution = function (isBadVersion: any) {
  return function (n: number): number {
    let r: number;
    let start = 1;
    let end = n;
    while (start <= end) {
      const mid = start + Math.floor((end - start) / 2);
      if (isBadVersion(mid)) {
        r = mid;
        end = mid - 1;
      } else {
        start = mid + 1;
      }
    }
    return r;
  };
};
