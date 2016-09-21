extern crate num;

use num::{Num, ToPrimitive};
use num::pow::pow;

pub fn convert<T: Num + PartialOrd + Copy + ToPrimitive>(number: &[T], from_base: T, to_base: T) -> Result<Vec<T>, ()> {
    let mut whole_value: T = T::zero();

    if from_base <= T::one() || to_base <= T::one() {
        return Err(());
    }

    for (i, digit) in number.iter().rev().enumerate() {
        if *digit >= from_base {
            return Err(());
        }

        whole_value = whole_value + pow(from_base, i) * (*digit);
    }

    // that's rather dirty
    if whole_value == T::zero() { return Ok(vec![]); }

    let mut digit = T::zero();
    let mut result: Vec<T> = vec![];

    loop {
        result.push(whole_value % pow(to_base, digit.to_usize().unwrap() + 1) / pow(to_base, digit.to_usize().unwrap()));

        digit = digit + T::one();

        if whole_value % pow(to_base, digit.to_usize().unwrap()) == whole_value {
            break;
        }
    }

    result.reverse();

    return Ok(result);
}
