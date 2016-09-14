pub struct Triangle {
    sides: [isize; 3],
}

impl Triangle {
    pub fn build(s: [isize; 3]) -> Result<Triangle, ()>  {
        let t: Triangle = Triangle { sides: s };
        
        if s.iter().all(|&s| s <= 0) {
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
        self.sides[0] != self.sides[1] && self.sides[1] != self.sides[2] && self.sides[2] != self.sides[0]
    }
}
