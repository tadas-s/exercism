use std::collections::HashMap;

pub fn count(nucleotide: char, strand: &str) -> usize {
    strand.chars().fold(0, |count, c| count + if nucleotide == c { 1 } else { 0 })
}

pub fn nucleotide_counts(strand: &str) -> HashMap<char, usize> {
    let mut counts = HashMap::with_capacity(4);
    
    for nucleotide in vec!['C', 'T', 'A', 'G'] {
        counts.insert(nucleotide, count(nucleotide, strand));
    }

    return counts;
}

