pub fn verse(n: i32) -> String {
    let next = if n > 0 { n - 1 } else { 99 };

    format!(
        "{} of beer on the wall, {} of beer.\n{}, {} of beer on the wall.\n",
        bottles(n, true),
        bottles(n, false),
        match n {
            0 => "Go to the store and buy some more",
            1 => "Take it down and pass it around",
            _ => "Take one down and pass it around",
        },
        bottles(next, false)
    )
    .to_string()
}

pub fn sing(start: i32, end: i32) -> String {
    (end..=start)
        .rev()
        .map(|v| verse(v))
        .collect::<Vec<_>>()
        .join("\n")
}

fn bottles(n: i32, uppercase: bool) -> String {
    match n {
        0 => (if uppercase {
            "No more bottles"
        } else {
            "no more bottles"
        })
        .to_string(),
        11 => n.to_string() + " bottles",
        _ if n % 10 == 1 => n.to_string() + " bottle",
        _ => n.to_string() + " bottles",
    }
}
