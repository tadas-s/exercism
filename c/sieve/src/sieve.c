#include "sieve.h"
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>
#include <math.h>

uint32_t sieve(uint32_t limit, uint32_t * primes, size_t max_primes) {
    bool *sieve = malloc(limit + 1);
    memset(sieve, 1, limit + 1);

    uint32_t number = 2;

    while (number * number <= limit) {
        uint32_t multiple = number * 2;

        while (multiple <= limit) {
            sieve[multiple] = false;
            multiple += number;
        }

        do {
            number++;
        } while (number <= limit && sieve[number] == false);
    }

    size_t primes_found = 0;

    for (size_t i = 2; (i <= limit) && (primes_found < max_primes); i++) {
        if (sieve[i]) {
            primes[primes_found] = i;
            primes_found++;
        }
    }

    free(sieve);

    return primes_found;
}