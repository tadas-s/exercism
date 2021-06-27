#include "bob.h"
#include <stdbool.h>
#include <string.h>
#include <ctype.h>

const char response_calm_down[] = "Calm down, I know what I'm doing!";
const char response_chill_out[] = "Whoa, chill out!";
const char response_whatever[] = "Whatever.";
const char response_sure[] = "Sure.";
const char response_fine[] = "Fine. Be that way!";

const char *hey_bob(char *greeting) {
    size_t length = strlen(greeting);
    bool is_question;
    bool has_alphabetic_chars = false;
    bool has_lowercase_chars = false;
    bool has_uppercase_chars = false;
    bool is_yelling;
    bool whitespace_only = true;
    char last_character = '\0';

    for (size_t i = 0; i < length; i++) {
        if (isalpha(greeting[i])) {
            has_alphabetic_chars = true;

            if(greeting[i] == toupper(greeting[i])) {
                has_uppercase_chars = true;
            } else {
                has_lowercase_chars = true;
            }
        }

        if (!isspace(greeting[i])) {
            last_character = greeting[i];
            whitespace_only = false;
        }
    }

    is_question = last_character == '?';
    is_yelling = has_alphabetic_chars && has_uppercase_chars && !has_lowercase_chars;

    if (whitespace_only) {
        return response_fine;
    } else if (is_yelling && is_question) {
        return response_calm_down;
    } else if (is_yelling) {
        return response_chill_out;
    } else if (is_question) {
        return response_sure;
    }

    return response_whatever;
}
