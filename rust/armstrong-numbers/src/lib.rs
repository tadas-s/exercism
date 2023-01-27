pub fn is_armstrong_number(num: u32) -> bool {
    let digits = num
        .to_string()
        .chars()
        .filter_map(|c| c.to_digit(10))
        .collect::<Vec<_>>();
    let to_power = digits.len() as u32;

    let sum_attempt = digits
        .iter()
        .map(|d| d.pow(to_power))
        .fold(Some(0 as u32), |acc, m| {
            acc.and_then(|acc| acc.checked_add(m))
        });

    match sum_attempt {
        Some(sum) => sum == num,
        None => false,
    }
}
