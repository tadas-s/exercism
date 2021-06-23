#include "nucleotide_count.h"
#include <string.h>
#include <stdlib.h>
#include <stdio.h>
#include <stdint.h>

char *count(const char *dna_strand) {
    uint64_t a = 0, c = 0, g = 0, t = 0;
    char *result;

    for(size_t i = 0; dna_strand[i] != '\0'; i++) {
        switch(dna_strand[i]) {
            case 'A': a++; break;
            case 'C': c++; break;
            case 'G': g++; break;
            case 'T': t++; break;
            default:
                result = malloc(1);
                result[0] = '\0';
                return result;
            break;
        }
    }

    // max unsinged 64 bit integer will take 20 characters.
    // So total buffer length with annotations, separators and zero terminator
    // is 4*20 + 4*2 + 3 + 1 = 92
    result = malloc(92);

    snprintf(result, 92, "A:%lu C:%lu G:%lu T:%lu", a, c, g, t);

    return result;
}