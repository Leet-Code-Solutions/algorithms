int lengthOfLongestSubstring(char * s) {
    int hash[128] = { 0 };
    int ptr = -1, ans = 0;

    for (int i = 0; s[i]; i++) {
        if (i) hash[s[i - 1]] = 0;
        while (s[ptr + 1] && !hash[s[ptr + 1]]) hash[s[ptr++ + 1]] = 1;
        if (ans < ptr - i + 1) ans = ptr - i + 1; 
    }
    return ans;
}
