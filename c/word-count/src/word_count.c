#include "word_count.h"
#include <stdlib.h>
#include <string.h>
#include <ctype.h>

int count_word(const char *word_start, const char *word_end, word_count_word_t *words);

int count_words(const char *sentence, word_count_word_t * words) {
    int total = 0;
    int err;
    word_count_word_t result[MAX_WORDS];
    memset(result, 0, sizeof(word_count_word_t) * MAX_WORDS);

    if (sentence == NULL) {
        memcpy(words, result, sizeof(word_count_word_t) * MAX_WORDS);
        return 0;
    }

    const char *end_of_sentence = sentence + strlen(sentence);
    const char *current_word = NULL;

    for (const char *current_char = sentence; current_char < end_of_sentence; current_char++) {
        char previous_char = current_char > sentence ? *(current_char - 1) : '\0';
        char next_char = current_char < end_of_sentence ? *(current_char + 1) : '\0';

        if (!current_word && isalnum(*current_char)) {
            current_word = current_char;
        } else if (*current_char == '\'' && isalnum(previous_char) && isalnum(next_char)) {
            // apostrophe - continue "scanning" the word
        } else if (current_word && !isalnum(*current_char)) {
            err = count_word(current_word, current_char - 1, result);
            if (err < 0) { return err; }
            current_word = NULL;
        }
    }

    if (current_word) {
        err = count_word(current_word, end_of_sentence - 1, result);
        if (err < 0) { return err; }
    }

    memcpy(words, result, sizeof(word_count_word_t) * MAX_WORDS);

    for (size_t i = 0; i < MAX_WORDS; i++) {
        if (result[i].count > 0) {
            total++;
        };
    }

    return total;
}

int count_word(const char *word_start, const char *word_end, word_count_word_t *words) {
    word_count_word_t *slot = NULL;

    if (word_end - word_start + 1 > MAX_WORD_LENGTH) {
        return EXCESSIVE_LENGTH_WORD;
    }

    char *normalized = malloc(word_end - word_start + 2);
    strncpy(normalized, word_start, word_end - word_start + 1);
    normalized[word_end - word_start + 1] = '\0';

    for (size_t i = 0; i < strlen(normalized); i++) {
        normalized[i] = tolower(normalized[i]);
    }

    size_t i = 0;

    do {
        if (strlen(words[i].text) == 0) {
            slot = &words[i];
            strcpy(words[i].text, normalized);
            words[i].count = 1;
        } else if (strcmp(words[i].text, normalized) == 0) {
            slot = &words[i];
            words[i].count++;
        } else {
            i++;
        }
    } while (slot == NULL && i < MAX_WORDS);

    free(normalized);

    if (slot == NULL) {
        return EXCESSIVE_NUMBER_OF_WORDS;
    }

    return slot->count;
}