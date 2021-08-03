#include "secret_handshake.h"
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>

const char *events[] = {
    [0] = "wink",
    [1] = "double blink",
    [2] = "close your eyes",
    [3] = "jump"
};

const char **commands(size_t number) {
    bool reverse = number & (1 << 4);

    const char **result = calloc(4, sizeof(char *));

    size_t offset = 0;

    for (int i = reverse ? 3 : 0; reverse ? i >= 0 : i < 4; reverse ? i-- : i++) {
        if (number & (1 << i)) {
            result[offset++] = events[i];
        }
    }

    return result;
}
