use num::{checked_pow, pow, CheckedMul, Num, ToPrimitive};
use std::fmt::Debug;

#[derive(Debug, PartialEq, Eq)]
pub enum Error<T> {
    InvalidInputBase,
    InvalidOutputBase,
    InvalidDigit(T),
}

///
/// Convert a number between two bases.
///
/// A number is any slice of digits.
/// A digit is any unsigned integer (e.g. u8, u16, u32, u64, or usize).
/// Bases are specified as unsigned integers.
///
/// Return the corresponding Error enum if the conversion is impossible.
///
///
/// You are allowed to change the function signature as long as all test still pass.
///
///
/// Example:
/// Input
///   number: &[4, 2]
///   from_base: 10
///   to_base: 2
/// Result
///   Ok(vec![1, 0, 1, 0, 1, 0])
///
/// The example corresponds to converting the number 42 from decimal
/// which is equivalent to 101010 in binary.
///
///
/// Notes:
///  * The empty slice ( "[]" ) is equal to the number 0.
///  * Never output leading 0 digits, unless the input number is 0, in which the output must be `[0]`.
///    However, your function must be able to process input with leading 0 digits.
///
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

        if digit > T::zero() {
            // to mult overflow if input has leading zero and is close to format maximum
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

    Ok(result)
}
