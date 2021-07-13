#ifndef PYTHAGOREAN_TRIPLET
#define PYTHAGOREAN_TRIPLET

#include <stdlib.h>

typedef struct {
    unsigned int a;
    unsigned int b;
    unsigned int c;
} triplet_t;

typedef struct {
    size_t count;
    triplet_t triplets[];
} triplets_t;

triplets_t *triplets_with_sum(unsigned int sum);
void free_triplets(triplets_t *triplets);

#endif
