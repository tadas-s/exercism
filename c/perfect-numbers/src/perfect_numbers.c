#include "perfect_numbers.h"

kind classify_number(int32_t number) {
    int32_t sum = 0;

    if (number < 1) {
        return ERROR;
    }

    for (int32_t d = 1; d <= (number / 2); d++) {
        if (number % d == 0) {
            sum += d;
        }
    }

    if (sum == number) {
        return PERFECT_NUMBER;
    } else if (sum > number) {
        return ABUNDANT_NUMBER;
    } else {
        return DEFICIENT_NUMBER;
    }
}
