#include "anagram.h"
#include <stdbool.h>
#include <string.h>
#include <strings.h>
#include <ctype.h>

bool is_anagram(const char *w1, const char *w2);

void find_anagrams(const char *subject, struct candidates *candidates) {
    for (size_t i = 0; i < candidates->count; i++) {
        if (is_anagram(subject, candidates->candidate[i].word)) {
            candidates->candidate[i].is_anagram = IS_ANAGRAM;
        } else {
            candidates->candidate[i].is_anagram = NOT_ANAGRAM;
        }
    }
}

bool is_anagram(const char *w1, const char *w2) {
    const size_t length = strlen(w1);

    if (length != strlen(w2)) { return false; }

    if (strcasecmp(w1, w2) == 0) { return false; }

    for (size_t i = 0; i < length; i++) {
        char c = tolower(w1[i]);
        size_t count_w1 = 0, count_w2 = 0;

        for (size_t j = 0; j < length; j++) {
            if (c == tolower(w1[j])) { count_w1++; }
            if (c == tolower(w2[j])) { count_w2++; }
        }

        if (count_w1 != count_w2) {
            return false;
        }
    }

    return true;
}

