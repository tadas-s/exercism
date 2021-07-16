#include "atbash_cipher.h"
#include <stdlib.h>
#include <string.h>
#include <ctype.h>

void add_space(char *input, size_t *result_offset);

char *atbash_encode(const char *input) {
    const size_t length = strlen(input);
    // Max result length: up to original + spaces every 5 chars + zero termination
    const size_t max_result_length = length + length / 5 + 1;

    char *result = malloc(max_result_length);
    memset(result, '\0', max_result_length);

    size_t result_offset = 0;

    for (size_t i = 0; i < length; i++) {
        char c = tolower(input[i]);

        if (c >= 'a' && c <= 'z') {
            add_space(result, &result_offset);
            result[result_offset] = 'z' - (c - 'a');
            result_offset++;
        } else if (isdigit(c)) {
            add_space(result, &result_offset);
            result[result_offset] = c;
            result_offset++;
        }
    }

    return result;
}

void add_space(char *input, size_t *result_offset) {
    if (*result_offset && *result_offset % 6 == 5) {
        input[*result_offset] = ' ';
        (*result_offset)++;
    }
}

char *atbash_decode(const char *input) {
    const size_t length = strlen(input) + 1;

    char *result = malloc(length);
    memset(result, '\0', length);

    size_t result_offset = 0;

    for (size_t i = 0; i < length; i++) {
        char c = tolower(input[i]);

        if (c >= 'a' && c <= 'z') {
            result[result_offset] = 'z' - (c - 'a');
            result_offset++;
        } else if (!isspace(c)) {
            result[result_offset] = c;
            result_offset++;
        }
    }

    return result;
}
