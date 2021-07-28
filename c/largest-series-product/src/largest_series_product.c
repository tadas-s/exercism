#include "largest_series_product.h"
#include <string.h>

int64_t largest_series_product(char *digits, size_t span) {
    const size_t length = strlen(digits);
    int64_t max = -1;

    for (size_t i = 0; i < (length - span + 1); i++) {
        int64_t product = 1;

        for (size_t j = 0; j < span; j++) {
            int8_t digit = digits[i+j] - '0';

            if (digit < 0 || digit > 9) { return -1; }

            product = product * digit;
        }

        if (product > max) {
            max = product;
        }
    }

    return max;
}