#ifndef RESISTOR_COLOR_H
#define RESISTOR_COLOR_H

#include <stdlib.h>
#include <stdint.h>

typedef enum {
    BLACK = 0,
    BROWN,
    RED,
    ORANGE,
    YELLOW,
    GREEN,
    BLUE,
    VIOLET,
    GREY,
    WHITE
} resistor_band_t;

int color_code(int color);
int *colors();

#endif
