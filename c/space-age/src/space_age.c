#include "space_age.h"
#include <math.h>

float age(planet_t planet, int64_t seconds) {
    switch(planet) {
        case MERCURY: return seconds / 7600543.8; break;
        case VENUS: return seconds / 19414149.1; break;
        case EARTH: return seconds / 31557600.0; break;
        case MARS: return seconds / 59354032.7; break;
        case JUPITER: return seconds / 374355659.1; break;
        case SATURN: return seconds / 929292362.9; break;
        case URANUS: return seconds / 2651370019.3; break;
        case NEPTUNE: return seconds / 5200418560.0; break;
    };

    return NAN;
}