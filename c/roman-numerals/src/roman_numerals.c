#include "roman_numerals.h"
#include <stdlib.h>
#include <string.h>
#include <math.h>
#include <sys/types.h>

const char *units[][9] = {
    { "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX" },
    { "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC" },
    { "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM" },
    { "M", "MM", "MMM" }
};

char *to_roman_numeral(unsigned int number) {
    if (number == 0 || number > 4000) {
        return NULL;
    }

    // Guesstimate the max resulting string size: up to 4 characters per digit
    // and 1 character for termination.
    const size_t size = (floor(log10(number)) + 1) * 4 + 1;
    char *buffer = calloc(size, sizeof(char));
    char *offset = buffer;

    for(ssize_t i = 3; i >= 0; i--) {
        const unsigned int digit = (number % (unsigned int)pow(10, i + 1)) / (unsigned int)pow(10, i);
        if (digit > 0) {
            strcpy(offset, units[i][digit - 1]);
            offset += strlen(units[i][digit - 1]);
        }
    }

    return buffer;
}