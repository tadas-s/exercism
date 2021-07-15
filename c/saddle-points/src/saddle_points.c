#include "saddle_points.h"
#include <math.h>
#include <string.h>

saddle_points_t *saddle_points(const size_t rows, const size_t columns, uint8_t matrix[rows][columns]) {
    uint8_t column_minimums[columns];
    uint8_t row_maximums[rows];

    memset(column_minimums, 255, columns);
    memset(row_maximums, 0, rows);

    saddle_points_t *result = malloc(sizeof(saddle_points_t));
    result->count = 0;

    for (size_t row = 0; row < rows; row++) {
        for (size_t column = 0; column < columns; column++) {
            column_minimums[column] = fmin(column_minimums[column], matrix[row][column]);
            row_maximums[row] = fmax(row_maximums[row], matrix[row][column]);
        }
    }

    for (size_t row = 0; row < rows; row++) {
        for (size_t column = 0; column < columns; column++) {
            if (matrix[row][column] == column_minimums[column]
                && matrix[row][column] == row_maximums[row]) {
                    result->count++;
                    result = realloc(result, sizeof(saddle_points_t)
                                     + result->count * sizeof(saddle_point_t));
                    result->points[result->count - 1] =
                        (saddle_point_t){ row + 1, column + 1};
                }
        }
    }

    return result;
}

void free_saddle_points(saddle_points_t *points) {
    free(points);
}

