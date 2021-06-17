#ifndef RESISTOR_COLOR_TRIO_H
#define RESISTOR_COLOR_TRIO_H

#include <math.h>

typedef enum {
    BLACK = 0, BROWN, RED, ORANGE, YELLOW,
    GREEN, BLUE, VIOLET, GREY, WHITE
} resistor_band_t;

typedef enum {
    OHMS, KILOOHMS
} resistor_unit_t;

typedef struct {
    int value;
    int unit;
} resistor_value_t;

resistor_value_t color_code(resistor_band_t colors[]);

#endif
