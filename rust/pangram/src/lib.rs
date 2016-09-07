static ALPHABET: &'static str = "abcdefghijklmnopqrstuvwxyz";

pub fn is_pangram(s: &str) -> bool {
    let lowercased = s.to_lowercase();
    ALPHABET.chars().all(|c| lowercased.contains(c))
}
