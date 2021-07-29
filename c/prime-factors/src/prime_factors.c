#include "prime_factors.h"

size_t find_factors(uint64_t n, uint64_t factors[static MAXFACTORS]) {
    uint64_t factor = 2;
    size_t found_factors = 0;

    while (n > 1) {
        if (n % factor == 0) {
            factors[found_factors] = factor;
            n = n / factor;
            found_factors++;
        } else {
            factor++;
        }
    }

    return found_factors;
}