#[derive(Debug, PartialEq, Eq)]
pub enum Error {
    InvalidInputBase,
    InvalidOutputBase,
    InvalidDigit(u32),
}

pub fn convert(number: &[u32], from_base: u32, to_base: u32) -> Result<Vec<u32>, Error> {
    let mut u64number: usize = 0;

    if from_base < 2 { return Err(Error::InvalidInputBase); }
    if to_base < 2 { return Err(Error::InvalidOutputBase); }

    for (i, &digit) in number.iter().rev().enumerate() {
        if digit >= from_base {
            return Err(Error::InvalidDigit(digit));
        }

        u64number += (from_base as usize).pow(i as u32) * (digit as usize);
    }

    let mut digit = 0;
    let mut result: Vec<u32> = vec![];

    loop {
        result.push((u64number % (to_base as usize).pow(digit + 1) / (to_base as usize).pow(digit)) as u32);

        digit += 1;

        if u64number % (to_base as usize).pow(digit) == u64number {
            break;
        }
    }

    result.reverse();

    return Ok(result);
}
