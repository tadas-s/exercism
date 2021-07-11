#include "roman_numerals.h"
#include <stdlib.h>
#include <string.h>
#include <math.h>
#include <sys/types.h>

char *to_roman_numeral(unsigned int number) {
    const char numerals[] = "IVXLCDM  ";

    if (number > 4000) {
        return NULL;
    }

    // Guesstimate the max resulting string size: up to 4 characters per digit
    // and 1 character for termination.
    const size_t size = (floor(log10(number)) + 1) * 4 + 1;
    char *buffer = malloc(size);
    memset(buffer, '\0', size);

    size_t offset = 0;

    for(ssize_t i = 3; i >= 0; i--) {
        char one = numerals[i * 2];
        char five = numerals[i * 2 + 1];
        char ten = numerals[i * 2 + 2];

        switch((number % (unsigned int)pow(10, i + 1)) / (unsigned int)pow(10, i)) {
            case 1:
                buffer[offset] = one;
                offset += 1;
            break;
            case 2:
                buffer[offset++] = one;
                buffer[offset++] = one;
            break;
            case 3:
                buffer[offset++] = one;
                buffer[offset++] = one;
                buffer[offset++] = one;
            break;
            case 4:
                buffer[offset++] = one;
                buffer[offset++] = five;
            break;
            case 5:
                buffer[offset++] = five;
            break;
            case 6:
                buffer[offset++] = five;
                buffer[offset++] = one;
            break;
            case 7:
                buffer[offset++] = five;
                buffer[offset++] = one;
                buffer[offset++] = one;
            break;
            case 8:
                buffer[offset++] = five;
                buffer[offset++] = one;
                buffer[offset++] = one;
                buffer[offset++] = one;
            break;
            case 9:
                buffer[offset++] = one;
                buffer[offset++] = ten;
            break;
        }
    }

    return buffer;
}