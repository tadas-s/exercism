#include "matching_brackets.h"
#include <stdbool.h>
#include <stdlib.h>

typedef struct {
    size_t length;
    bool error;
} bracket_length_t;

bool is_openning_bracket(const char c);
bool is_closing_bracket(const char c);
char get_closing_bracket(const char c);
bracket_length_t bracketed_string_lenght(const char *input);

bool is_paired(const char *input) {
    size_t offset = 0;

    while (input[offset] != '\0') {
        const char c = input[offset];

        if (is_openning_bracket(c)) {
            bracket_length_t bl = bracketed_string_lenght(&input[offset]);

            if (bl.error) {
                return false;
            } else {
                offset += bl.length + 2;
            }
        } else if (is_closing_bracket(c)) {
            return false;
        } else {
            offset++;
        }
    }

    return true;
}

bracket_length_t bracketed_string_lenght(const char *input) {
    bracket_length_t result = { 0, false };
    size_t offset = 1;
    const char cb = get_closing_bracket(input[0]);

    while (input[offset] != '\0') {
        const char c = input[offset];

        if (is_openning_bracket(c)) {
            bracket_length_t bl = bracketed_string_lenght(&input[offset]);

            if (bl.error) {
                return (bracket_length_t){ 0, true };
            } else {
                offset += bl.length + 2;
                result.length += bl.length + 2;
            }
        } else if (c == cb) {
            return result;
        } else if (is_closing_bracket(c)) {
            return (bracket_length_t){ 0, true };
        } else {
            offset++;
            result.length++;
        }
    }

    return (bracket_length_t){ 0, true };
}

bool is_openning_bracket(const char c) {
    switch(c) {
        case '(':
        case '[':
        case '{':
            return true;
        break;

        default:
            return false;
        break;
    }
}

bool is_closing_bracket(const char c) {
    switch(c) {
        case ')':
        case ']':
        case '}':
            return true;
        break;

        default:
            return false;
        break;
    }
}

char get_closing_bracket(const char c) {
    switch(c) {
        case '(': return ')'; break;
        case '[': return ']'; break;
        case '{': return '}'; break;
        default: return '\0'; break;
    }
}