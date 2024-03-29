#include "square_root.h"

int32_t square_root(int32_t num) {
    int32_t res = 0;
    int32_t bit = 1 << 30;

    while (bit > num) {
        bit >>= 2;
    }

    while (bit != 0) {
        if (num >= res + bit) {
            num -= res + bit;
            res = (res >> 1) + bit;
        } else {
            res >>= 1;
        }

        bit >>= 2;
    }

    return res;
}