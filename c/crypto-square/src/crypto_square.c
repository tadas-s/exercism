#include "crypto_square.h"
#include <stdlib.h>
#include <string.h>
#include <ctype.h>
#include <math.h>

char *ciphertext(const char *input) {
    const size_t length = strlen(input);
    char *buff = calloc(length + 1, sizeof(char));

    size_t normalized_length = 0;

    for (size_t i = 0; i < length; i++) {
        if (isalnum(input[i])) {
            buff[normalized_length++] = tolower(input[i]);
        }
    }

    size_t columns = floor(sqrt(normalized_length));
    size_t rows = floor(sqrt(normalized_length));
    if (columns * rows < normalized_length) { rows++; }
    if (columns * rows < normalized_length) { columns++; }

    size_t final_length = columns > 0 ? columns * rows + rows - 1 : 0;

    char *crypto_square = malloc(final_length + 1);
    memset(crypto_square, ' ', final_length);
    crypto_square[final_length] = '\0';

    for (size_t c = 0; c < columns; c++) {
        for (size_t r = 0; r < rows; r++) {
            if (c * rows + r < normalized_length) {
                crypto_square[r * (columns + 1) + c] = buff[c * rows + r];
            }
        }
    }

    free(buff);

    return crypto_square;
}