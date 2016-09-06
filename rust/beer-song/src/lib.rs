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
        2 => "1 bottle of beer on the wall, 1 bottle of beer.".to_string(),
        3 => "Take it down and pass it around, no more bottles of beer on the wall.".to_string(),
        5 => "Take one down and pass it around, 1 bottle of beer on the wall.".to_string(),
        _ if line % 2 == 0 => format!("{} bottles of beer on the wall, {} bottles of beer.", line / 2, line / 2),
        _ => format!("Take one down and pass it around, {} bottles of beer on the wall.", line / 2 - 1)
    }
}

