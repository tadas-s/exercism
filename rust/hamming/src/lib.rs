pub fn hamming_distance(a: &str, b: &str) -> Result<u64, ()> {
    if a.len() !=  b.len() {
        return Err(());
    }

    Ok(
        a.chars()
            .zip(b.chars())
            .fold(0, |diffs, (ca, cb)| diffs + if ca == cb { 0 } else { 1 })
    )
}
