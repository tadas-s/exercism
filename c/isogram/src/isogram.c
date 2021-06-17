#include "isogram.h"

bool is_isogram(const char phrase[]) {
    if (phrase == NULL) { return false; }

    size_t length = strlen(phrase);

    if (length == 0) { return true; }

    for (size_t i = 0; i < length - 1; i++) {
        if (isalpha(phrase[i])) {
            for(size_t j = i + 1; j < length; j++) {
                if (tolower(phrase[i]) == tolower(phrase[j])) {
                    return false;
                }
            }
        }
    }

    return true;
}