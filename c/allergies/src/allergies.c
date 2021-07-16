#include "allergies.h"
#include <string.h>

bool is_allergic_to(const allergen_t allergen, const uint16_t overall_score) {
    return (1 << allergen) & overall_score;
}

allergen_list_t get_allergens(const uint16_t overall_score) {
    allergen_list_t list;
    list.count = 0;

    for (allergen_t allergen = ALLERGEN_EGGS; allergen < ALLERGEN_COUNT; allergen++) {
        if (is_allergic_to(allergen, overall_score)) {
            list.count++;
            list.allergens[allergen] = true;
        } else {
            list.allergens[allergen] = false;
        }
    }

    return list;
}
