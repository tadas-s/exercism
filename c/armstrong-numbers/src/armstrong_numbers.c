#include "armstrong_numbers.h"
#include <math.h>

bool is_armstrong_number(int candidate) {
    int no_of_digits = floor(log10f(candidate)) + 1;
    int sum_of_pows = 0;

    for (int i = 0; i < no_of_digits; i++) {
        int digit = (candidate % (int)pow(10, i + 1)) / pow(10, i);
        sum_of_pows += pow(digit, no_of_digits);
    }

    return sum_of_pows == candidate;
}