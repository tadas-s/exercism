#include "pascals_triangle.h"
#include <stdlib.h>
#include <string.h>

void free_triangle(uint8_t ** triangle, size_t rows) {
    const size_t allocated = rows < 1 ? 1 : rows;

    for (size_t i = 0; i < allocated; i++) {
        free(triangle[i]);
    }

    free(triangle);
}

uint8_t **create_triangle(size_t rows) {
    const size_t allocate = rows < 1 ? 1 : rows;

    uint8_t **triangle = malloc(allocate * sizeof(uint8_t *));

    for (size_t i = 0; i < allocate; i++) {
        triangle[i] = calloc(allocate, sizeof(uint8_t));
    }

    if (rows < 1) {
        return triangle;
    }

    triangle[0][0] = 1;

    for (size_t row = 1; row < rows; row++) {
        for(size_t column = 0; column < rows; column++) {
            const uint8_t a = column > 0 ? triangle[row - 1][column - 1] : 0;
            const uint8_t b = triangle[row - 1][column];
            triangle[row][column] = a + b;
        }
    }

    return triangle;
}
