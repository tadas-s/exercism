use std::collections::HashMap;

static NUCLEOTIDES: [char; 4] = ['C', 'T', 'A', 'G'];

pub fn count(nucleotide: char, dna: &str) -> Result<usize, char> {
    if !NUCLEOTIDES.contains(&nucleotide) {
        return Err(nucleotide);
    }

    if let Some(bad_nucleotide_in_sequence) = dna.chars().find(|n| !NUCLEOTIDES.contains(n)) {
        return Err(bad_nucleotide_in_sequence);
    }

    Ok(dna.chars().filter(|&n| nucleotide == n).count())
}

pub fn nucleotide_counts(dna: &str) -> Result<HashMap<char, usize>, char> {
    let mut counts = HashMap::with_capacity(NUCLEOTIDES.len());

    for nucleotide in NUCLEOTIDES {
        match count(nucleotide, dna) {
            Ok(count) => counts.insert(nucleotide, count),
            Err(nucleotide) => return Err(nucleotide),
        };
    }

    Ok(counts)
}
