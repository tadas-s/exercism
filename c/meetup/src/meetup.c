#include "meetup.h"
#include <time.h>
#include <string.h>

int meetup_day_of_month(unsigned int year, unsigned int month, const char *week,
                        const char *day_of_week) {

    int tm_wday;
    int nth_weekday;
    int nth_counter = 0;
    int nth_last = -1;

    if (strcmp("Sunday", day_of_week) == 0) {
        tm_wday = 0;
    } else if (strcmp("Monday", day_of_week) == 0) {
        tm_wday = 1;
    } else if (strcmp("Tuesday", day_of_week) == 0) {
        tm_wday = 2;
    } else if (strcmp("Wednesday", day_of_week) == 0) {
        tm_wday = 3;
    } else if (strcmp("Thursday", day_of_week) == 0) {
        tm_wday = 4;
    } else if (strcmp("Friday", day_of_week) == 0) {
        tm_wday = 5;
    } else if (strcmp("Saturday", day_of_week) == 0) {
        tm_wday = 6;
    } else {
        return -1;
    }

    if (strcmp("first", week) == 0) {
        nth_weekday = 0;
    } else if (strcmp("second", week) == 0) {
        nth_weekday = 1;
    } else if (strcmp("third", week) == 0) {
        nth_weekday = 2;
    } else if (strcmp("fourth", week) == 0) {
        nth_weekday = 3;
    } else if (strcmp("fifth", week) == 0) {
        nth_weekday = 4;
    } else if (strcmp("last", week) == 0) {
        nth_weekday = -1;
    }

    for(int day_of_month = 1; day_of_month <= 31; day_of_month++) {
        struct tm *t;
        time_t timestamp;

        time(&timestamp);

        t = localtime(&timestamp);

        t->tm_sec = t->tm_min = 0;
        t->tm_hour = 12;
        // tm_year is the number of full calendar years since 1900
        t->tm_year = year - 1900;
        // tm_mon is 0 based
        t->tm_mon = month - 1;
        // while tm_day starts with 1
        t->tm_mday = day_of_month;

        time_t err = mktime(t);

        if (err < 0) {
            return -1;
        }

        // If month got bumped by mktime - we're trying to loop through
        // a month with less than 31 days.
        if (t->tm_mon != (int)(month - 1)) {
            break;
        }

        if (strcmp("teenth", week) == 0) {
            if (t->tm_mday >= 13 && t->tm_mday <= 19 && t->tm_wday == tm_wday) {
                return t->tm_mday;
            }
        } else if (nth_weekday >= 0 && nth_weekday <= 4 && t->tm_wday == tm_wday) {
            if (nth_weekday == nth_counter) {
                return t->tm_mday;
            }
            nth_counter++;
        } else if (nth_weekday == -1 && t->tm_wday == tm_wday) {
            nth_last = t->tm_mday;
        }
    }

    if (nth_weekday == -1) {
        return nth_last;
    } else {
        return -1;
    }
}