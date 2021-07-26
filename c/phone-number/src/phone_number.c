#include "phone_number.h"
#include <stdlib.h>
#include <string.h>
#include <ctype.h>
#include <stdbool.h>

#define NO_OF_DIGITS 10

void phone_number_reset(char *input);

char *phone_number_clean(const char *input) {
    size_t length = strlen(input);
    size_t digits = 0;

    char *result = malloc(NO_OF_DIGITS + 1);
    phone_number_reset(result);

    for (size_t i = 0; i < length; i++) {
        if (isdigit(input[i]) && !(digits == 0 && input[i] == '1')) {
            if (digits == NO_OF_DIGITS) {
                // empty error response - too many digits
                phone_number_reset(result);
                return result;
            }

            result[digits] = input[i];
            digits++;
        }
    }

    // No enough digits
    if (digits < NO_OF_DIGITS) { phone_number_reset(result); return result; }
    // Area code starts with 0
    if (result[0] == '0') { phone_number_reset(result); return result; }
    // Area code starts with 1
    if (result[0] == '1') { phone_number_reset(result); return result; }
    // Exchange code starts with 0
    if (result[3] == '0') { phone_number_reset(result); return result; }
    // Exchange code starts with 1
    if (result[3] == '1') { phone_number_reset(result); return result; }

    return result;
}

void phone_number_reset(char *input) {
    memset(input, '0', NO_OF_DIGITS);
    input[NO_OF_DIGITS] = '\0';
}