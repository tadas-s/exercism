use std::collections::HashMap;

pub fn count(nucleotide: char, dna: &str) -> Result<usize, char> {
    let mut counts = nucleotide_counts(dna)?;

    counts.remove(&nucleotide).ok_or(nucleotide)
}

pub fn nucleotide_counts(dna: &str) -> Result<HashMap<char, usize>, char> {
    let mut counts = HashMap::from([('C', 0), ('T', 0), ('A', 0), ('G', 0)]);

    for nucleotide in dna.chars() {
        counts.get_mut(&nucleotide).map(|count| *count += 1).ok_or(nucleotide)?
    }

    Ok(counts)
}
