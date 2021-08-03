#include "run_length_encoding.h"
#include <string.h>
#include <stdlib.h>
#include <stdio.h>
#include <ctype.h>

char *encode(const char *text) {
    // Resulting string is always less or equal length of original
    char *result = calloc(strlen(text), sizeof(char) + 1);
    char buff[30];
    size_t count = 1;

    for (size_t i = 0; text[i] != '\0'; i++) {
        if (text[i] == text[i + 1]) {
            count++;
        } else {
            if (count > 1) {
                snprintf(buff, 30, "%ld%c", count, text[i]);
            } else {
                snprintf(buff, 30, "%c", text[i]);
            }
            strcat(result, buff);
            count = 1;
        }
    }

    // In case compression worked great - trim the excess allocated memory.
    result = realloc(result, strlen(result) + 1);

    return result;
}

char *decode(const char *data) {
    char *result = calloc(1, sizeof(char));
    size_t result_length = 0, result_offset = 0;
    size_t offset = 0;
    size_t count;
    char c;

    while (data[offset] != '\0') {
        if (isdigit(data[offset])) {
            char *next;
            count = strtold(&data[offset], &next);
            c = *next;
            offset = next - data + 1;
        } else {
            count = 1;
            c = data[offset];
            offset++;
        }

        if (result_offset + count >= result_length) {
            // double the result buffer each time it runs out
            result_length = (result_offset + count) * 2;
            result = realloc(result, result_length);
        }

        for (size_t i = 0; i < count; i++) {
            result[result_offset++] = c;
        }

        result[result_offset] = '\0';
    }

    // Trim the excess allocated memory
    result = realloc(result, strlen(result) + 1);

    return result;
}
