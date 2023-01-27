/// Check a Luhn checksum.
pub fn is_valid(code: &str) -> bool {
    let chars = code.chars().collect::<Vec<_>>();

    if chars.iter().any(|&c| !(c.is_ascii_digit() || c == ' ')) {
        return false;
    }

    let digits = chars
        .iter()
        .filter(|c| c.is_ascii_digit())
        .filter_map(|c| c.to_digit(10))
        .collect::<Vec<_>>();

    if digits.len() < 2 {
        return false;
    }

    let parity = digits.len() % 2;

    let checksum = digits
        .iter()
        .enumerate()
        .map(|(i, &digit)| {
            if i % 2 == parity && digit > 4 {
                digit * 2 - 9
            } else if i % 2 == parity {
                digit * 2
            } else {
                digit
            }
        })
        .sum::<u32>();

    checksum % 10 == 0
}
