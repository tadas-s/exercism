#include "palindrome_products.h"
#include <stdlib.h>
#include <stdio.h>

product_t *get_palindrome_product(int from, int to) {
    product_t * result = malloc(sizeof(product_t));
    *result = (product_t){ -1, -1, NULL, NULL, "" };

    if (from > to) {
        snprintf(result->error, MAXERR,
                 "invalid input: min is %d and max is %d", from, to);
        return result;
    }

    for (int a = from; a <= to; a++) {
        for (int b = a; b <= to; b++) {
            const int p = a * b;

            if (is_palindrome(p)) {
                if (result->smallest == -1 || p < result->smallest) {
                    result->smallest = p;
                    free_factor(result->factors_sm);
                    result->factors_sm = NULL;
                    result->factors_sm = add_factor(a, b, result->factors_sm);
                } else if (p == result->smallest) {
                    result->factors_sm = add_factor(a, b, result->factors_sm);
                }

                if (result->largest == -1 || p > result->largest) {
                    result->largest = p;
                    free_factor(result->factors_lg);
                    result->factors_lg = NULL;
                    result->factors_lg = add_factor(a, b, result->factors_lg);
                } else if (p == result->largest) {
                    result->factors_lg = add_factor(a, b, result->factors_lg);
                }
            }
        }
    }

    if (result->smallest == -1) {
        snprintf(result->error, MAXERR,
                 "no palindrome with factors in the range %d to %d", from, to);
    }

    return result;
}

bool is_palindrome(int n) {
    char buff[31];

    int printed = snprintf(buff, 30, "%d", n);

    for (int i = 0; i < (printed / 2 + printed % 2); i++) {
        if (buff[i] != buff[printed - i - 1]) {
            return false;
        }
    }

    return true;
}

factor_t * add_factor(int a, int b, factor_t * others) {
    if (others) {
        factor_t * tail = others;
        while (tail->next) { tail = tail->next; }
        tail->next = malloc(sizeof(factor_t));
        *tail->next = (factor_t){ a, b, NULL };
        return others;
    } else {
        factor_t * result = malloc(sizeof(factor_t));
        *result = (factor_t){ a, b, NULL };
        return result;
    }
}

void free_product(product_t * p) {
    free_factor(p->factors_sm);
    free_factor(p->factors_lg);
    free(p);
}

void free_factor(factor_t * p) {
    if (!p) { return; }

    if (p->next) {
        free_factor(p->next);
    }

    free(p);
}