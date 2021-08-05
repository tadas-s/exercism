#include "diamond.h"
#include <stdlib.h>
#include <string.h>

char **make_diamond(const char letter) {
    const size_t middle = letter - 'A';
    const size_t size = middle * 2 + 1;

    char *diamond = calloc(size * (size + 1), sizeof(char));
    char **lines = calloc(size, sizeof(char *));

    for (size_t i = 0; i < size; i++) {
        lines[i] = &diamond[i * (size + 1)];
        memset(&diamond[i * (size + 1)], ' ', size);
    }

    for (size_t i = 0; i <= middle; i++) {
        lines[i][middle - i] = lines[i][middle + i] = 'A' + i;
        lines[size - i - 1][middle - i] = lines[size - i - 1][middle + i] = 'A' + i;
    }

    return lines;
}
