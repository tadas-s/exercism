#include "clock.h"
#include <stdio.h>
#include <string.h>

clock_t clock_create(int hour, int minute) {
    clock_t result;

    int time = hour * 60 + minute;

    if (time < 0) {
        time = time - (time / (24 * 60) - 1) * 24 * 60;
    }

    const int normalized_hour = (time / 60) % 24;
    const int normalized_minute = time % 60;

    snprintf(result.text, MAX_STR_LEN, "%02d:%02d", normalized_hour, normalized_minute);

    return result;
}

clock_t clock_add(clock_t clock, int minute_add) {
    int hour, minute, time;

    sscanf(clock.text, "%d:%d", &hour, &minute);

    time = hour * 60 + minute + minute_add;

    return clock_create(0, time);
}

clock_t clock_subtract(clock_t clock, int minute_subtract) {
    int hour, minute, time;

    sscanf(clock.text, "%d:%d", &hour, &minute);

    time = hour * 60 + minute - minute_subtract;

    return clock_create(0, time);
}

bool clock_is_equal(clock_t a, clock_t b) {
    return strncmp(a.text, b.text, MAX_STR_LEN) == 0;
}