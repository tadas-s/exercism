use std::cmp::Ordering::{Less, Equal, Greater};

pub fn find<T: AsRef<[U]>, U: Ord>(array: T, key: U) -> Option<usize> {
    let array = array.as_ref();

    if array.len() == 0 {
        return None;
    }

    let mut midpoint = array.len() / 2;
    let mut left: usize = 0;
    let mut right: usize = array.len() - 1;

    loop {
        match key.cmp(&array[midpoint]) {
            Less if right == left => return None,
            Less => {
                right = midpoint - 1;
                midpoint = (left + right) / 2;
            },
            Equal => return Some(midpoint),
            Greater if right == left => return None,
            Greater => {
                left = midpoint + 1;
                midpoint = (left + right) / 2;
            },
        };
    }
}
