#ifndef PALINDROME_PRODUCTS_H
#define PALINDROME_PRODUCTS_H

#include <stdbool.h>

#define MAXERR 100

typedef struct factors {
   int factor_a;
   int factor_b;
   struct factors *next;
} factor_t;

struct product {
   int smallest;
   int largest;
   factor_t *factors_sm;
   factor_t *factors_lg;
   char error[MAXERR];
};

typedef struct product product_t;

product_t *get_palindrome_product(int from, int to);
bool is_palindrome(int n);
factor_t * add_factor(int a, int b, factor_t * others);
void free_product(product_t * p);
void free_factor(factor_t * p);

#endif
