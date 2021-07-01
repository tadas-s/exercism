#include "grade_school.h"
#include <stdlib.h>
#include <string.h>
#include <strings.h>

static roster_t roster = { 0 };

void clear_roster() {
    memset(&roster, '\0', sizeof(roster_t));
}

bool add_student(char *name, uint8_t grade) {
    if (roster.count >= MAX_STUDENTS) {
        return false;
    }

    size_t i = 0;

    // Find where the specific grade starts
    while (roster.students[i].grade && roster.students[i].grade < grade) {
        i++;
    }

    // Find a spot by name within the same grade
    while (roster.students[i].grade == grade && (strcasecmp(roster.students[i].name, name) < 0)) {
        i++;
    }

    memmove(&roster.students[i+1], &roster.students[i], (MAX_STUDENTS - 1 - i) * sizeof(student_t));

    roster.students[i].grade = grade;
    roster.students[i].name = name;

    roster.count++;

    return true;
}

roster_t get_roster() {
    return roster;
}

roster_t get_grade(uint8_t desired_grade) {
    roster_t result = { 0 };

    for (size_t i = 0; i < MAX_STUDENTS; i++) {
        if (roster.students[i].grade == desired_grade) {
            result.students[result.count] = roster.students[i];
            result.count++;
        }
    }

    return result;
}
