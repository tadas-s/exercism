struct Primes {
    current: Option<u32>
}

impl Primes {
    fn new() -> Primes {
        Primes { current: None }
    }

    /**
     * Trial division prime test: https://en.wikipedia.org/wiki/Trial_division
     */
    fn is_prime(n: u32) -> bool {
        if n <= 1 {
            return false;
        } else if n <= 3 {
            return true;
        } else if n % 2 == 0 || n % 3 == 0 {
            return false;
        } else {
            let mut divisor = 5;
            while divisor * divisor <= n {
                if n % divisor == 0 || n % (divisor + 2) == 0 {
                    return false;
                }
                divisor += 6
            }
            return true;
        }
    }
}

impl Iterator for Primes {
    type Item = u32;

    fn next(&mut self) -> Option<u32> {
        match self.current {
            Some(prime) => {
                let mut next_prime = prime;

                loop {
                    next_prime = next_prime + 1;

                    if Primes::is_prime(next_prime) {
                        break;
                    }
                }
                self.current = Some(next_prime);
            }
            _ => { self.current = Some(2) }
        }

        return self.current;
    }
}

pub fn nth(n: u32) -> u32 {
    let mut primes: Primes = Primes::new();
    return primes.nth(n as usize).unwrap();
}
