///
/// Convert a number between two bases.
///
/// A number is any slice of digits.
/// A digit is any unsigned integer (e.g. u8, u16, u32, u64, or usize).
/// Bases are specified as unsigned integers.
///
/// Return an `Err(.)` if the conversion is impossible.
/// The tests do not test for specific values inside the `Err(.)`.
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
///  * Never output leading 0 digits. However, your function must be able to
///     process input with leading 0 digits.
///
#[allow(unused_variables)]
pub fn convert(number: &[u32], from_base: u32, to_base: u32) -> Result<Vec<u32>, ()> {
    let mut u64number: usize = 0;

    if from_base < 2 || to_base < 2 {
        return Err(());
    }

    for (i, digit) in number.iter().rev().enumerate() {
        if *digit >= from_base {
            return Err(());
        }

        u64number += (from_base as usize).pow(i as u32) * (*digit as usize);
    }

    // that's rather dirty
    if u64number == 0 { return Ok(vec![]); }

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
