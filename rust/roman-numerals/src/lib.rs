pub struct Roman(pub isize);

impl From<isize> for Roman {
    fn from(n: isize) -> Roman {
        Roman(n)
    }
}

impl ToString for Roman {
    fn to_string(&self) -> String {
        let ones: isize = self.0 % 10;
        let tens: isize = self.0 / 10 % 10;
        let hundreds: isize = self.0 / 100 % 10;
        let thousands: isize = self.0 / 1000 % 10;
        
        match thousands {
            1 => "M",
            2 => "MM",
            3 => "MMM",
            _ => "",
        }.to_string() +
        match hundreds {
            1 => "C",
            2 => "CC",
            3 => "CCC",
            4 => "CD",
            5 => "D",
            6 => "DC",
            7 => "DCC",
            8 => "DCCC",
            9 => "CM",
            _ => "",
        } +
        match tens {
            1 => "X",
            2 => "XX",
            3 => "XXX",
            4 => "XL",
            5 => "L",
            6 => "LX",
            7 => "LXX",
            8 => "LXXX",
            9 => "XC",
            _ => "",
        } +
        match ones {
            1 => "I",
            2 => "II",
            3 => "III",
            4 => "IV",
            5 => "V",
            6 => "VI",
            7 => "VII",
            8 => "VIII",
            9 => "IX",
            _ => "",
        }
    }
}
