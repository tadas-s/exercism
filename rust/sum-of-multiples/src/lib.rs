pub fn sum_of_multiples(limit: u32, factors: &[u32]) -> u32 {
    let mut multiples = factors
        .iter()
        .filter(|&f| *f > 0)
        .flat_map(|f| (1..=((limit - 1) / f)).map(move |multiplier| f * multiplier))
        .collect::<Vec<u32>>();

    multiples.sort();
    multiples.dedup();

    multiples.iter().sum()
}
