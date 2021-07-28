#include "sublist.h"
#include <stdbool.h>
#include <string.h>

comparison_result_t check_lists(int *a, int *b, size_t a_count, size_t b_count) {
    if (a_count < b_count && check_list_is_sublist(a, b, a_count, b_count)) {
        return SUBLIST;
    } else if (a_count == b_count && check_list_is_sublist(a, b, a_count, b_count)) {
        return EQUAL;
    } else if (a_count > b_count && check_list_is_sublist(b, a, b_count, a_count)) {
        return SUPERLIST;
    } else {
        return UNEQUAL;
    }
}

// Function assumes a size is smaller or equal to b
bool check_list_is_sublist(int *a, int *b, size_t a_count, size_t b_count) {
    for (size_t i = 0; i < (b_count - a_count + 1); i++) {
        if (memcmp(a, &b[i], a_count * sizeof(int)) == 0) {
            return true;
        }
    }
    return false;
}