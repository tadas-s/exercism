#include "rna_transcription.h"
#include <stdlib.h>
#include <string.h>

char *to_rna(const char *dna) {
    if (dna == NULL) {
        return NULL;
    }

    char *rna = malloc(strlen(dna) + 1);
    size_t offset = 0;

    while (1) {
        switch (dna[offset]) {
            case '\0':
                rna[offset] = '\0';
                return rna;
                break;
            case 'G':
                rna[offset] = 'C';
                break;
            case 'C':
                rna[offset] = 'G';
                break;
            case 'T':
                rna[offset] = 'A';
                break;
            case 'A':
                rna[offset] = 'U';
                break;
            default:
                free(rna);
                return NULL;
                break;
        }

        offset++;
    }

    return rna;
}