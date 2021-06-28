#include "sum_of_multiples.h"

unsigned int sum(const unsigned int *factors, const size_t number_of_factors, const unsigned int limit) {
    unsigned int total = 0;

    for (unsigned int n = 1; n < limit; n++) {
        for (size_t factor = 0; factor < number_of_factors; factor++) {
            if (factors[factor] > 0 && n % factors[factor] == 0) {
                total += n;
                break;
            }
        }
    }

    return total;
}