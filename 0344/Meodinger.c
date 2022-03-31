void reverseString(char* s, int size) {
    int l = 0, r = size - 1;
    while (l < r) {
        s[r] = s[l] ^ s[r];
        s[l] = s[l] ^ s[r];
        s[r] = s[l] ^ s[r];

        l++;
        r--;
    }
}