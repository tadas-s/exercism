#include "pythagorean_triplet.h"
#include <stdlib.h>

triplets_t *triplets_with_sum(unsigned int sum) {
    triplets_t *result = malloc(sizeof(triplets_t));
    result->count = 0;

    for (unsigned int a = 1; a <= sum / 3; a++) {
        for(unsigned int b = a + 1; b <= sum / 2; b++) {
            unsigned int c = sum - a - b;

            if (a * a + b * b == c * c) {
                result->count++;
                result = realloc(result, sizeof(triplets_t) + result->count * sizeof(triplet_t));
                result->triplets[result->count - 1] = (triplet_t){ a, b, c};
            }
        }
    }

    return result;
}

void free_triplets(triplets_t *triplets) {
    free(triplets);
}
