struct Primes {
    current: Option<u32>,
}

impl Primes {
    fn new() -> Primes {
        Self { current: None }
    }

    /**
    * Trial division prime test: https://en.wikipedia.org/wiki/Trial_division
    */
    fn is_prime(n: u32) -> bool {
        match n {
            1 => false,
            2 | 3 => true,
            _ if n % 2 == 0 || n % 3 == 0 => false,
            _ => !(5..)
                .step_by(6)
                .take_while(|divisor| divisor * divisor <= n)
                .any(|divisor| n % divisor == 0 || n % (divisor + 2) == 0)
        }
    }
}

impl Iterator for Primes {
    type Item = u32;

    fn next(&mut self) -> Option<u32> {
        match self.current {
            Some(prime) => {
                self.current = ((prime + 1)..).find(|&number| Primes::is_prime(number));
            }
            _ => self.current = Some(2)
        }

        self.current
    }
}

pub fn nth(n: u32) -> u32 {
    let mut primes = Primes::new();

    primes.nth(n as usize).unwrap()
}
