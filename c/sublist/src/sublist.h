#ifndef SUBLIST_H
#define SUBLIST_H

#include <stddef.h>
#include <stdbool.h>

typedef enum {
   EQUAL,
   UNEQUAL,
   SUBLIST,
   SUPERLIST
} comparison_result_t;

comparison_result_t check_lists(int *a, int *b, size_t a_count, size_t b_count);
bool check_list_is_sublist(int *a, int *b, size_t a_count, size_t b_count);

#endif
