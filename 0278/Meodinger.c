
int biSearch(int l, int r) {
    int m;
    while (l <= r) {
        m = l + (r - l) / 2; 
        
        if (isBadVersion(m))
            r = m - 1;
         else
            l = m + 1;
    }
    return l;
} 


int firstBadVersion(int n) {
    return biSearch(0, n);    
}