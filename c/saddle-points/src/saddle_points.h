#ifndef SADDLE_POINTS_H
#define SADDLE_POINTS_H

#include <stdlib.h>
#include <stdint.h>

typedef struct {
    size_t row, column;
} saddle_point_t;

typedef struct {
    size_t count;
    saddle_point_t points[];
} saddle_points_t;

saddle_points_t *saddle_points(const size_t rows, const size_t columns,
                               uint8_t matrix[rows][columns]);
void free_saddle_points(saddle_points_t *points);

#endif
