static DIRECTIONS: &'static [(i32, i32)] = &[
    (-1, -1),
    (0, -1),
    (1, -1),
    (-1, 0),
    (1, 0),
    (-1, 1),
    (0, 1),
    (1, 1),
];

pub fn annotate(minefield: &[&str]) -> Vec<String> {
    let mut marked: Vec<String> = Vec::with_capacity(minefield.len());

    for y in 0..minefield.len() {
        let mut this_line = "".to_string();

        for (x, square) in minefield[y].char_indices() {
            if square == '*' {
                this_line.push('*');
                continue;
            }

            let mines = DIRECTIONS
                .iter()
                .filter_map(|d| square_at(minefield, x, d.0, y, d.1))
                .filter(|s| *s == '*')
                .count();

            if mines > 0 {
                this_line.push(mines.to_string().chars().nth(0).unwrap())
            } else {
                this_line.push(' ');
            }
        }

        marked.push(this_line);
    }

    marked
}

fn square_at(minefield: &[&str], x: usize, dx: i32, y: usize, dy: i32) -> Option<char> {
    if (y as i32 + dy) < 0
        || (x as i32 + dx) < 0
        || (y as i32 + dy) >= minefield.len() as i32
        || (x as i32 + dx) >= minefield[y].len() as i32
    {
        return None;
    }

    minefield[(y as i32 + dy) as usize]
        .chars()
        .nth((x as i32 + dx) as usize)
}
