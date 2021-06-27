#include "beer_song.h"
#include <stdlib.h>
#include <string.h>
#include <stdio.h>

char *line1(unsigned int bottles);
char *line2(unsigned int bottles);

void recite(unsigned int start_bottles, unsigned int take_down, char *buffer) {
    buffer[0] = '\0';

    for(int i = start_bottles; i > ((int)start_bottles - (int)take_down); i--) {
        char *l1 = line1(i);
        strcat(buffer, l1);
        free(l1);

        char *l2 = line2(i);
        strcat(buffer, l2);
        free(l2);

        if (i > ((int)start_bottles - (int)take_down + 1)) {
            strcat(buffer, "\n");
        }
    }
}

char *line1(unsigned int bottles) {
    // ~50 char template + 2 x up to 20 digit numbers means 100 chars should be enough.
    char *buffer = malloc(100);

    if (bottles == 1 || (bottles % 10 == 1 && bottles != 11)) {
        sprintf(buffer, "%u bottle of beer on the wall, %u bottle of beer.\n", bottles, bottles);
    } else if (bottles > 1) {
        sprintf(buffer, "%u bottles of beer on the wall, %u bottles of beer.\n", bottles, bottles);
    } else {
        sprintf(buffer, "No more bottles of beer on the wall, no more bottles of beer.\n");
    }

    return buffer;
}

char *line2(unsigned int bottles) {
    char *buffer = malloc(100);

    if (bottles > 1 && !(bottles == 12 || bottles % 10 == 2)) {
        sprintf(buffer, "Take one down and pass it around, %u bottles of beer on the wall.\n", bottles - 1);
    } else if (bottles > 1) {
        sprintf(buffer, "Take one down and pass it around, %u bottle of beer on the wall.\n", bottles - 1);
    } else if (bottles == 1) {
        sprintf(buffer, "Take it down and pass it around, no more bottles of beer on the wall.\n");
    } else {
        sprintf(buffer, "Go to the store and buy some more, 99 bottles of beer on the wall.\n");
    }

    return buffer;
}