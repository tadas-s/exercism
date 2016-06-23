pub fn verse(verse: u8) -> String {
   format!(
       "{} of beer on the wall, {} of beer on the wall\nGo to the store and buy some more, {} bottles of beer on the wall.",
       pluralize_bottle_pile_of(verse),
       pluralize_bottle_pile_of(verse),
       pluralize_bottle_pile_of(next(verse))
   )
}

pub fn sing(from_verse: u8, to_verse: u8) -> String {
    "".to_string()
}

fn pluralize_bottle_pile_of(number: u8) -> String {
    match number {
        0 => "no more bottles".to_string(),
        _ if (number % 10 == 1) && (number != 11) => format!("{} bottle", number),
        _ => format!("{} bottles", number)
    }
}

fn next(n: u8) -> u8 {
    if n > 0 { n - 1 } else { 99 }
}
