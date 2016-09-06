pub fn square_of_sum(number: u64) -> u64 {
    (1..number + 1).sum::<u64>().pow(2)
}

pub fn sum_of_squares(number: u64) -> u64 {
    (1..number + 1).map(|n| n.pow(2)).sum()
}

pub fn difference(number: u64) -> u64 {
    square_of_sum(number) - sum_of_squares(number)
}
