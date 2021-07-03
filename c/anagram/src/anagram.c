#include "anagram.h"
#include <stdbool.h>
#include <utf8/utf8.h>

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
    const void *w1_scap_p = w1;
    const size_t length = utf8size(w1);

    if (length != utf8size(w2)) { return false; }

    if (utf8casecmp(w1, w2) == 0) { return false; }

    while (1) {
        utf8_int32_t c;
        size_t count_w1 = 0, count_w2 = 0;
        const void *w1sp = w1;
        const void *w2sp = w2;

        w1_scap_p = utf8codepoint(w1_scap_p, &c);

        if (c == 0) { break; }

        c = utf8lwrcodepoint(c);

        while (1) {
            utf8_int32_t c1, c2;

            w1sp = utf8codepoint(w1sp, &c1);
            if (c1 == 0) { break; }

            w2sp = utf8codepoint(w2sp, &c2);
            if (c2 == 0) { break; }

            if (c == utf8lwrcodepoint(c1)) { count_w1++; }
            if (c == utf8lwrcodepoint(c2)) { count_w2++; }
        }

        if (count_w1 != count_w2) {
            return false;
        }
    }

    return true;
}

