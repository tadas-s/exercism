#include "binary_search_tree.h"
#include <stdlib.h>
#include <stdbool.h>

void sorted_data_r(int **data, size_t *data_count, size_t *max_count,
                   node_t *tree);

node_t *build_tree(int *tree_data, size_t tree_data_len) {
    if (tree_data_len < 1) { return NULL; }

    node_t *tree = NULL;

    for (size_t i = 0; i < tree_data_len; i++) {
        node_t *n = calloc(1, sizeof(node_t));
        n->data = tree_data[i];

        if (!tree) {
            tree = n;
        } else {
            node_t *parent_node = tree;
            bool placed = false;

            do {
                if (n->data > parent_node->data && !parent_node->right) {
                    parent_node->right = n;
                    placed = true;
                } else if (n->data > parent_node->data && parent_node->right) {
                    parent_node = parent_node->right;
                } else if (n->data <= parent_node->data && !parent_node->left) {
                    parent_node->left = n;
                    placed = true;
                } else if (n->data <= parent_node->data && parent_node->left) {
                    parent_node = parent_node->left;
                }

            } while(!placed);
        }
    }

    return tree;
}

void free_tree(node_t * tree) {
    if (tree->left) { free_tree(tree->left); }
    if (tree->right) { free_tree(tree->right); }

    free(tree);

    return;
}

void sorted_data_r(int **data, size_t *data_count, size_t *max_count,
                   node_t *tree)
{
    if (tree->left) {
        sorted_data_r(data, data_count, max_count, tree->left);
    }

    (*data)[(*data_count)++] = tree->data;

    if ((*data_count) == (*max_count)) {
        *max_count *= 2;
        *data = realloc(*data, (*max_count) * sizeof(int));
    }

    if (tree->right) {
        sorted_data_r(data, data_count, max_count, tree->right);
    }
}

int *sorted_data(node_t * tree) {
    size_t data_count = 0;
    size_t max_count = 50;
    int *data = calloc(max_count, sizeof(int));

    sorted_data_r(&data, &data_count, &max_count, tree);

    data = realloc(data, data_count * sizeof(int));

    return data;
}