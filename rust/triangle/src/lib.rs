extern crate num;

use num::Num;

pub struct Triangle<T: Num> {
    sides: [T; 3],
}

impl<T> Triangle<T>
where
    T: Num + PartialOrd + Copy,
{
    pub fn build(sides: [T; 3]) -> Option<Triangle<T>> {
        if sides.iter().any(|&s| s <= T::zero()) {
            return None;
        }

        if (sides[0] + sides[1] < sides[2])
            || (sides[1] + sides[2] < sides[0])
            || (sides[2] + sides[0] < sides[1])
        {
            return None;
        }

        Some(Triangle { sides })
    }

    pub fn is_equilateral(&self) -> bool {
        self.sides[0] == self.sides[1] && self.sides[1] == self.sides[2]
    }

    pub fn is_isosceles(&self) -> bool {
        !self.is_equilateral()
            && ((self.sides[0] == self.sides[1])
                || (self.sides[1] == self.sides[2])
                || (self.sides[2] == self.sides[0]))
    }

    pub fn is_scalene(&self) -> bool {
        !self.is_equilateral() && !self.is_isosceles()
    }
}
