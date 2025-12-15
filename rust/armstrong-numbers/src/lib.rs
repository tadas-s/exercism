pub fn is_armstrong_number(num: u32) -> bool {
    let digits = num
        .to_string()
        .chars()
        .filter_map(|c| c.to_digit(10))
        .collect::<Vec<_>>();

    let exponent = digits.len() as u32;

    let sum = digits
        .iter()
        .map(|d| d.pow(exponent))
        .sum();

    num == sum
}
