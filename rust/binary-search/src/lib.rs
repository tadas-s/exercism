use std::cmp::Ordering::{Equal, Greater, Less};

pub fn find<T: AsRef<[U]>, U: Ord>(array: T, key: U) -> Option<usize> {
    let array = array.as_ref();

    let mut midpoint: usize;
    let mut left: usize = 0;
    let mut right: usize = array.len();

    while left < right {
        midpoint = (left + right) / 2;

        match array[midpoint].cmp(&key) {
            Equal => return Some(midpoint),
            Less => left = midpoint + 1,
            Greater => right = midpoint,
        };
    }

    None
}
