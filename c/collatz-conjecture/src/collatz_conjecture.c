#include "collatz_conjecture.h"

int steps(int n) {
    int s = 0;

    if (n < 1) {
        return ERROR_VALUE;
    }

    while (n != 1) {
        if (n % 2 == 0) {
            n = n / 2;
        } else {
            n = 3 * n + 1;
        }

        s++;
    };

    return s;
}
