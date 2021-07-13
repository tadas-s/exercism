#include "list_ops.h"
#include <string.h>

list_t *new_list(size_t length, list_element_t elements[]) {
    list_t *list = malloc(sizeof(list_t) + length * sizeof(list_element_t));
    list->length = length;
    memcpy(&list->elements, elements, length * sizeof(list_element_t));
    return list;
}

list_t *append_list(list_t * list1, list_t * list2) {
    const size_t length = list1->length + list2->length;
    list_t *list = malloc(sizeof(list_t) + length * sizeof(list_element_t));
    list->length = length;

    memcpy(&list->elements[0], list1->elements, list1->length * sizeof(list_element_t));
    memcpy(&list->elements[list1->length],
        list2->elements, list2->length * sizeof(list_element_t));

    return list;
}

list_t *filter_list(list_t * list, bool(*filter) (list_element_t)) {
    // Use temporary list for results
    list_t *tmp = malloc(sizeof(list_t) + list->length * sizeof(list_element_t));
    tmp->length = 0;

    // Final list will be allocated when we know exact filtered members
    list_t *final_list;

    for (size_t i = 0; i < list->length; i++) {
        if (filter(list->elements[i])) {
            tmp->elements[tmp->length++] = list->elements[i];
        }
    }

    final_list = new_list(tmp->length, tmp->elements);

    delete_list(tmp);

    return final_list;
}

list_t *map_list(list_t * list, list_element_t(*map) (list_element_t)) {
    list_t *mapped = malloc(sizeof(list_t) + list->length * sizeof(list_element_t));
    mapped->length = list->length;

    for (size_t i = 0; i < list->length; i++) {
        mapped->elements[i] = map(list->elements[i]);
    }

    return mapped;
}

list_element_t foldl_list(list_t * list, list_element_t initial,
                          list_element_t(*foldl) (list_element_t,
                                                  list_element_t)) {
    list_element_t result = initial;

    for (size_t i = 0; i < list->length; i++) {
        result = foldl(result, list->elements[i]);
    }

    return result;
}

list_element_t foldr_list(list_t * list, list_element_t initial,
                          list_element_t(*foldr) (list_element_t,
                                                  list_element_t)) {
    list_element_t result = initial;

    for (size_t i = list->length; i-- > 0;) {
        result = foldr(list->elements[i], result);
    }

    return result;
}

list_t *reverse_list(list_t * list) {
    list_t *reversed = malloc(sizeof(list_t) + list->length * sizeof(list_element_t));
    reversed->length = list->length;

    for (size_t i = 0; i < list->length; i++) {
        reversed->elements[list->length - 1 - i] = list->elements[i];
    }

    return reversed;
}

size_t length_list(list_t * list) {
    return list->length;
}

void delete_list(list_t * list) {
    free(list);
}

