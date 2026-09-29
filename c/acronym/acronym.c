#include "acronym.h"
#include <stddef.h>
#include <ctype.h>
#include <stdlib.h>
#include <string.h>

typedef enum {
    SEPARATOR,
    WORD
} acronym_scan_state_t;

char *abbreviate(const char *phrase) {
    if (phrase == NULL) { return NULL; }
    if (strcmp(phrase, "") == 0) { return NULL; }

    char *abbreviation = NULL;
    size_t abbreviation_size = 1;
    acronym_scan_state_t state = SEPARATOR;

    for (size_t i = 0; i < strlen(phrase); i++) {
        char current_character = phrase[i];

        switch (state) {
        case SEPARATOR:
            if (isalpha(current_character)) {
                abbreviation_size += 1;
                abbreviation = realloc(abbreviation, abbreviation_size);
                abbreviation[abbreviation_size - 2] = toupper(current_character);
                abbreviation[abbreviation_size - 1] = '\0';
                state = WORD;
            }
            break;

        case WORD:
            if (!isalpha(current_character) && current_character != '\'') {
                state = SEPARATOR;
            }
            break;

        default: break;
        }
    }

    return abbreviation;
}