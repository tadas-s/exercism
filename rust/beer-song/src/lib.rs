pub fn verse(verse: u8) -> String {
    format!("{}\n{}\n", line(verse * 2), line(verse * 2 + 1))
}

pub fn sing(start_with: u8, end_with: u8) -> String {
    (end_with..start_with + 1)
        .rev()
        .map(verse)
        .collect::<Vec<String>>()
        .join("\n")
}

fn line(line: u8) -> String {
    match line {
        0 => "No more bottles of beer on the wall, no more bottles of beer.".to_string(),
        1 => "Go to the store and buy some more, 99 bottles of beer on the wall.".to_string(),
        3 => "Take it down and pass it around, no more bottles of beer on the wall.".to_string(),
        _ if line % 2 == 0 => format!("{} of beer on the wall, {} of beer.", pluralize_bottle_pile_of(line / 2), pluralize_bottle_pile_of(line / 2)),
        _ if line % 2 == 1 => format!("Take one down and pass it around, {} of beer on the wall.", pluralize_bottle_pile_of(line / 2 - 1)),
        _ => panic!("This will never happen. This also looks like not a great code as well.")
    }
}

fn pluralize_bottle_pile_of(number: u8) -> String {
    match number {
        0 => "no more bottles".to_string(),
        _ if (number % 10 == 1) && (number != 11) => format!("{} bottle", number),
        _ => format!("{} bottles", number)
    }
}

