use std::fmt::{Display, Error, Formatter, Result};

const UNITS: &'static [[&'static str; 10]; 4] = &[
    ["", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"],
    ["", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"],
    ["", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"],
    ["", "M", "MM", "MMM", "", "", "", "", "", ""],
];

pub struct Roman {
    number: u32,
}

impl Display for Roman {
    fn fmt(&self, f: &mut Formatter<'_>) -> Result {
        if self.number == 0 || self.number > 3999 {
            return Err(Error);
        }

        write!(
            f,
            "{}{}{}{}",
            UNITS[3][(self.number / 1000) as usize],
            UNITS[2][(self.number % 1000 / 100) as usize],
            UNITS[1][(self.number % 100 / 10) as usize],
            UNITS[0][(self.number % 10) as usize],
        )
    }
}

impl From<u32> for Roman {
    fn from(num: u32) -> Self {
        Roman { number: num }
    }
}
