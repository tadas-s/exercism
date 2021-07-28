#include "binary.h"
#include <string.h>

int convert(const char *input) {
    const size_t length = strlen(input);
    int decimal = 0;

    for (size_t i = 0; i < length; i++) {
        switch(input[i]) {
            case '0': break;
            case '1':
                decimal |= 1 << (length - (i + 1));
            break;
            default: return INVALID; break;
        }
    }

    return decimal;
}