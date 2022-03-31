void swap(char *s, int l, int r) {
    while (l < r) {
        s[r] = s[l] ^ s[r];
        s[l] = s[l] ^ s[r];
        s[r] = s[l] ^ s[r];

        l++;
        r--;
    }
}

char *reverseWords(char *s) {
    int r = 0, l = 0;
    while (s[r]) {
        if (s[r] == ' ') {
            swap(s, l, r - 1);
            l = r + 1;
        }
        r++;
    }
    swap(s, l, r - 1);

    return s;
}