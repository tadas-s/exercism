#include "say.h"
#include <stdlib.h>
#include <string.h>
#include <math.h>

int say_1_to_999(int n, char *dest);
void strcat_with_space(char *dest, const char *string);

static char *names1[] = {
    [1] = "one", [2] = "two", [3] = "three", [4] = "four", [5] = "five",
    [6] = "six", [7] = "seven", [8] = "eigh", [9] = "nine",
    [10] = "ten", [11] = "eleven", [12] = "twelve", [13] = "thirteen",
    [14] = "fourteen", [15] = "fifteen", [16] = "sixteen", [17] = "seventeen",
    [18] = "eighteen", [19] = "nineteen",
};

static char *names3[] = {
    [2] = "twenty", [3] = "thirty", [4] = "forty", [5] = "fifty",
    [6] = "sixty", [7] = "seventy", [8] = "eighty", [9] = "ninety"
};

static char *names4[] = {
    [0] = "",
    [1] = "thousand",
    [2] = "million",
    [3] = "billion"
};

int say(int64_t n, char **dest) {
    if (n < 0 || n > 999999999999) { return -1; }

    *dest = calloc(1000, sizeof(char));

    if (n == 0) {
        strcat_with_space(*dest, "zero");
    } else {
        for (int i = 3; i >= 0; i--) {
            int64_t units = n % (int64_t)pow(10, 3 * (i + 1)) / (int64_t)pow(10, 3*i);

            if (units > 0) {
                say_1_to_999(units, *dest);
                strcat_with_space(*dest, names4[i]);
            }
        }
    }

    *dest = realloc(*dest, strlen(*dest) + 1);

    return 0;
}

int say_1_to_999(int n, char *dest) {
    if (n < 1 || n > 999) { return -1; }

    if (n / 100) {
        strcat_with_space(dest, names1[n / 100]);
        strcat_with_space(dest, "hundred");
    }

    if (n % 100) {
        if (n % 100 < 20) {
            strcat_with_space(dest, names1[n % 100]);
        } else if (n % 10) {
            strcat_with_space(dest, names3[(n % 100) / 10]);
            strcat(dest, "-");
            strcat(dest, names1[n % 10]);
        } else {
            strcat_with_space(dest, names3[n / 10]);
        }
    }

    return 0;
}

void strcat_with_space(char *dest, const char *string) {
    if (strlen(dest) > 0 && strlen(string) > 0) {
        strcat(dest, " ");
    }

    strcat(dest, string);
}