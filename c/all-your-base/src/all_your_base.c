#include "all_your_base.h"
#include <math.h>
#include <string.h>

size_t rebase(int8_t digits[DIGITS_ARRAY_SIZE], int16_t input_base,
              int16_t output_base, size_t input_length) {
    int n = 0;

    if (input_base < 2 || output_base < 2) { return 0; }

    for (size_t i = 0; i < input_length; i++) {
        if (digits[i] < 0 || digits[i] >= input_base) { return 0; }

        n += digits[i] * pow(input_base, input_length - (i + 1));
    }

    if (n == 0) { return 0; }

    int msd = -1; // index of most significant digit in the output

    for (size_t i = 0; i < DIGITS_ARRAY_SIZE; i++) {
        const int div = pow(output_base, DIGITS_ARRAY_SIZE - i - 1);
        digits[i] = n / div;

        if (msd == -1 && digits[i] > 0) {
            msd = i;
        }

        n = n % div;
    }

    // Move number forward so most significant "digit" is at zero index
    // of the digits array.
    memmove(digits, &digits[msd], DIGITS_ARRAY_SIZE - msd);

    return DIGITS_ARRAY_SIZE - msd;
}