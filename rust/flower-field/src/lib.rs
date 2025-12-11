pub fn annotate(garden: &[&str]) -> Vec<String> {
    let mut marked: Vec<String> = Vec::with_capacity(garden.len());

    for y in 0..garden.len() {
        let this_line = garden[y]
            .char_indices()
            .map(|(x, square)| {
                if square == '*' {
                    return '*';
                }

                let mines = DIRECTIONS
                    .iter()
                    .filter_map(|d| square_at(garden, x, d.0, y, d.1))
                    .filter(|s| *s == '*')
                    .count();

                if mines > 0 {
                    mines.to_string().chars().next().unwrap()
                } else {
                    ' '
                }
            })
            .collect();

        marked.push(this_line);
    }

    marked
}

#[rustfmt::skip]
static DIRECTIONS: &[(i32, i32)] = &[
    (-1, -1), (0, -1), (1, -1),
    (-1,  0),          (1,  0),
    (-1,  1), (0,  1), (1,  1),
];

fn square_at(garden: &[&str], x: usize, dx: i32, y: usize, dy: i32) -> Option<char> {
    if (y as i32 + dy) < 0
        || (x as i32 + dx) < 0
        || (y as i32 + dy) >= garden.len() as i32
        || (x as i32 + dx) >= garden[y].len() as i32
    {
        return None;
    }

    garden[(y as i32 + dy) as usize]
        .chars()
        .nth((x as i32 + dx) as usize)
}
