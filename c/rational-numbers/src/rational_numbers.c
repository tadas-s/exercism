#include "rational_numbers.h"
#include <stdlib.h>
#include <math.h>

rational_t reduce(rational_t n) {
    if (n.numerator == 0) {
        return (rational_t){ 0, 1 };
    } else {
        // Normalize the sign
        if (n.denominator < 0) {
            n = (rational_t){ -n.numerator, -n.denominator };
        }

        // Find a divider if any
        int max_divider = fmin(abs(n.numerator), abs(n.denominator));
        for (int divider = max_divider; divider > 1; divider--) {
            if (n.numerator % divider == 0 && n.denominator % divider == 0) {
                return (rational_t){ n.numerator / divider, n.denominator / divider };
            }
        }

        return n; // no divider found - just return the original
    }
}

rational_t add(rational_t a, rational_t b) {
    return reduce((rational_t) {
        (a.numerator * b.denominator + b.numerator * a.denominator),
        (a.denominator * b.denominator)
    });
}

rational_t subtract(rational_t a, rational_t b) {
    return reduce((rational_t) {
        (a.numerator * b.denominator - b.numerator * a.denominator),
        (a.denominator * b.denominator)
    });
}

rational_t multiply(rational_t a, rational_t b) {
    return reduce((rational_t) {
        (a.numerator * b.numerator), (a.denominator * b.denominator)
    });
}

rational_t divide(rational_t a, rational_t b) {
    return reduce((rational_t) {
        (a.numerator * b.denominator), (a.denominator * b.numerator)
    });
}

rational_t absolute(rational_t n) {
    return reduce((rational_t) {
        abs(n.numerator), abs(n.denominator)
    });
}

rational_t exp_rational(rational_t n, uint16_t e) {
    return reduce((rational_t) {
        pow(n.numerator, e), pow(n.denominator, e)
    });
}

float exp_real(uint16_t a, rational_t b) {
    return pow(a, (float)b.numerator / (float)b.denominator);
}
