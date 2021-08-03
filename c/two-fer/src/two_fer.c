#include "two_fer.h"
#include <stdio.h>

void two_fer(char *buffer, const char *name) {
    const char *you = name ? name : "you";
    sprintf(buffer, "One for %s, one for me.", you);
}
