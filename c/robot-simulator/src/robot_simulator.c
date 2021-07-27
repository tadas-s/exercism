#include "robot_simulator.h"
#include <stddef.h>

robot_status_t turn_left(robot_status_t robot);
robot_status_t turn_right(robot_status_t robot);
robot_status_t advance(robot_status_t robot);

robot_status_t robot_create(robot_direction_t direction, int x, int y) {
    return (robot_status_t){ direction, { x, y } };
}

void robot_move(robot_status_t * robot, const char *commands) {
    for (size_t i = 0; commands[i] != '\0'; i++) {
        switch (commands[i]) {
            case 'L': *robot = turn_left(*robot); break;
            case 'R': *robot = turn_right(*robot); break;
            case 'A': *robot = advance(*robot); break;
            default: break;
        }
    }
}

robot_status_t turn_left(robot_status_t robot) {
    switch (robot.direction) {
        case DIRECTION_NORTH: robot.direction = DIRECTION_WEST; break;
        case DIRECTION_EAST: robot.direction = DIRECTION_NORTH; break;
        case DIRECTION_SOUTH: robot.direction = DIRECTION_EAST; break;
        case DIRECTION_WEST: robot.direction = DIRECTION_SOUTH; break;
        default: break;
    }

    return robot;
}

robot_status_t turn_right(robot_status_t robot) {
    switch (robot.direction) {
        case DIRECTION_NORTH: robot.direction = DIRECTION_EAST; break;
        case DIRECTION_EAST: robot.direction = DIRECTION_SOUTH; break;
        case DIRECTION_SOUTH: robot.direction = DIRECTION_WEST; break;
        case DIRECTION_WEST: robot.direction = DIRECTION_NORTH; break;
        default: break;
    }

    return robot;
}

robot_status_t advance(robot_status_t robot) {
    switch (robot.direction) {
        case DIRECTION_NORTH: robot.position.y += 1; break;
        case DIRECTION_EAST: robot.position.x += 1; break;
        case DIRECTION_SOUTH: robot.position.y -= 1; break;
        case DIRECTION_WEST: robot.position.x -= 1; break;
        default: break;
    }

    return robot;
}