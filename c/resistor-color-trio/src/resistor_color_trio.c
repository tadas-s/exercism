#include "resistor_color_trio.h"

resistor_value_t color_code(resistor_band_t colors[]) {
    resistor_value_t result;
    result.value = (10 * colors[0] + colors[1]) * pow(10, colors[2]);
    result.unit = OHMS;

    if (0 == result.value % 1000) {
        result.value = result.value / 1000;
        result.unit += KILOOHMS;
    }

    return result;
}