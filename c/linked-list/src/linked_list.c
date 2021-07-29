#include "linked_list.h"
#include <stdlib.h>

struct list_node {
   struct list_node *prev, *next;
   ll_data_t data;
};

struct list {
   struct list_node *first, *last;
};

struct list *list_create(void) {
   struct list *l = malloc(sizeof(struct list));
   *l = (struct list){ NULL, NULL };
   return l;
}

bool list_is_empty(const struct list *list) {
   return list->first == NULL;
}

bool list_push(struct list *list, ll_data_t item_data) {
   struct list_node *n = malloc(sizeof(struct list_node));

   if (!n) { return false; }

   *n = (struct list_node){ NULL, NULL, item_data };

   if (list_is_empty(list)) {
      list->first = list->last = n;
   } else {
      list->last->next = n;
      n->prev = list->last;
      list->last = n;
   }

   return true;
}

ll_data_t list_pop(struct list *list) {
   struct list_node *popped = list->last;
   ll_data_t item = popped->data;

   list->last = list->last->prev;

   if (list->last) {
      list->last->next = NULL;
   } else {
      list->first = NULL;
   }

   free(popped);

   return item;
}

bool list_unshift(struct list *list, ll_data_t item_data) {
   struct list_node *n = malloc(sizeof(struct list_node));

   if (!n) { return false; }

   *n = (struct list_node){ NULL, NULL, item_data };

   if (list_is_empty(list)) {
      list->first = list->last = n;
   } else {
      list->first->prev = n;
      n->next = list->first;
      list->first = n;
   }

   return true;
}

ll_data_t list_shift(struct list *list) {
   struct list_node *shifted = list->first;
   ll_data_t item = shifted->data;

   list->first = list->first->next;

   if (list->first) {
      list->first->prev = NULL;
   } else {
      list->last = NULL;
   }

   free(shifted);

   return item;
}

void list_destroy(struct list *list) {
   struct list_node *to_free = list->first;

   while (to_free) {
      struct list_node *this_node = to_free;
      to_free = to_free->next;
      free(this_node);
   }

   free(list);
}