pub fn brackets_are_balanced(string: &str) -> bool {
    let mut stack: Vec<char> = Vec::new();

    for c in string.chars() {
        match c {
            '[' | '{' | '(' => stack.push(c),
            ']' | '}' | ')' => {
                if stack.pop() != Some(openning_bracket(c)) {
                    return false;
                }
            }
            _ => (),
        }
    }

    stack.len() == 0
}

fn openning_bracket(c: char) -> char {
    match c {
        ']' => '[',
        '}' => '{',
        ')' => '(',
        _ => panic!("No openning bracket for char '{}'", c),
    }
}
