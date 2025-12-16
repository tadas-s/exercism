pub fn factors(n: u64) -> Vec<u64> {
    let mut result = vec![];
    let mut remaining = n;
    let mut candidates = 2..=n;

    while remaining > 1 {
        let candidate = candidates.next().unwrap();

        while remaining % candidate == 0 {
            result.push(candidate);
            remaining /= candidate;
        }
    }

    result
}
