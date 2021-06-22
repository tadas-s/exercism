#include "pangram.h"
#include <stddef.h>
#include <stdint.h>
#include <string.h>
#include <ctype.h>

bool is_pangram(const char *sentence) {
    if (sentence == NULL) { return false; }

    const size_t length = strlen(sentence);

    // English alphabet has 26 letters. Use 32 bit word
    // to store bit flags for each letter seen.
    uint32_t chars_seen = 0;

    for(size_t i = 0; i < length; i++) {
        char c = tolower(sentence[i]);

        if(isalpha(c)) {
            chars_seen |= 1 << (c - 'a');

            if (chars_seen == 0x03ffffff) {
                return true;
            }
        }
    }

    return false;
}