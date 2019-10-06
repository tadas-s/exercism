pub fn factors(n: u64) -> Vec<u64> {
    let mut result = Vec::new();
    let mut remaining = n;
    let mut candidate = 2;

    while remaining > 1 {
        while remaining % candidate == 0 {
            result.push(candidate);
            remaining /= candidate;
        }

        candidate += 1;
    }

    result
}
