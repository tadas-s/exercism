#ifndef CIRCULAR_BUFFER_H
#define CIRCULAR_BUFFER_H

#include <stdlib.h>
#include <stdint.h>
#include <errno.h>

typedef int buffer_value_t;

typedef struct {
    size_t capacity, max_capacity, head;
    buffer_value_t *values;
} circular_buffer_t;

circular_buffer_t *new_circular_buffer(size_t capacity);
void delete_buffer(circular_buffer_t *buffer);
int16_t write(circular_buffer_t *buffer, buffer_value_t value);
int16_t read(circular_buffer_t *buffer, buffer_value_t *value);
int16_t clear_buffer(circular_buffer_t *buffer);
int16_t overwrite(circular_buffer_t *buffer, buffer_value_t value);

#endif
