pub fn is_pangram(s: &str) -> bool {
    let lowercased = s.to_lowercase();

    for c in ('a' as u8)..('z' as u8 + 1) {
        if !lowercased.contains(c as char) {
            return false;
        }
    }

    return true;
}
