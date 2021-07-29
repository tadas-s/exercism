#include "etl.h"
#include <stdlib.h>
#include <string.h>
#include <ctype.h>

int compare_new_map_item(const void *a, const void *b);

int convert(const legacy_map *input, const size_t input_len, new_map **output) {
    size_t entries = 0;

    for (size_t i = 0; i < input_len; i++) {
        entries += strlen(input[i].keys);
    }

    *output = malloc(sizeof(new_map) * entries);
    size_t output_offset = 0;

    for (size_t i = 0; i < input_len; i++) {
        for (size_t j = 0; input[i].keys[j] != '\0'; j++) {
            (*output)[output_offset] = (new_map){
                tolower(input[i].keys[j]), input[i].value
            };
            output_offset++;
        }
    }

    qsort(*output, entries, sizeof(new_map), compare_new_map_item);

    return entries;
}

int compare_new_map_item(const void *a, const void *b) {
    return ((new_map *)a)->key - ((new_map *)b)->key;
}