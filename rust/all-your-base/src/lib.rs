use num::{checked_pow, pow, CheckedMul, Num, ToPrimitive};
use std::fmt::Debug;

#[derive(Debug, PartialEq, Eq)]
pub enum Error<T> {
    InvalidInputBase,
    InvalidOutputBase,
    InvalidDigit(T),
}

pub fn convert<T: PartialOrd + Copy + Num + CheckedMul + ToPrimitive>(
    number: &[T],
    from_base: T,
    to_base: T,
) -> Result<Vec<T>, Error<T>> {
    let mut n = T::zero();

    if from_base <= T::one() {
        return Err(Error::InvalidInputBase);
    }
    if to_base <= T::one() {
        return Err(Error::InvalidOutputBase);
    }

    for (exponent, &digit) in number.iter().rev().enumerate() {
        if digit >= from_base {
            return Err(Error::InvalidDigit(digit));
        }

        if digit > T::zero() { // to mult overflow if input has leading zero and is close to format maximum
            n = n + pow(from_base, exponent) * digit;
        }
    }

    if n == T::zero() {
        return Ok(vec![T::zero()]);
    }

    let mut result: Vec<T> = vec![];

    for exponent in 0.. {
        if n % pow(to_base, exponent) == n {
            break;
        }

        if let Some(next_divisor) = checked_pow(to_base, exponent + 1) {
            result.push(n % next_divisor / pow(to_base, exponent));
        } else {
            // This is the last most significant digit
            result.push(n / pow(to_base, exponent));
            break;
        }
    }

    result.reverse();

    return Ok(result);
}
