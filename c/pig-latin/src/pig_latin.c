#include "pig_latin.h"
#include <string.h>
#include <ctype.h>
#include <stdlib.h>
#include <stdbool.h>

typedef enum {
    state_start,
    state_word,
    state_delimiter,
    state_end
} state_t;

size_t translate_word(char *dest, const char *word_start, const char *word_end);
bool is_vowel(const char c);
bool is_consonant(const char c);

char *translate(const char *phrase) {
    const size_t length = strlen(phrase);
    // tripple length string for approximate worst case result
    char *translated = calloc(length * 3, sizeof(char));
    const char *word_start = NULL;
    state_t state = state_start;
    size_t i = 0, j = 0;

    while (state != state_end) {
        switch(state) {
            case state_start:
                if (isalpha(phrase[i])) {
                    state = state_word;
                    word_start = &phrase[i];
                } else {
                    state = state_delimiter;
                }
            break;

            case state_word:
                if (!isalpha(phrase[i])) {
                    state = state_delimiter;
                    j += translate_word(&translated[j], word_start,
                                        &phrase[i - 1]);
                } else {
                    i++;
                }
            break;

            case state_delimiter:
                if (isalpha(phrase[i])) {
                    state = state_word;
                    word_start = &phrase[i];
                } else if (phrase[i] == '\0') {
                    state = state_end;
                } else {
                    translated[j++] = phrase[i++];
                }
            break;

            default: break;
        }
    }

    translated = realloc(translated, strlen(translated) + 1);

    return translated;
}

size_t translate_word(char *dest, const char *word_start, const char *word_end) {
    if (is_vowel(word_start[0]) || strncmp(word_start, "yt", 2) == 0 ||
          strncmp(word_start, "xr", 2) == 0) {
        memcpy(dest, word_start, word_end - word_start + 1);
    } else {
        const char *first_vowel = word_start;

        do {
            first_vowel++;
        } while(is_consonant(*first_vowel) && !(*first_vowel == 'y'));

        if (strncmp(first_vowel - 1, "qu", 2) == 0) { first_vowel++; }

        memcpy(dest, first_vowel, word_end - first_vowel + 1);
        memcpy(dest + (word_end - first_vowel + 1), word_start,
               first_vowel - word_start);
    }

    dest[word_end - word_start + 1] = 'a';
    dest[word_end - word_start + 2] = 'y';

    return (word_end - word_start + 3); // +2?
}

bool is_vowel(const char c) {
    return c == 'a' || c == 'e' || c == 'i' ||
           c == 'o' || c == 'u';
}

bool is_consonant(const char c) {
    return isalpha(c) && !is_vowel(c);
}