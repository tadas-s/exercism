#include "scrabble_score.h"
#include <stdint.h>
#include <stddef.h>
#include <ctype.h>

static uint8_t points[255] = {
    ['a'] = 1, ['e'] = 1, ['i'] = 1, ['o'] = 1,
    ['u'] = 1, ['l'] = 1, ['n'] = 1, ['r'] = 1,
    ['s'] = 1, ['t'] = 1, ['d'] = 2, ['g'] = 2,
    ['b'] = 3, ['c'] = 3, ['m'] = 3, ['p'] = 3,
    ['f'] = 4, ['h'] = 4, ['v'] = 4, ['w'] = 4,
    ['y'] = 4, ['k'] = 5, ['j'] = 8, ['x'] = 8,
    ['q'] = 10, ['z']= 10
};

unsigned int score(const char *word) {
    int result = 0;

    for (size_t i = 0; word[i] != '\0'; i++) {
        if (points[tolower(word[i])] == 0) { return 0; }
        result += points[tolower(word[i])];
    }

    return result;
}