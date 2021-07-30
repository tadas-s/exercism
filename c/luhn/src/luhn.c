#include "luhn.h"
#include <stdlib.h>
#include <string.h>
#include <stdint.h>
#include <sys/types.h>

bool luhn(const char *num) {
    const size_t length = strlen(num);
    size_t digits = 0;
    uint64_t sum = 0;

    for (ssize_t i = length - 1; i >= 0; i--) {
        if (num[i] >= '0' && num[i] <= '9') {
            const uint8_t digit = num[i] - '0';

            if (digits % 2) {
                sum += digit * 2 > 9 ? digit * 2 - 9 : digit * 2;
            } else {
                sum += digit;
            }

            digits++;
        } else if (num[i] == ' ') {
            continue; // skip spaces
        } else {
            return false;
        }
    }

    return digits > 1 && sum % 10 == 0;
}