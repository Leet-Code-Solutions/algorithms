/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
     min := 1
     max := n
    
     for {
         if max <= min {
             break
         }
          mid := min + (max-min)/2
         if isBadVersion(mid) {
             max = mid
             continue
         }

         min = mid + 1
     }

     return min
}
