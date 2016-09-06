pub fn hamming_distance(a: &str, b: &str) -> Result<u64, &'static str> {
    if a.len() !=  b.len() {
        return Err("Can only compare strings of equal length");
    }

    Ok(
        a.chars()
            .zip(b.chars())
            .map(|(ca, cb)| if ca == cb { 0 } else { 1 })
            .sum()
    )
}
