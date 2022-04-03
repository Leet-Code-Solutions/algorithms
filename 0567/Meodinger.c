#include <stdio.h>

#define true  1
#define false 0
typedef unsigned char bool;

bool checkInclusion(char *s1, char *s2) {
    int score[26] = {0}, i, j, len;

    for (i = 0; s1[i]; i++) {
        score[s1[i] - 'a'] += 1;
        if (!s2[i]) return false;
        score[s2[i] - 'a'] -= 1;
    }
    len = i;

    for (i = 0; s2[i + len - 1]; i++) {
        for (j = 0; j < 26; j++) {
            if (score[j]) break;
            if (j == 25) return true;
        }
        if (!s2[i + len]) return false;
        score[s2[i]       - 'a'] += 1;
        score[s2[i + len] - 'a'] -= 1;
    }

    return false;
}
