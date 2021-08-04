#include "wordy.h"
#include <string.h>
#include <stdlib.h>
#include <ctype.h>

typedef enum {
    token_type_end,
    token_type_error,
    token_type_operand,
    token_type_operator_plus,
    token_type_operator_minus,
    token_type_operator_multiply,
    token_type_operator_divide,
    token_type_fluf_by
} token_type_t;

typedef struct {
    token_type_t type;
    int operand;
} token_t;

token_t parse_token(const char *token);
void calculate(token_t *tokens, bool *success, int *result);

bool answer(const char *question, int *result) {
    const size_t length = strlen(question);
    char *buff = malloc(length + 1);
    strcpy(buff, question);

    // Ditch the questionmark from the end of question
    if (length > 1 && buff[length - 1] == '?') {
        buff[length - 1] = '\0';
    }

    // Skip the "What is " if it's in the question
    if (strncmp(buff, "What is ", 8) == 0) {
        memmove(buff, buff + 8, length - 8);
    }

    size_t i = 0;
    size_t tokens_size = 10;
    token_t *tokens;
    bool success = true;

    tokens = malloc(tokens_size * sizeof(token_t));

    for (char *token = strtok(buff, " "); token; token = strtok(NULL, " "))
    {
        token_t t = parse_token(token);

        if (t.type == token_type_error) { success = false; }

        tokens[i++] = t;

        if (i == tokens_size) {
            tokens_size *= 2;
            tokens = realloc(tokens, tokens_size * sizeof(token_t));
        }

        tokens[i] = (token_t) { token_type_end };
    }

    free(buff);

    calculate(tokens, &success, result);

    free(tokens);

    return success;
}

token_t parse_token(const char *token) {
    const size_t length = strlen(token);
    bool all_numeric = true;

    for (size_t i = 0; i < length; i++) {
        if (!(isdigit(token[i]) || (i == 0 && token[i] == '-'))) {
            all_numeric = false;
        }
    }

    if (all_numeric) {
        return (token_t) { token_type_operand, atoi(token) };
    } else if (strcmp(token, "plus") == 0) {
        return (token_t){ token_type_operator_plus, -1 };
    } else if (strcmp(token, "minus") == 0) {
        return (token_t){ token_type_operator_minus, -1 };
    } else if (strcmp(token, "multiplied") == 0) {
        return (token_t){ token_type_operator_multiply, -1 };
    } else if (strcmp(token, "divided") == 0) {
        return (token_t){ token_type_operator_divide, -1 };
    } else if (strcmp(token, "by") == 0) {
        return (token_t){ token_type_fluf_by, -1 };
    } else {
        return (token_t){ token_type_error, -1 };
    }
}

void calculate(token_t *tokens, bool *success, int *result) {
    bool done = false;
    size_t stack_index = 0;

    while (*success && !done) {
        token_t t1 = tokens[stack_index++];
        token_t t2 = tokens[stack_index++];

        if (t1.type == token_type_operand && t2.type == token_type_end) {
            *result = t1.operand;
            done = true;
        } else if (t2.type == token_type_operator_plus) {
            token_t t3 = tokens[stack_index++];

            if (t3.type == token_type_operand) {
                tokens[--stack_index] = (token_t) {
                    token_type_operand,
                    t1.operand + t3.operand
                };
            } else {
                *success = false;
            }
        } else if (t2.type == token_type_operator_minus) {
            token_t t3 = tokens[stack_index++];

            if (t3.type == token_type_operand) {
                tokens[--stack_index] = (token_t) {
                    token_type_operand,
                    t1.operand - t3.operand
                };
            } else {
                *success = false;
            }
        } else if (t2.type == token_type_operator_multiply) {
            token_t t3 = tokens[stack_index++];

            if (t3.type != token_type_fluf_by) { *success = false; }

            token_t t4 = tokens[stack_index++];

            if (t4.type == token_type_operand) {
                tokens[--stack_index] = (token_t) {
                    token_type_operand,
                    t1.operand * t4.operand
                };
            } else {
                *success = false;
            }
        } else if (t2.type == token_type_operator_divide) {
            token_t t3 = tokens[stack_index++];

            if (t3.type != token_type_fluf_by) { *success = false; }

            token_t t4 = tokens[stack_index++];

            if (t4.type == token_type_operand) {
                tokens[--stack_index] = (token_t) {
                    token_type_operand,
                    t1.operand / t4.operand
                };
            } else {
                *success = false;
            }
        } else {
            *success = false;
        }
    }
}
