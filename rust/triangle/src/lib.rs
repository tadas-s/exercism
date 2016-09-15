extern crate num;

use num::Num;

pub struct Triangle<T: Num> {
    sides: [T; 3],
}

impl<T> Triangle<T> where T: Num + PartialOrd + Copy {
    pub fn build(s: [T; 3]) -> Result<Triangle<T>, ()>  {
        let t: Triangle<T> = Triangle { sides: s };
        
        if s.iter().all(|&s| s <= T::zero()) {
            return Err(());
        }

        if (s[0] + s[1] < s[2]) || (s[1] + s[2] < s[0]) || (s[2] + s[0] < s[1]) {
            return Err(());
        }

        return Ok(t);
    }

    pub fn is_equilateral(&self) -> bool {
        self.sides[0] == self.sides[1] && self.sides[1] == self.sides[2]
    }

    pub fn is_isosceles(&self) -> bool {
        !self.is_equilateral() && (
            (self.sides[0] == self.sides[1]) || (self.sides[1] == self.sides[2]) || (self.sides[2] == self.sides[0])
        )
    }

    pub fn is_scalene(&self) -> bool {
        !self.is_equilateral() && !self.is_isosceles()
    }
}
