#include "circular_buffer.h"

circular_buffer_t *new_circular_buffer(size_t capacity) {
    circular_buffer_t *buffer = malloc(sizeof(circular_buffer_t));
    buffer->max_capacity = capacity;
    buffer->capacity = 0;
    buffer->values = malloc(sizeof(buffer_value_t) * capacity);
    buffer->head = 0;
    return buffer;
}

void delete_buffer(circular_buffer_t *buffer) {
    free(buffer->values);
    free(buffer);
}

int16_t write(circular_buffer_t *buffer, buffer_value_t value) {
    errno = 0;

    if (buffer->capacity == buffer->max_capacity) {
        errno = ENOBUFS;
        return EXIT_FAILURE;
    }

    return overwrite(buffer, value);
}

int16_t read(circular_buffer_t *buffer, buffer_value_t *value) {
    errno = 0;

    if (buffer->capacity == 0) {
        errno = ENODATA;
        return EXIT_FAILURE;
    }

    if (buffer->capacity > buffer->head) {
        *value = buffer->values[buffer->max_capacity - (buffer->capacity - buffer->head)];
    } else {
        *value = buffer->values[buffer->head - buffer->capacity];
    }

    buffer->capacity--;

    return EXIT_SUCCESS;
}

int16_t clear_buffer(circular_buffer_t *buffer) {
    errno = 0;

    buffer->head = 0;
    buffer->capacity = 0;

    return EXIT_SUCCESS;
}

int16_t overwrite(circular_buffer_t *buffer, buffer_value_t value) {
    errno = 0;

    buffer->values[buffer->head] = value;

    if (buffer->capacity < buffer->max_capacity) {
        buffer->capacity++;
    }

    if (buffer->head < (buffer->max_capacity - 1)) {
        buffer->head++;
    } else {
        buffer->head = 0;
    }

    return EXIT_SUCCESS;
}