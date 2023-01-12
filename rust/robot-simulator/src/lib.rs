// The code below is a stub. Just enough to satisfy the compiler.
// In order to pass the tests you can add-to or change any of this code.

use core::panic;
use Direction::*;

#[derive(PartialEq, Eq, Debug)]
pub enum Direction {
    North,
    East,
    South,
    West,
}

pub struct Robot {
    x: i32,
    y: i32,
    direction: Direction,
}

impl Robot {
    pub fn new(x: i32, y: i32, d: Direction) -> Self {
        Self { x: x, y: y, direction: d}
    }

    #[must_use]
    pub fn turn_right(self) -> Self {
        match self.direction {
            North => Self { direction: East, ..self },
            East => Self { direction: South, ..self },
            South => Self { direction: West, ..self },
            West => Self { direction: North, ..self },
        }
    }

    #[must_use]
    pub fn turn_left(self) -> Self {
        match self.direction {
            North => Self { direction: West, ..self },
            East => Self { direction: North, ..self },
            South => Self { direction: East, ..self },
            West => Self { direction: South, ..self },
        }
    }

    #[must_use]
    pub fn advance(self) -> Self {
        match self.direction {
            North => Self { y: self.y + 1, ..self },
            East => Self { x: self.x + 1, ..self },
            South => Self { y: self.y - 1, ..self },
            West => Self { x: self.x - 1, ..self },
        }
    }

    #[must_use]
    pub fn instructions(self, instructions: &str) -> Self {
        instructions.chars().fold(self, |robot, instruction| {
            match instruction {
                'R' => robot.turn_right(),
                'L' => robot.turn_left(),
                'A' => robot.advance(),
                _ => panic!("Unknown instruction: '{}'", instruction),
            }
        })
    }

    pub fn position(&self) -> (i32, i32) {
        (self.x, self.y)
    }

    pub fn direction(&self) -> &Direction {
        &self.direction
    }
}
