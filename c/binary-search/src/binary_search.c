#include "binary_search.h"

const int *binary_search(int value, const int *arr, size_t length) {
    int lower_bound = 0;
    int upper_bound = length - 1;

    while (lower_bound <= upper_bound) {
        int i = (lower_bound + upper_bound) / 2;

        if (arr[i] < value) {
            lower_bound = i + 1;
        } else if (arr[i] > value) {
            upper_bound = i - 1;
        } else {
            return &arr[i];
        }
    }

    return NULL;
}