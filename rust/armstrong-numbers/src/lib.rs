pub fn is_armstrong_number(num: u32) -> bool {
    let digits = num
        .to_string()
        .chars()
        .filter_map(|c| c.to_digit(10))
        .collect::<Vec<_>>();
    let to_power = digits.len() as u32;

    let sum_attempt =
        digits
            .iter()
            .map(|d| d.pow(to_power))
            .fold(Some(0 as u32), |acc, m| match acc {
                Some(sum) => sum.checked_add(m),
                None => None,
            });

    match sum_attempt {
        Some(sum) => sum == num,
        None => false,
    }
}
