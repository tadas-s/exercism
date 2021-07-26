#include "nth_prime.h"
#include <stdbool.h>

bool is_prime(uint32_t n);

uint32_t nth(uint32_t n) {
    uint32_t number = 2, sequence = 1;

    if (n < 1) { return 0; }

    while (sequence < n) {
        do {
            number++;
        } while (!is_prime(number));

        sequence++;
    }

    return number;
}

bool is_prime(uint32_t n) {
    if (n <= 1) {
        return false;
    } else if (n <= 3) {
        return true;
    } else if (n % 2 == 0 || n % 3 == 0) {
        return false;
    } else {
        uint32_t divisor = 5;

        while (divisor * divisor <= n) {
            if (n % divisor == 0 || n % (divisor + 2) == 0) {
                return false;
            }
            divisor += 6;
        }

        return true;
    }
}