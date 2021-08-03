#include "minesweeper.h"
#include <stdlib.h>
#include <string.h>

int mines_around(const char **minefield, const size_t rows, const size_t columns,
                 const size_t row, const size_t column);

char **annotate(const char **minefield, const size_t rows) {
    if (rows == 0) { return NULL; }

    // +1 for end "zero" marker
    char **result = calloc(rows + 1, sizeof(char *));
    const size_t columns = rows > 0 ? strlen(minefield[0]) : 0;

    for (size_t i = 0; i < rows; i++) {
        result[i] = calloc(columns, sizeof(char) + 1);

        for (size_t j = 0; j < columns; j++) {
            if (minefield[i][j] == '*') {
                result[i][j] = '*';
            } else {
                int mines = mines_around(minefield, rows, columns, i, j);

                if (mines) {
                    result[i][j] = '0' + mines;
                } else {
                    result[i][j] = ' ';
                }
            }
        }
    }

    result[rows] = NULL; // "zero" marker for the free_annotation procedure

    return result;
}

void free_annotation(char **annotation) {
    for (size_t i = 0; annotation[i] != NULL; i++) {
        free(annotation[i]);
    }

    free(annotation);
}

int mines_around(const char **minefield, const size_t rows, const size_t columns,
                 const size_t row, const size_t column)
{
    int mines = 0;

    for (int i = -1; i <= 1; i++) {
        if ((int)row + i < 0) continue;
        if ((int)row + i >= (int)rows) continue;

        for (int j = -1; j <= 1; j++) {
            if (i == 0 && j == 0) continue;
            if ((int)column + j < 0) continue;
            if ((int)column + j >= (int)columns) continue;

            if (minefield[row + i][column + j] == '*') {
                mines++;
            }
        }
    }

    return mines;
}
