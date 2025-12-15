pub fn recite(start_bottles: u32, take_down: u32) -> String {
    ((start_bottles - take_down + 1)..=start_bottles)
        .rev()
        .map(|bottles| verse(bottles))
        .collect::<Vec<_>>()
        .join("\n")
}

fn verse(bottles: u32) -> String {
    let remaining = bottles - 1;

    format!(
        concat!(
            "{} hanging on the wall,\n",
            "{} hanging on the wall,\n",
            "And if one green bottle should accidentally fall,\n",
            "There'll be {} hanging on the wall.\n"
        ),
        upcase(bottles_in_words(bottles)),
        upcase(bottles_in_words(bottles)),
        bottles_in_words(remaining)
    )
    .to_string()
}

fn bottles_in_words(n: u32) -> String {
    (match n {
        10 => "ten green bottles",
        9 => "nine green bottles",
        8 => "eight green bottles",
        7 => "seven green bottles",
        6 => "six green bottles",
        5 => "five green bottles",
        4 => "four green bottles",
        3 => "three green bottles",
        2 => "two green bottles",
        1 => "one green bottle",
        0 => "no green bottles",
        _ => panic!(),
    })
    .to_string()
}

fn upcase(s: String) -> String {
    s.chars()
        .enumerate()
        .map(|(i, c)| {
            if i == 0 {
                c.to_uppercase().to_string()
            } else {
                c.to_string()
            }
        })
        .collect()
}
