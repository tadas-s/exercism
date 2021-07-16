#include "series.h"
#include <stdlib.h>
#include <string.h>

slices_t slices(char *input_text, unsigned int substring_length) {
    const size_t length = strlen(input_text);
    slices_t result;

    if (substring_length > 0 && length >= substring_length) {
        result.substring_count = 1 + length - substring_length;
    } else {
        result.substring_count = 0;
        return result;
    }

    result.substring = malloc(result.substring_count * sizeof(char *));

    for (size_t i = 0; i < result.substring_count; i++) {
        result.substring[i] = malloc(substring_length + 1);
        strncpy(result.substring[i], input_text + i, substring_length);
        result.substring[i][substring_length] = '\0';
    }

    return result;
}