#include "hamming.h"
#include <stddef.h>

int compute(const char *lhs, const char *rhs) {
    if (lhs == NULL || rhs == NULL) {
        return -1;
    }

    size_t offset = 0;
    int distance = 0;

    while (1) {
        if (lhs[offset] == '\0' && rhs[offset] == '\0') {
            return distance;
        } else if (lhs[offset] == '\0' || rhs[offset] == '\0') {
            return -1;
        } else {
            if (lhs[offset] != rhs[offset]) {
                distance++;
            }
        }

        offset++;
    }

    return distance;
}