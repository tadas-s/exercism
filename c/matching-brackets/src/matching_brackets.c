#include "matching_brackets.h"
#include <stdlib.h>
#include <string.h>

bool is_openning_bracket(const char c);
bool is_closing_bracket(const char c);
char get_closing_bracket(const char c);

bool is_paired(const char *input) {
    char *stack = malloc(strlen(input));
    size_t stack_top = 0;

    for (size_t i = 0; input[i] != '\0'; i++) {
        const char c = input[i];

        if (is_openning_bracket(c)) {
            stack[stack_top] = c;
            stack_top++;
        } else if (is_closing_bracket(c)) {
            if (stack_top > 0 && get_closing_bracket(stack[stack_top - 1]) == c) {
                stack_top--;
            } else {
                free(stack);
                return false;
            }
        }
    }

    free(stack);

    return stack_top == 0;
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