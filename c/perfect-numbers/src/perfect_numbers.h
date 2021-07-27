#ifndef PERFECT_NUMBERS_H
#define PERFECT_NUMBERS_H

#include <stdint.h>

typedef enum {
   PERFECT_NUMBER = 1,
   ABUNDANT_NUMBER = 2,
   DEFICIENT_NUMBER = 3,
   ERROR = -1
} kind;

kind classify_number(int32_t number);

#endif
