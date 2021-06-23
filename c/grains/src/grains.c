#include "grains.h"

uint64_t square(uint8_t index) {
    if (index > 0 && index < 65) {
        return 1ull << (index - 1);
    } else {
        return 0;
    }
}

uint64_t total(void) {
    uint64_t total = 0;

    for(int i = 1; i < 65; i++) {
        total += square(i);
    }

    return total;
}
