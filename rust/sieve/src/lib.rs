use std::collections::BTreeSet;

pub fn primes_up_to(up_to: usize) -> Vec<usize> {
    let mut composite_numbers: BTreeSet<usize> = BTreeSet::new();
    let mut primes: Vec<usize> = Vec::new();

    for number in 2..up_to + 1 {
        let mut multiplier = 2;

        if !composite_numbers.contains(&number) {
            primes.push(number);
        }

        while number * multiplier <= up_to {
            composite_numbers.insert(number * multiplier);
            multiplier += 1;
        }
    }

    return primes;
}
